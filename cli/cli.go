// API呼び出しのCLI
package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"reflect"
	"strings"
	"time"
	"unicode"

	"gopkg.in/yaml.v1"

	log "github.com/Sirupsen/logrus"
	"github.com/mattn/go-shellwords"

	"github.com/iij/p2pubapi"
	"github.com/iij/p2pubapi/protocol"
	"github.com/urfave/cli"
	"github.com/wtnb75/go-cmdrepl"
)

func shortname(s string) string {
	// filter uppercase
	res := ""
	for _, c := range s {
		if unicode.IsUpper(c) {
			res += string(c)
		}
	}
	return res
}

func toEnv(s string) string {
	return strings.ToUpper(s)
}

func cmdfunc(c *cli.Context) error {
	if c.GlobalBool("verbose") {
		log.SetLevel(log.DebugLevel)
	}
	typ := protocol.TypeMap[c.Command.Name]

	if typ == nil {
		return fmt.Errorf("no such subcommand? %s", c.Command.Name)
	}
	var data = map[string]string{}
	arg := reflect.Zero(typ).Interface().(protocol.CommonArg)
	log.Debugf("found %s %+v", c.Command.Name, arg)
	req, opt := p2pubapi.ArgumentList(arg)
	for _, i := range append(req, opt...) {
		val := c.String(i)
		if val == "" {
			val = c.GlobalString(i)
		}
		if val == "" {
			continue
		}
		log.Debug("setting", i, val)
		data[i] = val
	}
	log.Debugf("data: %#+v", data)
	api := p2pubapi.NewAPI(c.GlobalString("AccessKey"), c.GlobalString("SecretKey"))
	var resp = map[string]interface{}{}
	if err := p2pubapi.ValidateMap(c.Command.Name, data); err != nil {
		log.Error(err)
		return err
	}
	if err := p2pubapi.CallWithMap(*api, c.Command.Name, data, resp); err != nil {
		log.Error(err)
		return err
	}
	log.Debug("resp", resp)
	switch c.GlobalString("format") {
	case "golang":
		fmt.Printf("%+v", resp)
	case "yaml":
		if out, err := yaml.Marshal(resp); err == nil {
			os.Stdout.Write(out)
		} else {
			log.Error(err)
		}
	case "pretty":
		if b, err := json.MarshalIndent(resp, "", "  "); err == nil {
			os.Stdout.Write(b)
		} else {
			log.Error(err)
		}
	default: // json
		enc := json.NewEncoder(os.Stdout)
		enc.Encode(resp)
	}
	return nil
}

// GetCLI cliモジュールのサブコマンド構造体を求める
func GetCLI(arg protocol.CommonArg) cli.Command {
	res := cli.Command{}
	res.Name = arg.APIName()
	res.Aliases = []string{shortname(res.Name), arg.JPName()}
	res.Usage = arg.JPName()
	res.Description = fmt.Sprintf("calls %s API\n    %s %s\n    see: %s (%s)\n", arg.APIName(), arg.Method(), arg.URI(), arg.Document(), arg.JPName())
	res.Action = cmdfunc
	req, opt := p2pubapi.ArgumentList(arg)
	res.Flags = []cli.Flag{}
	for _, o := range req {
		flg := cli.StringFlag{}
		flg.Name = fmt.Sprintf("%s", o)
		if strings.HasSuffix(flg.Name, "ServiceCode") {
			flg.Name = flg.Name + "," + strings.TrimSuffix(flg.Name, "ServiceCode")
		}
		flg.EnvVar = toEnv(o)
		flg.Usage = "required"
		res.Flags = append(res.Flags, flg)
	}
	for _, o := range opt {
		flg := cli.StringFlag{}
		flg.Name = fmt.Sprintf("%s", o)
		if strings.HasSuffix(flg.Name, "ServiceCode") {
			flg.Name = flg.Name + "," + strings.TrimSuffix(flg.Name, "ServiceCode")
		}
		flg.EnvVar = toEnv(o)
		flg.Usage = "optional"
		res.Flags = append(res.Flags, flg)
	}
	return res
}

func string2status(s string) p2pubapi.Status {
	switch strings.ToLower(s) {
	default:
		return p2pubapi.None
	case strings.ToLower(p2pubapi.InPreparation.String()):
		return p2pubapi.InPreparation
	case strings.ToLower(p2pubapi.InService.String()):
		return p2pubapi.InService
	case strings.ToLower(p2pubapi.Stopped.String()):
		return p2pubapi.Stopped
	case strings.ToLower(p2pubapi.Configuring.String()):
		return p2pubapi.Configuring
	case strings.ToLower(p2pubapi.Starting.String()):
		return p2pubapi.Starting
	case strings.ToLower(p2pubapi.Running.String()):
		return p2pubapi.Running
	case strings.ToLower(p2pubapi.Stopping.String()):
		return p2pubapi.Stopping
	case strings.ToLower(p2pubapi.Locked.String()):
		return p2pubapi.Locked
	case strings.ToLower(p2pubapi.Attached.String()):
		return p2pubapi.Attached
	case strings.ToLower(p2pubapi.NotAttached.String()):
		return p2pubapi.NotAttached
	case strings.ToLower(p2pubapi.Initializing.String()):
		return p2pubapi.Initializing
	case strings.ToLower(p2pubapi.Archiving.String()):
		return p2pubapi.Archiving
	case strings.ToLower(p2pubapi.Initialized.String()):
		return p2pubapi.Initialized
	case strings.ToLower(p2pubapi.Configured.String()):
		return p2pubapi.Configured
	case strings.ToLower(p2pubapi.Updating.String()):
		return p2pubapi.Updating
	}
}

func getwaitinfo(arg []string) (scode string, cst, sst p2pubapi.Status) {
	scode = arg[0]
	cst = p2pubapi.None
	sst = p2pubapi.None
	for _, v := range arg[1:len(arg)] {
		st := string2status(v)
		if st == p2pubapi.InPreparation || st == p2pubapi.InService {
			cst = st
		} else {
			sst = st
		}
	}
	return
}

func main() {
	app := cli.NewApp()
	app.EnableBashCompletion = true
	app.Version = version
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:   "AccessKey",
			EnvVar: "IIJAPI_ACCESS_KEY",
		},
		cli.StringFlag{
			Name:   "SecretKey",
			EnvVar: "IIJAPI_SECRET_KEY",
		},
		cli.BoolFlag{
			Name: "verbose, V",
		},
		cli.StringFlag{
			Name:   "format, f",
			Value:  "json",
			Usage:  "[json|yaml|pretty|golang]",
			EnvVar: "FORMAT",
		},
	}
	waitflag := []cli.Flag{
		cli.StringFlag{
			Name:   "GisServiceCode,Gis",
			EnvVar: "GISSERVICECODE,GisServiceCode",
		},
		cli.DurationFlag{
			Name:  "duration,max-wait",
			Value: time.Duration(10 * time.Minute),
		},
	}
	app.Commands = []cli.Command{
		{
			Name:    "setenv",
			Aliases: []string{"set", "export"},
			Usage:   "set environment variable",
			Action: func(ctxt *cli.Context) error {
				args := ctxt.Args()
				if !args.Present() {
					log.Error("require arguments")
					return fmt.Errorf("require arguments")
				}
				for _, v := range args {
					tmp := strings.SplitN(v, "=", 2)
					if len(tmp) == 1 {
						log.Debugf("set %s=%s", tmp[0], os.Getenv(tmp[0]))
					} else {
						oldval := os.Getenv(tmp[0])
						if tmp[1] == "" {
							log.Debugf("%s = %s -> (unset)", tmp[0], oldval)
						} else {
							os.Setenv(tmp[0], tmp[1])
							log.Debugf("%s = %s -> %s", tmp[0], oldval, os.Getenv(tmp[0]))
						}
					}
				}
				return nil
			},
		}, {
			Name:    "getenv",
			Aliases: []string{"get", "env"},
			Usage:   "get environment variable",
			Action: func(ctxt *cli.Context) error {
				args := ctxt.Args()
				if args.Present() {
					for _, k := range append([]string{args.First()}, args.Tail()...) {
						log.Infof("set %s=%s", k, os.Getenv(k))
					}
				} else {
					for _, k := range os.Environ() {
						log.Info(k)
					}
				}
				return nil
			},
		}, {
			Name:    "exit",
			Aliases: []string{"quit", "q"},
			Hidden:  true,
			Action: func(c *cli.Context) error {
				return fmt.Errorf("exit")
			},
		}, {
			Name:    "source",
			Aliases: []string{".", "exec"},
			Action: func(c *cli.Context) error {
				for _, k := range c.Args() {
					if fp, err := os.Open(k); err == nil {
						defer fp.Close()
						scanner := bufio.NewScanner(fp)
						for scanner.Scan() {
							var tokens []string
							tokens, err = shellwords.Parse(scanner.Text())
							log.Debug("execute ", tokens)
							if err = app.Run(append([]string{tokens[0]}, tokens...)); err != nil {
								log.Error("cmd error:", err)
								return err
							}
						}
					} else {
						log.Error("open error:", err)
						return err
					}
				}
				return nil
			},
		}, {
			Name:    "waitVM",
			Aliases: []string{"waitvm", "vmwait", "wait"},
			Flags:   waitflag,
			Action: func(c *cli.Context) error {
				arg := c.Args()
				if len(arg) == 0 {
					log.Error("Usage: waitVM ivmXXXXXXXX [contractStatus] [vmStatus]")
					return fmt.Errorf("Usage: waitVM ivmXXXXXXXX [contractStatus] [vmStatus]")
				}
				ivm, cst, sst := getwaitinfo(arg)
				api := p2pubapi.NewAPI(c.GlobalString("AccessKey"), c.GlobalString("SecretKey"))
				gis := c.String("GisServiceCode")
				res := p2pubapi.WaitVM(api, gis, ivm, cst, sst, c.Duration("duration"))
				return res
			},
		}, {
			Name:    "waitSystemStorage",
			Aliases: []string{"waitsst", "sstwait"},
			Flags:   waitflag,
			Action: func(c *cli.Context) error {
				arg := c.Args()
				if len(arg) == 0 {
					log.Error("Usage: waitSystemStorage ibaXXXXXXXX [contractStatus] [ibaStatus]")
					return fmt.Errorf("Usage: waitSystemStorage ibaXXXXXXXX [contractStatus] [ibaStatus]")
				}
				iba, cst, sst := getwaitinfo(arg)
				api := p2pubapi.NewAPI(c.GlobalString("AccessKey"), c.GlobalString("SecretKey"))
				gis := c.String("GisServiceCode")
				res := p2pubapi.WaitSystemStorage(api, gis, iba, cst, sst, c.Duration("duration"))
				return res
			},
		}, {
			Name:    "waitDataStorage",
			Aliases: []string{"waitst", "stwait"},
			Flags:   waitflag,
			Action: func(c *cli.Context) error {
				arg := c.Args()
				if len(arg) == 0 {
					log.Error("Usage: waitDataStorage i??XXXXXXXX [contractStatus] [ibaStatus]")
					return fmt.Errorf("Usage: waitSystemStorage i??XXXXXXXX [contractStatus] [ibaStatus]")
				}
				iba, cst, sst := getwaitinfo(arg)
				api := p2pubapi.NewAPI(c.GlobalString("AccessKey"), c.GlobalString("SecretKey"))
				gis := c.String("GisServiceCode")
				res := p2pubapi.WaitDataStorage(api, gis, iba, cst, sst, c.Duration("duration"))
				return res
			},
		},
	}
	for _, p := range protocol.APIlist {
		subcommand := GetCLI(p)
		app.Commands = append(app.Commands, subcommand)
	}
	if len(os.Args) == 1 {
		cmdrepl.CmdRepl("[P2PUB]# ", app)
	} else {
		app.Run(os.Args)
	}
}

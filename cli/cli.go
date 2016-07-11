// API呼び出しのCLI
package main

import (
	"encoding/json"
	"fmt"
	"os"
	"reflect"
	"strings"
	"unicode"

	"gopkg.in/yaml.v1"

	log "github.com/Sirupsen/logrus"

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
	res.Usage = ""
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
		flg.EnvVar = toEnv(o)
		flg.Usage = "optional"
		res.Flags = append(res.Flags, flg)
	}
	return res
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
	app.Commands = []cli.Command{
		{
			Name:    "setenv",
			Aliases: []string{"set"},
			Usage:   "set environment variable",
			Action: func(ctxt *cli.Context) error {
				args := ctxt.Args()
				if !args.Present() {
					log.Error("require arguments")
					return fmt.Errorf("require arguments")
				}
				k := args.First()
				preval := os.Getenv(k)
				if args.Get(1) == "" {
					os.Unsetenv(k)
					log.Infof("%s = %s -> (unset)", args.First(), preval)
				} else {
					os.Setenv(k, args.Get(1))
					log.Infof("%s = %s -> %s", args.First(), preval, os.Getenv(k))
				}
				return nil
			},
		}, {
			Name:    "getenv",
			Aliases: []string{"get"},
			Usage:   "get environment variable",
			Action: func(ctxt *cli.Context) error {
				args := ctxt.Args()
				if args.Present() {
					for _, k := range append([]string{args.First()}, args.Tail()...) {
						log.Infof("%s=%s", k, os.Getenv(k))
					}
				} else {
					for _, k := range os.Environ() {
						log.Info(k)
					}
				}
				return nil
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

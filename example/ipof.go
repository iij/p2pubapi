package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/iij/p2pubapi"
	"github.com/iij/p2pubapi/protocol"
	"github.com/urfave/cli"
)

// Usage: $0 [--global|--private] IvmServiceCode

func doMain(c *cli.Context) error {
	api := p2pubapi.NewAPI(c.String("access-key"), c.String("secret-key"))
	gis := c.String("gis-service-code")
	for _, v := range c.Args() {
		arg := protocol.VMGet{
			GisServiceCode: gis,
			IvmServiceCode: v,
		}
		if c.Bool("verbose") {
			fmt.Printf("arg=%+v\n", arg)
		}
		var res = protocol.VMGetResponse{}
		if err := p2pubapi.Call(*api, arg, &res); err != nil {
			return err
		}
		if c.Bool("verbose") {
			fmt.Printf("res=%+v\n", res)
		}
		for _, n := range res.NetworkList {
			if c.Bool("private") && strings.HasPrefix(n.NetworkType, "Private") {
				if c.Bool("mac") {
					fmt.Println(n.MacAddress)
					continue
				}
				for _, i := range n.IpAddressList {
					if c.Bool("v6") && i.IPv6.IpAddress != "" {
						fmt.Println(i.IPv6.IpAddress)
					} else if !c.Bool("v6") && i.IPv4.IpAddress != "" {
						fmt.Println(i.IPv4.IpAddress)
					}
				}
			} else if !c.Bool("private") && n.NetworkType == "Global" {
				if c.Bool("mac") {
					fmt.Println(n.MacAddress)
					continue
				}
				for _, i := range n.IpAddressList {
					if c.Bool("v6") && i.IPv6.IpAddress != "" {
						fmt.Println(i.IPv6.IpAddress)
					} else if !c.Bool("v6") && i.IPv4.IpAddress != "" {
						fmt.Println(i.IPv4.IpAddress)
					}
				}
			}
		}
	}
	return nil
}

func main() {
	app := cli.NewApp()
	app.Name = "ipof"
	app.Usage = "show ip address of vm"
	app.Author = ""
	app.Version = ""
	app.UsageText = "ipof [options] ivmXXXXXXXX..."
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:   "access-key,AccessKey,Access,access,a",
			EnvVar: "IIJAPI_ACCESS_KEY",
		},
		cli.StringFlag{
			Name:   "secret-key,SecretKey,Secret,secret,s",
			EnvVar: "IIJAPI_SECRET_KEY",
		},
		cli.StringFlag{
			Name:   "gis-service-code,GisServiceCode,Gis,gis,g",
			EnvVar: "GISSERVICECODE,GisServiceCode",
		},
		cli.BoolFlag{
			Name:  "private,p",
			Usage: "show private address",
		},
		cli.BoolFlag{
			Name:  "v6,ipv6,6",
			Usage: "show IPv6 address",
		},
		cli.BoolFlag{
			Name:  "mac,mac-address,m",
			Usage: "show MAC address",
		},
		cli.BoolFlag{
			Name: "verbose",
		},
	}
	app.Action = doMain
	app.Run(os.Args)
}

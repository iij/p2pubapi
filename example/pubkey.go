package main

// Github APIを使ってsshの公開鍵を取得し、P2PUBのVMに設定する
// Github Enterprizeの場合は--Endpoint https://<ghe host>/api/v3/ を指定する

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"path/filepath"

	"github.com/iij/p2pubapi"
	"github.com/iij/p2pubapi/protocol"

	log "github.com/Sirupsen/logrus"
	"github.com/urfave/cli"
)

type pubKeyAPI []struct {
	ID  int    `json:"id"`
	Key string `json:"key"`
}

func doMain(c *cli.Context) (err error) {
	if c.Bool("verbose") {
		log.SetLevel(log.DebugLevel)
	}
	api := p2pubapi.NewAPI(c.String("AccessKey"), c.String("SecretKey"))
	log.Debug("api", api)
	if len(api.AccessKey) == 0 || len(api.SecretKey) == 0 {
		log.Error("option required: AccessKey, SecretKey")
		return fmt.Errorf("option required: AccessKey, SecretKey")
	}

	// gh api
	var req *http.Request
	req, err = http.NewRequest("GET", c.String("Endpoint")+filepath.Join("users", c.String("Account"), "keys"), nil)
	if err != nil {
		log.Error(err)
		return
	}
	req.Header.Add("Accept", "application/vnd.github.v3+json")
	cl := http.Client{}
	var resp *http.Response
	resp, err = cl.Do(req)
	if err != nil {
		log.Error(err)
		return
	}
	dec := json.NewDecoder(resp.Body)
	var apires = pubKeyAPI{}
	if err = dec.Decode(&apires); err != nil {
		log.Error(err)
	}
	for _, k := range apires {
		log.Info("pubkey", k.ID, k.Key)
		arg := protocol.PublicKeyAdd{}
		arg.GisServiceCode = c.String("GisServiceCode")
		arg.IbaServiceCode = c.String("IbaServiceCode")
		arg.PublicKey = k.Key
		res := protocol.PublicKeyAddResponse{}
		if err = p2pubapi.Validate(arg); err != nil {
			log.Error(err)
			continue
		}
		if err = p2pubapi.Call(*api, arg, &res); err != nil {
			log.Error(err)
		} else {
			log.Infof("%s -> %s", res.Previous.ResourceStatus, res.Current.ResourceStatus)
		}
	}
	return
}

func main() {
	app := cli.NewApp()
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:   "AccessKey, a",
			Usage:  "Access Key",
			EnvVar: "IIJAPI_ACCESS_KEY",
		},
		cli.StringFlag{
			Name:   "SecretKey, s",
			Usage:  "Secret Key",
			EnvVar: "IIJAPI_SECRET_KEY",
		},
		cli.BoolFlag{
			Name:  "verbose, V",
			Usage: "verbose output",
		},
		cli.StringFlag{
			Name:   "GisServiceCode, g",
			Usage:  "gisXXXXXXXX",
			EnvVar: "GISSERVICECODE",
		},
		cli.StringFlag{
			Name:   "IbaServiceCode, i",
			Usage:  "ibaXXXXXXXX",
			EnvVar: "IBASERVICECODE",
		},
		cli.StringFlag{
			Name:   "Account, User, u",
			Usage:  "User Account for github.com",
			EnvVar: "USER",
		},
		cli.StringFlag{
			Name:  "Endpoint, e",
			Usage: "endpoint for github API",
			Value: "https://api.github.com/",
		},
	}
	app.Version = "0.0.1"
	app.Usage = "Public Key importer (github.com -> P2PUB)"
	app.Author = ""
	app.Action = doMain
	app.Run(os.Args)
}

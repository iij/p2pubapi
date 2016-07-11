package main

// 各種リスト系APIを使ってテーブル状に表示

import (
	"fmt"
	"os"
	"strings"
	"text/tabwriter"

	log "github.com/Sirupsen/logrus"

	"github.com/urfave/cli"

	"github.com/iij/p2pubapi"
	"github.com/iij/p2pubapi/protocol"
)

func newAPI(c *cli.Context) *p2pubapi.API {
	if c.GlobalBool("verbose") {
		log.SetLevel(log.DebugLevel)
	}
	return p2pubapi.NewAPI(c.GlobalString("AccessKey"), c.GlobalString("SecretKey"))
}

func lsgis(c *cli.Context) error {
	api := newAPI(c)
	arg := protocol.P2PUBContractServiceCodeListGetForSA{}
	arg.Item = "ServiceCode"
	if err := p2pubapi.Validate(arg); err != nil {
		log.Fatal(err)
		return err
	}
	res := protocol.P2PUBContractServiceCodeListGetForSAResponse{}
	err := p2pubapi.Call(*api, arg, &res)
	if err != nil {
		log.Fatal(err)
		return err
	}
	log.Println("res", res)
	w := new(tabwriter.Writer)
	w.Init(os.Stdout, 8, 0, 2, ' ', 0)
	fmt.Fprintln(w, "ServiceCode")
	for _, vm := range res.GisList {
		fmt.Fprintln(w, vm.ServiceCode)
	}
	w.Flush()
	return nil
}

func lsvm(c *cli.Context) error {
	api := newAPI(c)
	arg := protocol.VMListGet{}
	arg.GisServiceCode = c.String("GisServiceCode")
	if err := p2pubapi.Validate(arg); err != nil {
		log.Fatal(err)
		return err
	}
	res := protocol.VMListGetResponse{}
	err := p2pubapi.Call(*api, arg, &res)
	if err != nil {
		log.Fatal(err)
		return err
	}
	log.Println("res", res)
	w := new(tabwriter.Writer)
	w.Init(os.Stdout, 8, 0, 2, ' ', 0)
	fmt.Fprintln(w, "ServiceCode\tLabel\tFrom\tType\tOS\tStatus\tGlobal\tPrivate\tStorage")
	for _, vm := range res.VirtualServerList {
		x := []string{
			vm.ServiceCode, vm.Label, vm.StartDate, vm.Type, vm.OSType, vm.ResourceStatus,
		}
		var global = []string{}
		var private = []string{}
		var storage = []string{}
		for _, nw := range vm.NetworkList {
			var addrs = []string{}
			for _, ad := range nw.IpAddressList {
				addrs = append(addrs, ad.IPv4.IpAddress)
			}
			if nw.NetworkType == "Global" {
				global = append(global, addrs...)
			} else if strings.HasPrefix(nw.NetworkType, "Private") {
				private = append(private, addrs...)
			}
		}
		for _, st := range vm.StorageList {
			storage = append(storage, st.ServiceCode)
		}
		x = append(x, strings.Join(global, ", "), strings.Join(private, ", "), strings.Join(storage, ", "))
		fmt.Fprintln(w, strings.Join(x, "\t"))
	}
	w.Flush()

	return nil
}

func lsnw(c *cli.Context) error {
	api := newAPI(c)
	arg := protocol.GlobalAddressGetForGlobal{}
	arg.GisServiceCode = c.String("GisServiceCode")
	if err := p2pubapi.Validate(arg); err != nil {
		log.Fatal(err)
		return err
	}
	res := protocol.GlobalAddressGetForGlobalResponse{}
	err := p2pubapi.Call(*api, arg, &res)
	if err != nil {
		log.Fatal(err)
		return err
	}
	log.Println("res", res)
	adr := fmt.Sprintf("%s/%s %s", res.GlobalIpAddressNum.Assigned, res.GlobalIpAddressNum.Limit, res.GlobalIpAddressV.AddressNum)
	w := new(tabwriter.Writer)
	w.Init(os.Stdout, 8, 0, 2, ' ', 0)
	fmt.Fprintln(w, "VM\tType\tAddress "+adr+"\tDomainName")
	for _, nw := range res.AssignedResourceList {
		if len(nw.IPv4.Type) != 0 {
			fmt.Fprintf(w, "%s\t%s\t%s\t%s\n", nw.ServiceCode, nw.IPv4.Type, nw.IPv4.IpAddress, nw.IPv4.DomainName)
		}
		if len(nw.IPv6.Type) != 0 {
			fmt.Fprintf(w, "%s\t%s\t%s\t%s\n", nw.ServiceCode, nw.IPv6.Type, nw.IPv6.IpAddress, "")
		}
	}
	w.Flush()
	return nil
}

func lsip(c *cli.Context) error {
	api := newAPI(c)
	arg := protocol.GlobalAddressListGet{}
	arg.GisServiceCode = c.String("GisServiceCode")
	arg.IvmServiceCode = c.String("IvmServiceCode")

	if err := p2pubapi.Validate(arg); err != nil {
		log.Fatal(err)
		return err
	}
	res := protocol.GlobalAddressListGetResponse{}
	err := p2pubapi.Call(*api, arg, &res)
	if err != nil {
		log.Fatal(err)
		return err
	}
	log.Println("res", res)
	w := new(tabwriter.Writer)
	w.Init(os.Stdout, 8, 0, 2, ' ', 0)
	fmt.Fprintln(w, "Type\tAddress\tDomainName")
	for _, nw := range res.IpAddressList {
		if len(nw.IPv4.Type) != 0 {
			fmt.Fprintf(w, "%s\t%s\t%s\n", nw.IPv4.Type, nw.IPv4.IpAddress, nw.IPv4.DomainName)
		}
		if len(nw.IPv6.Type) != 0 {
			fmt.Fprintf(w, "%s\t%s\t%s\n", nw.IPv6.Type, nw.IPv6.IpAddress, "")
		}
	}
	w.Flush()
	return nil
}

func lspipv(c *cli.Context) error {
	api := newAPI(c)
	arg := protocol.PrivateNetworkListGet{}
	arg.GisServiceCode = c.String("GisServiceCode")
	arg.IvmServiceCode = c.String("IvmServiceCode")

	if err := p2pubapi.Validate(arg); err != nil {
		log.Fatal(err)
		return err
	}
	res := protocol.PrivateNetworkListGetResponse{}
	err := p2pubapi.Call(*api, arg, &res)
	if err != nil {
		log.Fatal(err)
		return err
	}
	log.Println("res", res)
	w := new(tabwriter.Writer)
	w.Init(os.Stdout, 8, 0, 2, ' ', 0)
	fmt.Fprintln(w, "ServiceCode\tLabel\tMacAddress\tURI")
	for _, nw := range res.NetworkList {
		x := []string{
			nw.ServiceCode, nw.Label, nw.MacAddress, nw.URI,
		}
		fmt.Fprintln(w, strings.Join(x, "\t"))
	}
	w.Flush()
	return nil
}

func lspip(c *cli.Context) error {
	api := newAPI(c)
	arg := protocol.PrivateNetworkVListGet{}
	arg.GisServiceCode = c.String("GisServiceCode")

	if err := p2pubapi.Validate(arg); err != nil {
		log.Fatal(err)
		return err
	}
	res := protocol.PrivateNetworkVListGetResponse{}
	err := p2pubapi.Call(*api, arg, &res)
	if err != nil {
		log.Fatal(err)
		return err
	}
	log.Println("res", res)
	w := new(tabwriter.Writer)
	w.Init(os.Stdout, 8, 0, 2, ' ', 0)
	fmt.Fprintln(w, "ServiceCode\tLabel\tStartAt\tAddress\tStatus")
	for _, nw := range res.PrivateNetworkList {
		x := []string{
			nw.ServiceCode, nw.Label, nw.StartDate, nw.NetworkAddress,
			nw.ContractStatus,
		}
		fmt.Fprintln(w, strings.Join(x, "\t"))
	}
	w.Flush()
	return nil
}

func lsst(c *cli.Context) error {
	api := newAPI(c)
	arg := protocol.StorageListGet{}
	arg.GisServiceCode = c.String("GisServiceCode")

	if err := p2pubapi.Validate(arg); err != nil {
		log.Println(arg.Document())
		log.Fatal(err)
		return err
	}
	res := protocol.StorageListGetResponse{}
	err := p2pubapi.Call(*api, arg, &res)
	if err != nil {
		log.Fatal(err)
		return err
	}
	log.Println("res", res)
	w := new(tabwriter.Writer)
	w.Init(os.Stdout, 8, 0, 2, ' ', 0)
	fmt.Fprintln(w, "ServiceCode\tLabel\tSize\tStartAt\tStatus")
	for _, st := range res.AdditionalStorageList {
		x := []string{
			st.ServiceCode, st.Label, st.StorageSize, st.StartDate, st.ResourceStatus,
		}
		fmt.Fprintln(w, strings.Join(x, "\t"))
	}
	w.Flush()
	return nil
}

func lssysst(c *cli.Context) error {
	api := newAPI(c)
	arg := protocol.SystemStorageListGet{}
	arg.GisServiceCode = c.String("GisServiceCode")

	if err := p2pubapi.Validate(arg); err != nil {
		log.Println(arg.Document())
		log.Fatal(err)
		return err
	}
	res := protocol.SystemStorageListGetResponse{}
	err := p2pubapi.Call(*api, arg, &res)
	if err != nil {
		log.Fatal(err)
		return err
	}
	log.Println("res", res)
	w := new(tabwriter.Writer)
	w.Init(os.Stdout, 8, 0, 2, ' ', 0)
	fmt.Fprintln(w, "ServiceCode\tLabel\tSize\tStartAt\tStatus\tAttachedTo\tType\tOS")
	for _, st := range res.SystemStorageList {
		x := []string{
			st.ServiceCode, st.Label, st.StorageSize, st.StartDate, st.ContractStatus + "/" + st.ResourceStatus,
			st.AttachedVirtualServer.ServiceCode, st.AttachedVirtualServer.Type, st.AttachedVirtualServer.OSType,
		}
		fmt.Fprintln(w, strings.Join(x, "\t"))
	}
	w.Flush()
	return nil
}

func lssc(c *cli.Context) error {
	api := newAPI(c)
	arg := protocol.P2PUBContractGetForSA{}
	arg.GisServiceCode = c.String("GisServiceCode")
	var res = protocol.P2PUBContractGetForSAResponse{}
	err := p2pubapi.Call(*api, arg, &res)
	if err != nil {
		log.Fatal(err)
		return err
	}
	log.Printf("res: %+v", res)
	w := new(tabwriter.Writer)
	w.Init(os.Stdout, 8, 0, 2, ' ', 0)
	fmt.Fprintln(w, "Type\tServiceCode\tnote")
	for _, c := range res.VirtualServerList {
		fmt.Fprintln(w, strings.Join([]string{"VM", c.ServiceCode, c.Category}, "\t"))
	}
	for _, c := range res.PrivateNetworkList {
		fmt.Fprintln(w, strings.Join([]string{"PrivateNet", c.ServiceCode}, "\t"))
	}
	for _, c := range res.SystemStorageList {
		fmt.Fprintln(w, strings.Join([]string{"SystemStorage", c.ServiceCode}, "\t"))
	}
	for _, c := range res.AdditionalStorageList {
		fmt.Fprintln(w, strings.Join([]string{"AddStorage", c.ServiceCode, c.Category}, "\t"))
	}
	if res.StorageArchive.ServiceCode != "" {
		fmt.Fprintln(w, strings.Join([]string{"Archive", res.StorageArchive.ServiceCode}, "\t"))
	}
	if res.GlobalIpAddress.ServiceCode != "" {
		fmt.Fprintln(w, strings.Join([]string{"GlobalAddress", res.GlobalIpAddress.ServiceCode}, "\t"))
	}
	for _, c := range res.FwLbList {
		fmt.Fprintln(w, strings.Join([]string{"FwLb", c.ServiceCode}, "\t"))
	}
	w.Flush()
	return nil
}

func lsarchive(c *cli.Context) error {
	api := newAPI(c)
	arg := protocol.CustomOSImageListGet{
		GisServiceCode: c.String("GisServiceCode"),
		IarServiceCode: c.String("IarServiceCode"),
	}
	var res = protocol.CustomOSImageListGetResponse{}
	if err := p2pubapi.Call(*api, arg, &res); err != nil {
		log.Fatal(err)
		return err
	}
	w := new(tabwriter.Writer)
	w.Init(os.Stdout, 8, 0, 2, ' ', 0)
	fmt.Fprintln(w, "Type\tStatus\tImageId\tSize\tParent\tOS\tDateTime")
	for _, i := range res.ImageList {
		fmt.Fprintf(w, "%s\t%s\t%s\t%s\t%s\t%s\t%s\n", i.Type, i.Status, i.ImageId, i.ImageSize, i.SrcServiceCode, i.OSType, i.ArchivedDateTime)
	}
	w.Flush()
	return nil
}

func main() {
	app := cli.NewApp()
	app.EnableBashCompletion = true
	gisflag := cli.StringFlag{
		Name:   "GisServiceCode, g",
		Usage:  "gisXXXXXXXX",
		EnvVar: "GISSERVICECODE",
	}
	ivmflag := cli.StringFlag{
		Name:   "IvmServiceCode, i",
		Usage:  "ivmXXXXXXXX",
		EnvVar: "IVMSERVICECODE",
	}
	iarflag := cli.StringFlag{
		Name:   "IarServiceCode, a",
		Usage:  "iarXXXXXXXX",
		EnvVar: "IARSERVICECODE",
	}
	app.Commands = []cli.Command{
		{
			Name:   "gis",
			Action: lsgis,
		},
		{
			Name:    "vm",
			Aliases: []string{"virtualmachine", "virtualserver", "vms"},
			Usage:   "show VM list",
			Action:  lsvm,
			Flags:   []cli.Flag{gisflag},
		},
		{
			Name:    "net",
			Aliases: []string{"network", "nw"},
			Usage:   "show network list",
			Action:  lsnw,
			Flags:   []cli.Flag{gisflag},
		},
		{
			Name:    "global",
			Aliases: []string{"globaladdress", "ip"},
			Usage:   "show global ip address list",
			Action:  lsip,
			Flags:   []cli.Flag{gisflag, ivmflag},
		},
		{
			Name:    "private",
			Aliases: []string{"privateaddress", "pip"},
			Usage:   "show private ip address list",
			Action:  lspip,
			Flags:   []cli.Flag{gisflag},
		},
		{
			Name:    "vmprivate",
			Aliases: []string{"vmprivateaddress", "pipv", "vpip"},
			Usage:   "show private ip address list for vm",
			Action:  lspipv,
			Flags:   []cli.Flag{gisflag, ivmflag},
		},
		{
			Name:    "storage",
			Aliases: []string{"st", "disk"},
			Usage:   "show storage list",
			Action:  lsst,
			Flags:   []cli.Flag{gisflag},
		},
		{
			Name:    "systemstorage",
			Aliases: []string{"sst", "sdisk"},
			Usage:   "show systemstorage list",
			Action:  lssysst,
			Flags:   []cli.Flag{gisflag},
		},
		{
			Name:    "servicecode",
			Aliases: []string{"sc"},
			Usage:   "show servicecode list",
			Action:  lssc,
			Flags:   []cli.Flag{gisflag},
		},
		{
			Name:    "archive",
			Aliases: []string{"arc", "iar"},
			Usage:   "show archive storage list",
			Action:  lsarchive,
			Flags:   []cli.Flag{gisflag, iarflag},
		},
	}
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
	}
	app.Version = "0.0.1"
	app.Usage = "list P2PUB"
	app.Author = ""
	log.SetOutput(os.Stderr)
	form := log.TextFormatter{}
	form.FullTimestamp = true
	log.SetFormatter(&form)
	log.SetLevel(log.WarnLevel)

	app.Run(os.Args)
}

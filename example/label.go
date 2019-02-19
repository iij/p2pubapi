package main

import (
	"fmt"
	"log"
	"os"

	"github.com/iij/p2pubapi"
	"github.com/iij/p2pubapi/protocol"
)

func usage() {
	fmt.Println("Usage: ", os.Args[0], "serviceCode [newLabel]")
}

var gis = os.Getenv("GISSERVICECODE")
var accesskey = os.Getenv("IIJAPI_ACCESS_KEY")
var secretkey = os.Getenv("IIJAPI_SECRET_KEY")

func newapi() *p2pubapi.API {
	return p2pubapi.NewAPI(accesskey, secretkey)
}

func showIvm(ivm string) error {
	arg := protocol.VMGet{
		GisServiceCode: gis,
		IvmServiceCode: ivm,
	}
	var res = protocol.VMGetResponse{}
	api := newapi()
	if err := p2pubapi.Call(*api, arg, &res); err != nil {
		return err
	}
	log.Println(ivm, res.Label)
	return nil
}

func setIvm(ivm, newlabel string) error {
	arg := protocol.VMLabelSet{
		GisServiceCode: gis,
		IvmServiceCode: ivm,
		Name:           newlabel,
	}
	var res = protocol.VMLabelSetResponse{}
	api := newapi()
	if err := p2pubapi.Call(*api, arg, &res); err != nil {
		return err
	}
	log.Println(ivm, res)
	return nil
}

func showSystemStorage(storage string) error {
	arg := protocol.SystemStorageGet{
		GisServiceCode: gis,
		StorageServiceCode: storage,
	}
	var res = protocol.SystemStorageGetResponse{}
	api := newapi()
	if err := p2pubapi.Call(*api, arg, &res); err != nil {
		return err
	}
	log.Println(storage, res.Label)
	return nil
}

func setSystemStorage(storage, newlabel string) error {
	arg := protocol.SystemStorageLabelSet{
		GisServiceCode: gis,
		StorageServiceCode: storage,
		Name:           newlabel,
	}
	var res = protocol.SystemStorageLabelSetResponse{}
	api := newapi()
	if err := p2pubapi.Call(*api, arg, &res); err != nil {
		return err
	}
	log.Println(storage, res)
	return nil
}

func showIar(iar string) error {
	arg := protocol.CustomOSImageListGet{
		GisServiceCode: gis,
		IarServiceCode: iar,
	}
	var res = protocol.CustomOSImageListGetResponse{}
	api := newapi()
	if err := p2pubapi.Call(*api, arg, &res); err != nil {
		return err
	}
	for _, v := range res.ImageList {
		log.Println(iar, v.ImageId, v.Label)
	}
	return nil
}

func setIar(iar, imageid, newlabel string) error {
	arg := protocol.CustomOSImageLabelSet{
		GisServiceCode: gis,
		IarServiceCode: iar,
		ImageId:        imageid,
		Name:           newlabel,
	}
	var res = protocol.CustomOSImageLabelSetResponse{}
	api := newapi()
	if err := p2pubapi.Call(*api, arg, &res); err != nil {
		return err
	}
	log.Println(iar, imageid, res)
	return nil
}

func show(sc string) {
	switch sc[0:3] {
	case "ivm":
		showIvm(sc)
	case "iba", "ica":
		showSystemStorage(sc)
	case "iar":
		showIar(sc)
	default:
		log.Println("sc", sc[0:3])
	}
}

func set(sc, val1, val2 string) {
	switch sc[0:3] {
	case "ivm":
		setIvm(sc, val1)
	case "iba", "ica":
		setSystemStorage(sc, val1)
	case "iar":
		setIar(sc, val1, val2)
	default:
		log.Println("sc", sc[0:3])
	}
}

// Usage: $0 serviceCode [newLabel]
func main() {
	switch len(os.Args) {
	case 1:
		usage()
	case 2:
		show(os.Args[1])
	case 3:
		set(os.Args[1], os.Args[2], "")
	case 4:
		set(os.Args[1], os.Args[2], os.Args[3])
	}
}

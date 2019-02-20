package main

// Usage:
// export IIJAPI_ACCESS_KEY=<your access key>
// export IIJAPI_SECRET_KEY=<your secret key>
// export GISSERVICECODE=<your gis code>
// $0 servicecode...

import (
	"fmt"
	"log"
	"os"

	"github.com/iij/p2pubapi"
	"github.com/iij/p2pubapi/protocol"
)

func cancelVM(api *p2pubapi.API, gis, sc string) error {
	arg := protocol.VMCancel{
		GisServiceCode: gis,
		IvmServiceCode: sc,
	}
	var res = protocol.VMCancelResponse{}
	if err := p2pubapi.Call(*api, arg, &res); err != nil {
		return err
	}
	return nil
}

func cancelSystemStorage(api *p2pubapi.API, gis, sc string) error {
	arg := protocol.SystemStorageCancel{
		GisServiceCode: gis,
		StorageServiceCode: sc,
	}
	var res = protocol.SystemStorageCancelResponse{}
	if err := p2pubapi.Call(*api, arg, &res); err != nil {
		return err
	}
	return nil
}

func cancelDataStorage(api *p2pubapi.API, gis, sc string) error {
	arg := protocol.StorageCancel{
		GisServiceCode: gis,
		StorageServiceCode: sc,
	}
	var res = protocol.StorageCancelResponse{}
	if err := p2pubapi.Call(*api, arg, &res); err != nil {
		return err
	}
	return nil
}

func cancelArchiveStorage(api *p2pubapi.API, gis, sc string) error {
	arg := protocol.StorageArchiveCancel{
		GisServiceCode: gis,
		IarServiceCode: sc,
	}
	var res = protocol.StorageArchiveCancelResponse{}
	if err := p2pubapi.Call(*api, arg, &res); err != nil {
		return err
	}
	return nil
}

func cancelGlobalAddress(api *p2pubapi.API, gis, sc string) error {
	arg := protocol.GlobalAddressVCancel{
		GisServiceCode: gis,
		IgaServiceCode: sc,
	}
	var res = protocol.GlobalAddressVCancelResponse{}
	if err := p2pubapi.Call(*api, arg, &res); err != nil {
		return err
	}
	return nil
}

func cancelPrivateNetwork(api *p2pubapi.API, gis, sc string) error {
	arg := protocol.PrivateNetworkVCancel{
		GisServiceCode: gis,
		IvlServiceCode: sc,
	}
	var res = protocol.PrivateNetworkVCancelResponse{}
	if err := p2pubapi.Call(*api, arg, &res); err != nil {
		return err
	}
	return nil
}

func cancelFwLb(api *p2pubapi.API, gis, sc string) error {
	arg := protocol.FwLbCancel{
		GisServiceCode: gis,
		IflServiceCode: sc,
	}
	var res = protocol.FwLbCancelResponse{}
	if err := p2pubapi.Call(*api, arg, &res); err != nil {
		return err
	}
	return nil
}

func main() {
	api := p2pubapi.NewAPI(os.Getenv("IIJAPI_ACCESS_KEY"), os.Getenv("IIJAPI_SECRET_KEY"))
	gis := os.Getenv("GISSERVICECODE")
	for _, v := range os.Args[1:len(os.Args)] {
		var err error
		log.Println("try cancel", v)
		switch v[0:3] {
		default:
			err = fmt.Errorf("unknown service code prefix: %s (%s)", v[0:3], v)
		case "ivm", "ivd":
			err = cancelVM(api, gis, v)
		case "iba", "ica":
			err = cancelSystemStorage(api, gis, v)
		case "ibg", "ibb", "icg", "icb":
			err = cancelDataStorage(api, gis, v)
		case "iar":
			err = cancelArchiveStorage(api, gis, v)
		case "iga":
			err = cancelGlobalAddress(api, gis, v)
		case "ivl":
			err = cancelPrivateNetwork(api, gis, v)
		case "ifl":
			err = cancelFwLb(api, gis, v)
		}
		if err != nil {
			log.Println(err)
		}
	}
}

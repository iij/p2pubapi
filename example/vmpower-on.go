package main

// Usage:
//   export IIJAPI_ACCESS_KEY=<YOUR ACCESSS KEY>
//   export IIJAPI_SECRET_KEY=<YOUR SECRET KEY>
//   export GISSERVICECODE=<YOUR GIS CODE>
//   $0

import (
	"log"
	"os"

	"github.com/iij/p2pubapi"
	"github.com/iij/p2pubapi/protocol"
)

func main() {
	api := p2pubapi.NewAPI(os.Getenv("IIJAPI_ACCESS_KEY"), os.Getenv("IIJAPI_SECRET_KEY"))
	// list VMs
	listarg := protocol.VMListGet{}
	listarg.GisServiceCode = os.Getenv("GISSERVICECODE")
	var listresp = protocol.VMListGetResponse{}
	if err := p2pubapi.Call(*api, listarg, &listresp); err != nil {
		log.Println("List API error", err)
	}

	// Power On Stopped VMs
	arg := protocol.VMPower{}
	arg.GisServiceCode = os.Getenv("GISSERVICECODE")
	arg.Power = "On"
	for _, v := range listresp.VirtualServerList {
		if v.ResourceStatus == "Stopped" {
			log.Println("Power On VM:", v.ServiceCode, v.OSType)
			arg.IvmServiceCode = v.ServiceCode
			var resp = protocol.VMPowerResponse{}
			if err := p2pubapi.Call(*api, arg, &resp); err != nil {
				log.Println("API error", err)
			}
			log.Printf("%+v", resp)
		}
	}
}

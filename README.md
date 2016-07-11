# Golang binding for P2PUB API

## Install

- go get github.com/iij/p2pubapi
- go build -o p2pub github.com/iij/p2pubapi/cli
- install -c p2pub /usr/local/bin/p2pub

...or [download binary](https://github.com/iij/p2pubapi/releases)

# Use p2pub command

```
# vi setup.sh
IIJAPI_ACCESS_KEY=<YOUR ACCESS KEY>
IIJAPI_SECRET_KEY=<YOUR SECRET KEY>
GISSERVICECODE=<YOUR GIS SERVICE CODE>
# . ./setup.sh
# p2pub VMListGet
{... (result json)}
# p2pub --format yaml VMListGet
(result YAML)
```

## Example

- list your GIS ServiceCode
    - p2pub P2PUBContractServiceCodeListGetForSA --Item ServiceCode | jq -r '.GisList[].ServiceCode'
- list your VM's IP Address
    - p2pub VMListGet | jq -r '.VirtualServerList[].NetworkList[].IpAddressList[].IPv4.IpAddress'

# Usage for Golang users

[API reference](http://godoc.org/github.com/iij/p2pubapi)

```go
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
```

## Example

- [vmpower-on.go](example/vmpower-on.go)
    - Power On all VMs
- [ls.go](example/ls.go)
    - Call ListGet API and show result as Text Table
- [pubkey.go](example/pubkey.go)
    - SSH Public Key importer
    - get public key from github API
    - set public key to System Storage with P2PUB API
- [label.go](example/label.go)
    - set/get labels

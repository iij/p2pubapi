package p2pubapi

import (
	"fmt"
	"time"

	"github.com/iij/p2pubapi/protocol"
)

const (
	pollInterval = time.Duration(10 * time.Second)
)

// WaitVM wait vm status (contract status, resource status)
// Contract Status(cstatus): InPreparation/InService
// Resource Status(rstatus): Stopped/Configuring/Starting/Running/Stopping/Locked
func WaitVM(api *API, gis, ivm string, cstatus, rstatus Status, maxwait time.Duration) error {
	start := time.Now()
	for {
		arg := protocol.VMGet{
			GisServiceCode: gis,
			IvmServiceCode: ivm,
		}
		var res = protocol.VMGetResponse{}
		if err := Call(*api, arg, &res); err != nil {
			return err
		}
		if (cstatus == None || cstatus.String() == res.ContractStatus) &&
			(rstatus == None || rstatus.String() == res.ResourceStatus) {
			break
		}
		if time.Since(start) > maxwait {
			return fmt.Errorf("timeout")
		}
		time.Sleep(pollInterval)
	}
	return nil
}

// WaitSystemStorage wait iba status (contract status, resource status)
// Contract Status(cstatus): InPreparation/InService
// Resource Status(rstatus): Attached/NotAttached/Initializing/Configuring/Archiving
func WaitSystemStorage(api *API, gis, iba string, cstatus, rstatus Status, maxwait time.Duration) error {
	start := time.Now()
	for {
		arg := protocol.SystemStorageGet{
			GisServiceCode: gis,
			IbaServiceCode: iba,
		}
		var res = protocol.SystemStorageGetResponse{}
		if err := Call(*api, arg, &res); err != nil {
			return err
		}
		if (cstatus == None || cstatus.String() == res.ContractStatus) &&
			(rstatus == None || rstatus.String() == res.ResourceStatus) {
			break
		}
		if time.Since(start) > maxwait {
			return fmt.Errorf("timeout")
		}
		time.Sleep(pollInterval)
	}
	return nil
}

// WaitDataStorage wait ibg status (contract status, resource status)
// Contract Status(cstatus): InPreparation/InService
// Resource Status(rstatus): Attached/NotAttached/Initializing/Configuring/Archiving
func WaitDataStorage(api *API, gis, ibg string, cstatus, rstatus Status, maxwait time.Duration) error {
	start := time.Now()
	for {
		arg := protocol.StorageGet{
			GisServiceCode: gis,
			IbgServiceCode: ibg,
		}
		var res = protocol.StorageGetResponse{}
		if err := Call(*api, arg, &res); err != nil {
			return err
		}
		if (cstatus == None || cstatus.String() == res.ContractStatus) &&
			(rstatus == None || rstatus.String() == res.ResourceStatus) {
			break
		}
		if time.Since(start) > maxwait {
			return fmt.Errorf("timeout")
		}
		time.Sleep(pollInterval)
	}
	return nil
}

// WaitArchiveStorage wait iar status (contract status)
// Contract Status(cstatus): InPreparation/InService
func WaitArchiveStorage(api *API, gis, iar string, cstatus Status, maxwait time.Duration) error {
	start := time.Now()
	for {
		arg := protocol.StorageArchiveGet{
			GisServiceCode: gis,
			IarServiceCode: iar,
		}
		var res = protocol.StorageArchiveGetResponse{}
		if err := Call(*api, arg, &res); err != nil {
			return err
		}
		if cstatus == None || cstatus.String() == res.ContractStatus {
			break
		}
		if time.Since(start) > maxwait {
			return fmt.Errorf("timeout")
		}
		time.Sleep(pollInterval)
	}
	return nil
}

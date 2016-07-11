package protocol

import (
	"reflect"
)

// DataDeviceStorageDisconnect データデバイスストレージ切断 (同期)
//  http://manual.iij.jp/p2/pubapi/59939612.html
type DataDeviceStorageDisconnect struct {
	GisServiceCode string `json:"-"` // P2契約のサービスコード(gis########)
	IvmServiceCode string `json:"-"` // 仮想サーバのサービスコード(ivm########, ivd########)
	PciSlot        string `json:"-"` // PCIスロット(16進数)
}

// URI /{{.GisServiceCode}}/virtual-servers/{{.IvmServiceCode}}/data-devices/{{.PciSlot}}.json
func (t DataDeviceStorageDisconnect) URI() string {
	return "/{{.GisServiceCode}}/virtual-servers/{{.IvmServiceCode}}/data-devices/{{.PciSlot}}.json"
}

// APIName DataDeviceStorageDisconnect
func (t DataDeviceStorageDisconnect) APIName() string {
	return "DataDeviceStorageDisconnect"
}

// Method DELETE
func (t DataDeviceStorageDisconnect) Method() string {
	return "DELETE"
}

// Document http://manual.iij.jp/p2/pubapi/59939612.html
func (t DataDeviceStorageDisconnect) Document() string {
	return "http://manual.iij.jp/p2/pubapi/59939612.html"
}

// JPName データデバイスストレージ切断
func (t DataDeviceStorageDisconnect) JPName() string {
	return "データデバイスストレージ切断"
}
func init() {
	APIlist = append(APIlist, DataDeviceStorageDisconnect{})
	TypeMap["DataDeviceStorageDisconnect"] = reflect.TypeOf(DataDeviceStorageDisconnect{})
}

// DataDeviceStorageDisconnectResponse データデバイスストレージ切断のレスポンス
type DataDeviceStorageDisconnectResponse struct {
	*CommonResponse
	OSType      string `json:",omitempty"` // OS種別(Windows/Linux)
	URI         string `json:",omitempty"` // ストレージの情報へアクセスするためのURI(URI)
	Category    string `json:",omitempty"` // ストレージ種別(BestEffort/Gurantee)
	ServiceCode string `json:",omitempty"` // ストレージのサービスコード(iba########)
	Type        string `json:",omitempty"` // ストレージ品目
}

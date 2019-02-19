package protocol

import (
	"reflect"
)

// DataDeviceStorageGet データデバイスストレージ情報取得 (同期)
//  http://manual.iij.jp/p2/pubapi/59939603.html
type DataDeviceStorageGet struct {
	GisServiceCode string `json:"-"` // P2契約のサービスコード(gis########)
	IvmServiceCode string `json:"-"` // 仮想サーバのサービスコード(ivm########, ivd########)
	PciSlot        string `json:"-"` // PCIスロット(16進数)
}

// URI /{{.GisServiceCode}}/virtual-servers/{{.IvmServiceCode}}/data-devices/{{.PciSlot}}.json
func (t DataDeviceStorageGet) URI() string {
	return "/{{.GisServiceCode}}/virtual-servers/{{.IvmServiceCode}}/data-devices/{{.PciSlot}}.json"
}

// APIName DataDeviceStorageGet
func (t DataDeviceStorageGet) APIName() string {
	return "DataDeviceStorageGet"
}

// Method GET
func (t DataDeviceStorageGet) Method() string {
	return "GET"
}

// Document http://manual.iij.jp/p2/pubapi/59939603.html
func (t DataDeviceStorageGet) Document() string {
	return "http://manual.iij.jp/p2/pubapi/59939603.html"
}

// JPName データデバイスストレージ情報取得
func (t DataDeviceStorageGet) JPName() string {
	return "データデバイスストレージ情報取得"
}
func init() {
	APIlist = append(APIlist, DataDeviceStorageGet{})
	TypeMap["DataDeviceStorageGet"] = reflect.TypeOf(DataDeviceStorageGet{})
}

// DataDeviceStorageGetResponse データデバイスストレージ情報取得のレスポンス
type DataDeviceStorageGetResponse struct {
	*CommonResponse
	OSType      string `json:",omitempty"` // OS種別(None/Windows/Linux)
	URI         string `json:",omitempty"` // ストレージの情報へアクセスするためのURI(URI)
	Category    string `json:",omitempty"` // ストレージ種別(BestEffort/Guarantee)
	ServiceCode string `json:",omitempty"` // ストレージのサービスコード(iba/ibb/ibg/ica/icb/icg########)
	Type        string `json:",omitempty"` // ストレージ品目
}

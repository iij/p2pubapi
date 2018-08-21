package protocol

import (
	"reflect"
)

// BootDeviceStorageConnect ブートデバイスストレージ接続 (同期)
//  http://manual.iij.jp/p2/pubapi/59939555.html
type BootDeviceStorageConnect struct {
	GisServiceCode string `json:"-"` // P2契約のサービスコード(gis########)
	IvmServiceCode string `json:"-"` // 仮想サーバのサービスコード(ivm########, ivd########)
	IbaServiceCode string // システムストレージタイプ S のサービスコード(iba########)
	IcaServiceCode string // システムストレージタイプ X のサービスコード(ica########)
}

// URI /{{.GisServiceCode}}/virtual-servers/{{.IvmServiceCode}}/boot-device.json
func (t BootDeviceStorageConnect) URI() string {
	return "/{{.GisServiceCode}}/virtual-servers/{{.IvmServiceCode}}/boot-device.json"
}

// APIName BootDeviceStorageConnect
func (t BootDeviceStorageConnect) APIName() string {
	return "BootDeviceStorageConnect"
}

// Method PUT
func (t BootDeviceStorageConnect) Method() string {
	return "PUT"
}

// Document http://manual.iij.jp/p2/pubapi/59939555.html
func (t BootDeviceStorageConnect) Document() string {
	return "http://manual.iij.jp/p2/pubapi/59939555.html"
}

// JPName ブートデバイスストレージ接続
func (t BootDeviceStorageConnect) JPName() string {
	return "ブートデバイスストレージ接続"
}
func init() {
	APIlist = append(APIlist, BootDeviceStorageConnect{})
	TypeMap["BootDeviceStorageConnect"] = reflect.TypeOf(BootDeviceStorageConnect{})
}

// BootDeviceStorageConnectResponse ブートデバイスストレージ接続のレスポンス
type BootDeviceStorageConnectResponse struct {
	*CommonResponse
	OSType      string `json:",omitempty"` // OS種別(Windows/Linux)
	URI         string `json:",omitempty"` // ストレージの情報へアクセスするためのURI(URI)
	PciSlot     string `json:",omitempty"` // PCIスロット(16進数)
	ServiceCode string `json:",omitempty"` // ストレージのサービスコード(iba########)
	Type        string `json:",omitempty"` // ストレージ品目
}

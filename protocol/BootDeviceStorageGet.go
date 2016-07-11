package protocol

import (
	"reflect"
)

// BootDeviceStorageGet ブートデバイスストレージ情報取得 (同期)
//  http://manual.iij.jp/p2/pubapi/59939569.html
type BootDeviceStorageGet struct {
	GisServiceCode string `json:"-"` // P2契約のサービスコード(gis########)
	IvmServiceCode string `json:"-"` // 仮想サーバのサービスコード(ivm########, ivd########)
}

// URI /{{.GisServiceCode}}/virtual-servers/{{.IvmServiceCode}}/boot-device.json
func (t BootDeviceStorageGet) URI() string {
	return "/{{.GisServiceCode}}/virtual-servers/{{.IvmServiceCode}}/boot-device.json"
}

// APIName BootDeviceStorageGet
func (t BootDeviceStorageGet) APIName() string {
	return "BootDeviceStorageGet"
}

// Method GET
func (t BootDeviceStorageGet) Method() string {
	return "GET"
}

// Document http://manual.iij.jp/p2/pubapi/59939569.html
func (t BootDeviceStorageGet) Document() string {
	return "http://manual.iij.jp/p2/pubapi/59939569.html"
}

// JPName ブートデバイスストレージ情報取得
func (t BootDeviceStorageGet) JPName() string {
	return "ブートデバイスストレージ情報取得"
}
func init() {
	APIlist = append(APIlist, BootDeviceStorageGet{})
	TypeMap["BootDeviceStorageGet"] = reflect.TypeOf(BootDeviceStorageGet{})
}

// BootDeviceStorageGetResponse ブートデバイスストレージ情報取得のレスポンス
type BootDeviceStorageGetResponse struct {
	*CommonResponse
	OSType      string `json:",omitempty"` // OS種別(Windows/Linux)
	URI         string `json:",omitempty"` // ストレージの情報へアクセスするためのURI(URI)
	PciSlot     string `json:",omitempty"` // PCIスロット(16進数)
	ServiceCode string `json:",omitempty"` // ストレージのサービスコード(iba########)
	Type        string `json:",omitempty"` // ストレージ 品目
}

package protocol

import (
	"reflect"
)

// BootDeviceStorageDisconnect ブートデバイスストレージ切断 (同期)
//  http://manual.iij.jp/p2/pubapi/59939581.html
type BootDeviceStorageDisconnect struct {
	GisServiceCode string `json:"-"` // P2契約のサービスコード(gis########)
	IvmServiceCode string `json:"-"` // 仮想サーバのサービスコード(ivm########, ivd########)
}

// URI /{{.GisServiceCode}}/virtual-servers/{{.IvmServiceCode}}/boot-device.json
func (t BootDeviceStorageDisconnect) URI() string {
	return "/{{.GisServiceCode}}/virtual-servers/{{.IvmServiceCode}}/boot-device.json"
}

// APIName BootDeviceStorageDisconnect
func (t BootDeviceStorageDisconnect) APIName() string {
	return "BootDeviceStorageDisconnect"
}

// Method DELETE
func (t BootDeviceStorageDisconnect) Method() string {
	return "DELETE"
}

// Document http://manual.iij.jp/p2/pubapi/59939581.html
func (t BootDeviceStorageDisconnect) Document() string {
	return "http://manual.iij.jp/p2/pubapi/59939581.html"
}

// JPName ブートデバイスストレージ切断
func (t BootDeviceStorageDisconnect) JPName() string {
	return "ブートデバイスストレージ切断"
}
func init() {
	APIlist = append(APIlist, BootDeviceStorageDisconnect{})
	TypeMap["BootDeviceStorageDisconnect"] = reflect.TypeOf(BootDeviceStorageDisconnect{})
}

// BootDeviceStorageDisconnectResponse ブートデバイスストレージ切断のレスポンス
type BootDeviceStorageDisconnectResponse struct {
	*CommonResponse
	Current struct {
		URI         string `json:",omitempty"` // 空文字列(文字列)
		ServiceCode string `json:",omitempty"` // 空文字列(文字列)
	} `json:",omitempty"`
	Previous struct {
		URI         string `json:",omitempty"` // 切り離し前のシステムストレージの情報へアクセスするURI(URI)
		ServiceCode string `json:",omitempty"` // 切り離し前のシステムストレージのサービスコード(iba########)
	} `json:",omitempty"`
	PciSlot string `json:",omitempty"` // PCIスロット(16進数)
}

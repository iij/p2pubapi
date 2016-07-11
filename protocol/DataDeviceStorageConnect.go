package protocol

import (
	"reflect"
)

// DataDeviceStorageConnect データデバイスストレージ接続 (同期)
//  http://manual.iij.jp/p2/pubapi/59939590.html
type DataDeviceStorageConnect struct {
	GisServiceCode string `json:"-"`          // P2契約のサービスコード(gis########)
	IvmServiceCode string `json:"-"`          // 仮想サーバのサービスコード(ivm########, ivd########)
	IbgServiceCode string `json:",omitempty"` // 追加ストレージ性能保証タイプのサービスコード(ibg########)
	IbbServiceCode string `json:",omitempty"` // 追加ストレージベストエフォートタイプのサービスコード(ibb########)
	IbaServiceCode string `json:",omitempty"` // システムストレージのサービスコード(iba########)
}

// URI /{{.GisServiceCode}}/virtual-servers/{{.IvmServiceCode}}/data-devices.json
func (t DataDeviceStorageConnect) URI() string {
	return "/{{.GisServiceCode}}/virtual-servers/{{.IvmServiceCode}}/data-devices.json"
}

// APIName DataDeviceStorageConnect
func (t DataDeviceStorageConnect) APIName() string {
	return "DataDeviceStorageConnect"
}

// Method POST
func (t DataDeviceStorageConnect) Method() string {
	return "POST"
}

// Document http://manual.iij.jp/p2/pubapi/59939590.html
func (t DataDeviceStorageConnect) Document() string {
	return "http://manual.iij.jp/p2/pubapi/59939590.html"
}

// JPName データデバイスストレージ接続
func (t DataDeviceStorageConnect) JPName() string {
	return "データデバイスストレージ接続"
}
func init() {
	APIlist = append(APIlist, DataDeviceStorageConnect{})
	TypeMap["DataDeviceStorageConnect"] = reflect.TypeOf(DataDeviceStorageConnect{})
}

// DataDeviceStorageConnectResponse データデバイスストレージ接続のレスポンス
type DataDeviceStorageConnectResponse struct {
	*CommonResponse
	PciSlot     string `json:",omitempty"` // PCIスロット(16進数)
	ServiceCode string `json:",omitempty"` // ストレージのサービスコード(iba########)
	OSType      string `json:",omitempty"` // OS種別(Windows/Linux)
	URI         string `json:",omitempty"` // ストレージの情報へアクセスするためのURI(URI)
	Category    string `json:",omitempty"` // ストレージ種別(BestEffort/Guarantee)
	Type        string `json:",omitempty"` // ストレージ品目
}

package protocol

import (
	"reflect"
)

// DataDeviceStorageListGet データデバイスストレージ情報一覧取得 (同期)
//  http://manual.iij.jp/p2/pubapi/59939598.html
type DataDeviceStorageListGet struct {
	GisServiceCode string `json:"-"` // P2契約のサービスコード(gis########)
	IvmServiceCode string `json:"-"` // 仮想サーバのサービスコード(ivm########, ivd########)
}

// URI /{{.GisServiceCode}}/virtual-servers/{{.IvmServiceCode}}/data-devices.json
func (t DataDeviceStorageListGet) URI() string {
	return "/{{.GisServiceCode}}/virtual-servers/{{.IvmServiceCode}}/data-devices.json"
}

// APIName DataDeviceStorageListGet
func (t DataDeviceStorageListGet) APIName() string {
	return "DataDeviceStorageListGet"
}

// Method GET
func (t DataDeviceStorageListGet) Method() string {
	return "GET"
}

// Document http://manual.iij.jp/p2/pubapi/59939598.html
func (t DataDeviceStorageListGet) Document() string {
	return "http://manual.iij.jp/p2/pubapi/59939598.html"
}

// JPName データデバイスストレージ情報一覧取得
func (t DataDeviceStorageListGet) JPName() string {
	return "データデバイスストレージ情報一覧取得"
}
func init() {
	APIlist = append(APIlist, DataDeviceStorageListGet{})
	TypeMap["DataDeviceStorageListGet"] = reflect.TypeOf(DataDeviceStorageListGet{})
}

// DataDeviceStorageListGetResponse データデバイスストレージ情報一覧取得のレスポンス
type DataDeviceStorageListGetResponse struct {
	*CommonResponse
	StorageList []struct {
		PciSlot     string `json:",omitempty"` // PCIスロット(16進数)
		ServiceCode string `json:",omitempty"` // ストレージのサービスコード(iba/ibb/ibg/ica/icb/icg########)
		OSType      string `json:",omitempty"` // OS種別(None/Windows/Linux)
		URI         string `json:",omitempty"` // ストレージの情報へアクセスするためのURI(URI)
		Category    string `json:",omitempty"` // ストレージ種別(BestEffort/Guarantee)
		Type        string `json:",omitempty"` // ストレージ品目
	}
}

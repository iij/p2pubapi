package protocol

import (
	"reflect"
)

// SystemStorageListGet システムストレージ情報一覧取得 (同期)
//  http://manual.iij.jp/p2/pubapi/59939923.html
type SystemStorageListGet struct {
	GisServiceCode string `json:"-"`                // P2契約のサービスコード(gis########)
	StartIndex     string `json:"-" p2pub:",query"` // StartIndexだけスキップした位置から情報を取得する(数値)
	Count          string `json:"-" p2pub:",query"` // 取得する情報の最大数(数値)
}

// URI /{{.GisServiceCode}}/system-storages.json
func (t SystemStorageListGet) URI() string {
	return "/{{.GisServiceCode}}/system-storages.json"
}

// APIName SystemStorageListGet
func (t SystemStorageListGet) APIName() string {
	return "SystemStorageListGet"
}

// Method GET
func (t SystemStorageListGet) Method() string {
	return "GET"
}

// Document http://manual.iij.jp/p2/pubapi/59939923.html
func (t SystemStorageListGet) Document() string {
	return "http://manual.iij.jp/p2/pubapi/59939923.html"
}

// JPName システムストレージ情報一覧取得
func (t SystemStorageListGet) JPName() string {
	return "システムストレージ情報一覧取得"
}
func init() {
	APIlist = append(APIlist, SystemStorageListGet{})
	TypeMap["SystemStorageListGet"] = reflect.TypeOf(SystemStorageListGet{})
}

// SystemStorageListGetResponse システムストレージ情報一覧取得のレスポンス
type SystemStorageListGetResponse struct {
	*CommonResponse
	SystemStorageList []struct {
		ResourceStatus        string `json:",omitempty"` // ストレージステータス
		AttachedVirtualServer struct {
			OSType      string `json:",omitempty"` // OS種別(Windows/Linux)
			URI         string `json:",omitempty"` // 仮想サーバの情報へアクセスするためのURI(URI)
			PciSlot     string `json:",omitempty"` // PCIスロット(16進数)
			ServiceCode string `json:",omitempty"` // 接続された仮想サーバのサービスコード(ivd########/ivm########)
			Type        string `json:",omitempty"` // 仮想サーバ品目
		} `json:",omitempty"`
		Label          string `json:",omitempty"` // ラベル(文字列)
		ServiceCode    string `json:",omitempty"` // システムストレージのサービスコード(iba########)
		OSType         string `json:",omitempty"` // OS種別(Linux/Windows)
		StorageGroup   string `json:",omitempty"` // ストレージグループ(Z/Y)
		StartDate      string `json:",omitempty"` // 利用開始日(YYYYMMDD)
		StorageSize    string `json:",omitempty"` // ストレージ容量(数値)
		ContractStatus string `json:",omitempty"` // 契約状態
		StopDate       string `json:",omitempty"` // 解約予定日(YYYYMMDD)
		Type           string `json:",omitempty"` // システムストレージ品目
	}
}

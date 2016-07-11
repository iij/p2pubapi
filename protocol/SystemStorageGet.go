package protocol

import (
	"reflect"
)

// SystemStorageGet システムストレージ情報取得 (同期)
//  http://manual.iij.jp/p2/pubapi/59939955.html
type SystemStorageGet struct {
	GisServiceCode string `json:"-"` // P2契約のサービスコード(gis########)
	IbaServiceCode string `json:"-"` // システムストレージのサービスコード(iba########)
}

// URI /{{.GisServiceCode}}/system-storages/{{.IbaServiceCode}}.json
func (t SystemStorageGet) URI() string {
	return "/{{.GisServiceCode}}/system-storages/{{.IbaServiceCode}}.json"
}

// APIName SystemStorageGet
func (t SystemStorageGet) APIName() string {
	return "SystemStorageGet"
}

// Method GET
func (t SystemStorageGet) Method() string {
	return "GET"
}

// Document http://manual.iij.jp/p2/pubapi/59939955.html
func (t SystemStorageGet) Document() string {
	return "http://manual.iij.jp/p2/pubapi/59939955.html"
}

// JPName システムストレージ情報取得
func (t SystemStorageGet) JPName() string {
	return "システムストレージ情報取得"
}
func init() {
	APIlist = append(APIlist, SystemStorageGet{})
	TypeMap["SystemStorageGet"] = reflect.TypeOf(SystemStorageGet{})
}

// SystemStorageGetResponse システムストレージ情報取得のレスポンス
type SystemStorageGetResponse struct {
	*CommonResponse
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

package protocol

import (
	"reflect"
)

// SystemStorageAdd システムストレージ追加申込 (非同期)
//  http://manual.iij.jp/p2/pubapi/59939812.html
type SystemStorageAdd struct {
	GisServiceCode string `json:"-"`          // P2契約のサービスコード(gis########)
	Encryption     string `json:",omitempty"` // 暗号化 ("Yes", "No")
	StorageGroup   string `json:",omitempty"` // ストレージグループ。省略時はどちらかのグループへ自動的に割り当てられます(Z/Y)
	Type           string // ストレージ品目
}

// URI /{{.GisServiceCode}}/system-storages.json
func (t SystemStorageAdd) URI() string {
	return "/{{.GisServiceCode}}/system-storages.json"
}

// APIName SystemStorageAdd
func (t SystemStorageAdd) APIName() string {
	return "SystemStorageAdd"
}

// Method POST
func (t SystemStorageAdd) Method() string {
	return "POST"
}

// Document http://manual.iij.jp/p2/pubapi/59939812.html
func (t SystemStorageAdd) Document() string {
	return "http://manual.iij.jp/p2/pubapi/59939812.html"
}

// JPName システムストレージ追加申込
func (t SystemStorageAdd) JPName() string {
	return "システムストレージ追加申込"
}
func init() {
	APIlist = append(APIlist, SystemStorageAdd{})
	TypeMap["SystemStorageAdd"] = reflect.TypeOf(SystemStorageAdd{})
}

// SystemStorageAddResponse システムストレージ追加申込のレスポンス
type SystemStorageAddResponse struct {
	*CommonResponse
	ResourceStatus string `json:",omitempty"` // ストレージステータス
	Label          string `json:",omitempty"` // ラベル(文字列)
	ServiceCode    string `json:",omitempty"` // システムストレージのサービスコード(iba########/ica########)
	OSType         string `json:",omitempty"` // OS種別(Linux/Windows)
	StorageGroup   string `json:",omitempty"` // ストレージグループ(Z/Y)
	StartDate      string `json:",omitempty"` // 利用開始日(YYYYMMDD)
	StorageSize    string `json:",omitempty"` // ストレージ容量(数値)
	ContractStatus string `json:",omitempty"` // 契約状態
	StopDate       string `json:",omitempty"` // 解約予定日(YYYYMMDD)
	Type           string `json:",omitempty"` // システムストレージ品目
	Mode           string `json:",omitempty"` // ストレージのモード ("Basic", "Backup")
	Encryption     string `json:",omitempty"` // 暗号化 ("Yes", "No")
}

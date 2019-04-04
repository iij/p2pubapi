package protocol

import (
	"reflect"
)

// StorageAdd 追加ストレージ追加申込 (同期)
//  http://manual.iij.jp/p2/pubapi/59940088.html
type StorageAdd struct {
	GisServiceCode string `json:"-"`          // P2契約のサービスコード(gis########)
	Encryption     string // 暗号化 ("Yes", "No")
	StorageGroup   string `json:",omitempty"` // ストレージグループ(Z/Y)
	Type           string // 追加ストレージ品目
}

// URI /{{.GisServiceCode}}/additional-storages.json
func (t StorageAdd) URI() string {
	return "/{{.GisServiceCode}}/additional-storages.json"
}

// APIName StorageAdd
func (t StorageAdd) APIName() string {
	return "StorageAdd"
}

// Method POST
func (t StorageAdd) Method() string {
	return "POST"
}

// Document http://manual.iij.jp/p2/pubapi/59940088.html
func (t StorageAdd) Document() string {
	return "http://manual.iij.jp/p2/pubapi/59940088.html"
}

// JPName 追加ストレージ追加申込
func (t StorageAdd) JPName() string {
	return "追加ストレージ追加申込"
}
func init() {
	APIlist = append(APIlist, StorageAdd{})
	TypeMap["StorageAdd"] = reflect.TypeOf(StorageAdd{})
}

// StorageAddResponse 追加ストレージ追加申込のレスポンス
type StorageAddResponse struct {
	*CommonResponse
	ResourceStatus string `json:",omitempty"` // ストレージステータス
	Label          string `json:",omitempty"` // ラベル(文字列)
	ServiceCode    string `json:",omitempty"` // 追加ストレージのサービスコード(ibg########/ibb########/icg########/icb########)
	OSType         string `json:",omitempty"` // OS種別(None)
	StorageGroup   string `json:",omitempty"` // ストレージグループ(Z/Y)
	StartDate      string `json:",omitempty"` // 利用開始日(YYYYMMDD)
	StorageSize    string `json:",omitempty"` // ストレージ容量(数値)
	Category       string `json:",omitempty"` // 追加ストレージ種別(BestEffort/Gurantee)
	ContractStatus string `json:",omitempty"` // 契約状態
	StopDate       string `json:",omitempty"` // 解約予定日(YYYYMMDD。未設定ならば空文字)
	Type           string `json:",omitempty"` // 追加ストレージ品目
	Mode           string `json:",omitempty"` // ストレージのモード ("Basic", "Backup")
	Encryption     string `json:",omitempty"` // 暗号化 ("Yes", "No")
}

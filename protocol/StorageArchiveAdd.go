package protocol

import (
	"reflect"
)

// StorageArchiveAdd ストレージアーカイブ追加申込 (同期)
//  http://manual.iij.jp/p2/pubapi/59940198.html
type StorageArchiveAdd struct {
	GisServiceCode string `json:"-"` // P2契約のサービスコード(gis########)
	ArchiveSize    string // ストレージアーカイブの容量（GB）(数値（10～1000）。10GB単位のみ指定可能)
}

// URI /{{.GisServiceCode}}/storage-archive.json
func (t StorageArchiveAdd) URI() string {
	return "/{{.GisServiceCode}}/storage-archive.json"
}

// APIName StorageArchiveAdd
func (t StorageArchiveAdd) APIName() string {
	return "StorageArchiveAdd"
}

// Method POST
func (t StorageArchiveAdd) Method() string {
	return "POST"
}

// Document http://manual.iij.jp/p2/pubapi/59940198.html
func (t StorageArchiveAdd) Document() string {
	return "http://manual.iij.jp/p2/pubapi/59940198.html"
}

// JPName ストレージアーカイブ追加申込
func (t StorageArchiveAdd) JPName() string {
	return "ストレージアーカイブ追加申込"
}
func init() {
	APIlist = append(APIlist, StorageArchiveAdd{})
	TypeMap["StorageArchiveAdd"] = reflect.TypeOf(StorageArchiveAdd{})
}

// StorageArchiveAddResponse ストレージアーカイブ追加申込のレスポンス
type StorageArchiveAddResponse struct {
	*CommonResponse
	Stat struct {
		Total string `json:",omitempty"` // ストレージアーカイブの容量（MB）(文字列)
		Free  string `json:",omitempty"` // ストレージアーカイブの空き容量（MB）(文字列)
		Used  string `json:",omitempty"` // ストレージアーカイブの使用量（MB）(文字列)
	} `json:",omitempty"`
	ServiceCode    string `json:",omitempty"` // ストレージアーカイブのサービスコード(iar########)
	ArchiveSize    string `json:",omitempty"` // ストレージアーカイブの容量（GB）(数値)
	StartDate      string `json:",omitempty"` // 利用開始日(YYYYMMDD)
	ContractStatus string `json:",omitempty"` // 契約状態
	StopDate       string `json:",omitempty"` // 解約予定日(YYYYMMDD。未設定ならば空文字列)
}

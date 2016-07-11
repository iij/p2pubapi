package protocol

import (
	"reflect"
)

// StorageArchiveGet ストレージアーカイブ情報取得 (同期)
//  http://manual.iij.jp/p2/pubapi/59940285.html
type StorageArchiveGet struct {
	GisServiceCode string `json:"-"` // P2契約のサービスコード(gis########)
	IarServiceCode string `json:"-"` // ストレージアーカイブのサービスコード(iar########)
}

// URI /{{.GisServiceCode}}/storage-archive/{{.IarServiceCode}}.json
func (t StorageArchiveGet) URI() string {
	return "/{{.GisServiceCode}}/storage-archive/{{.IarServiceCode}}.json"
}

// APIName StorageArchiveGet
func (t StorageArchiveGet) APIName() string {
	return "StorageArchiveGet"
}

// Method GET
func (t StorageArchiveGet) Method() string {
	return "GET"
}

// Document http://manual.iij.jp/p2/pubapi/59940285.html
func (t StorageArchiveGet) Document() string {
	return "http://manual.iij.jp/p2/pubapi/59940285.html"
}

// JPName ストレージアーカイブ情報取得
func (t StorageArchiveGet) JPName() string {
	return "ストレージアーカイブ情報取得"
}
func init() {
	APIlist = append(APIlist, StorageArchiveGet{})
	TypeMap["StorageArchiveGet"] = reflect.TypeOf(StorageArchiveGet{})
}

// StorageArchiveGetResponse ストレージアーカイブ情報取得のレスポンス
type StorageArchiveGetResponse struct {
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
	StopDate       string `json:",omitempty"` // 解約予定日(YYYYMMDD)
}

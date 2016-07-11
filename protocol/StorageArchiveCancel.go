package protocol

import (
	"reflect"
)

// StorageArchiveCancel ストレージアーカイブ解約申込 (同期)
//  http://manual.iij.jp/p2/pubapi/59940330.html
type StorageArchiveCancel struct {
	GisServiceCode string `json:"-"`                // P2契約のサービスコード(gis########)
	IarServiceCode string `json:"-"`                // ストレージアーカイブのサービスコード(iar########)
	StopDate       string `json:"-" p2pub:",query"` // 解約予定日。省略した場合は即時(YYYYMMDD)
}

// URI /{{.GisServiceCode}}/storage-archive/{{.IarServiceCode}}.json
func (t StorageArchiveCancel) URI() string {
	return "/{{.GisServiceCode}}/storage-archive/{{.IarServiceCode}}.json"
}

// APIName StorageArchiveCancel
func (t StorageArchiveCancel) APIName() string {
	return "StorageArchiveCancel"
}

// Method DELETE
func (t StorageArchiveCancel) Method() string {
	return "DELETE"
}

// Document http://manual.iij.jp/p2/pubapi/59940330.html
func (t StorageArchiveCancel) Document() string {
	return "http://manual.iij.jp/p2/pubapi/59940330.html"
}

// JPName ストレージアーカイブ解約申込
func (t StorageArchiveCancel) JPName() string {
	return "ストレージアーカイブ解約申込"
}
func init() {
	APIlist = append(APIlist, StorageArchiveCancel{})
	TypeMap["StorageArchiveCancel"] = reflect.TypeOf(StorageArchiveCancel{})
}

// StorageArchiveCancelResponse ストレージアーカイブ解約申込のレスポンス
type StorageArchiveCancelResponse struct {
	*CommonResponse
	Current struct {
		StopDate string `json:",omitempty"` // 設定した解約予定日(YYYYMMDD)
	} `json:",omitempty"`
	Previous struct {
		StopDate string `json:",omitempty"` // 設定前の解約予定日(YYYYMMDD)
	} `json:",omitempty"`
}

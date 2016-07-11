package protocol

import (
	"reflect"
)

// StorageArchiveSizeChange ストレージアーカイブサイズ変更申込 (同期)
//  http://manual.iij.jp/p2/pubapi/59940310.html
type StorageArchiveSizeChange struct {
	GisServiceCode string `json:"-"` // P2契約のサービスコード(gis########)
	IarServiceCode string `json:"-"` // ストレージアーカイブのサービスコード(iar########)
	ArchiveSize    string // ストレージアーカイブの容量（GB）(数値。10～1000GBの範囲で10GB単位で指定可能)
}

// URI /{{.GisServiceCode}}/storage-archive/{{.IarServiceCode}}.json
func (t StorageArchiveSizeChange) URI() string {
	return "/{{.GisServiceCode}}/storage-archive/{{.IarServiceCode}}.json"
}

// APIName StorageArchiveSizeChange
func (t StorageArchiveSizeChange) APIName() string {
	return "StorageArchiveSizeChange"
}

// Method PUT
func (t StorageArchiveSizeChange) Method() string {
	return "PUT"
}

// Document http://manual.iij.jp/p2/pubapi/59940310.html
func (t StorageArchiveSizeChange) Document() string {
	return "http://manual.iij.jp/p2/pubapi/59940310.html"
}

// JPName ストレージアーカイブサイズ変更申込
func (t StorageArchiveSizeChange) JPName() string {
	return "ストレージアーカイブサイズ変更申込"
}
func init() {
	APIlist = append(APIlist, StorageArchiveSizeChange{})
	TypeMap["StorageArchiveSizeChange"] = reflect.TypeOf(StorageArchiveSizeChange{})
}

// StorageArchiveSizeChangeResponse ストレージアーカイブサイズ変更申込のレスポンス
type StorageArchiveSizeChangeResponse struct {
	*CommonResponse
	Current struct {
		ArchiveSize string `json:",omitempty"` // 設定したストレージアーカイブの容量（GB）(数値)
	} `json:",omitempty"`
	Previous struct {
		ArchiveSize string `json:",omitempty"` // 設定前のストレージアーカイブの容量（GB）(数値)
	} `json:",omitempty"`
}

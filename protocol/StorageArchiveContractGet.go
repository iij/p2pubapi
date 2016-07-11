package protocol

import (
	"reflect"
)

// StorageArchiveContractGet ストレージアーカイブ契約状態取得 (同期)
//  http://manual.iij.jp/p2/pubapi/59940295.html
type StorageArchiveContractGet struct {
	GisServiceCode string `json:"-"` // P2契約のサービスコード(gis########)
	IarServiceCode string `json:"-"` // ストレージアーカイブのサービスコード(iar########)
}

// URI /{{.GisServiceCode}}/storage-archive/{{.IarServiceCode}}.json
func (t StorageArchiveContractGet) URI() string {
	return "/{{.GisServiceCode}}/storage-archive/{{.IarServiceCode}}.json"
}

// APIName StorageArchiveContractGet
func (t StorageArchiveContractGet) APIName() string {
	return "StorageArchiveContractGet"
}

// Method GET
func (t StorageArchiveContractGet) Method() string {
	return "GET"
}

// Document http://manual.iij.jp/p2/pubapi/59940295.html
func (t StorageArchiveContractGet) Document() string {
	return "http://manual.iij.jp/p2/pubapi/59940295.html"
}

// JPName ストレージアーカイブ契約状態取得
func (t StorageArchiveContractGet) JPName() string {
	return "ストレージアーカイブ契約状態取得"
}
func init() {
	APIlist = append(APIlist, StorageArchiveContractGet{})
	TypeMap["StorageArchiveContractGet"] = reflect.TypeOf(StorageArchiveContractGet{})
}

// StorageArchiveContractGetResponse ストレージアーカイブ契約状態取得のレスポンス
type StorageArchiveContractGetResponse struct {
	*CommonResponse
	ContractStatus string `json:",omitempty"` // 契約状態
}

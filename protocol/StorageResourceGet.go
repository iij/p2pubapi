package protocol

import (
	"reflect"
)

// StorageResourceGet 追加ストレージリソース状態取得 (同期)
//  http://manual.iij.jp/p2/pubapi/59940176.html
type StorageResourceGet struct {
	GisServiceCode string `json:"-"`                // P2契約のサービスコード(gis########)
	IbgServiceCode string `json:"-"`                // 追加ストレージのサービスコード(ibb########, ibg########)
	Item           string `json:"-" p2pub:",query"` // 取得するフィールド("ResourceStatus")
}

// URI /{{.GisServiceCode}}/additional-storages/{{.IbgServiceCode}}.json
func (t StorageResourceGet) URI() string {
	return "/{{.GisServiceCode}}/additional-storages/{{.IbgServiceCode}}.json"
}

// APIName StorageResourceGet
func (t StorageResourceGet) APIName() string {
	return "StorageResourceGet"
}

// Method GET
func (t StorageResourceGet) Method() string {
	return "GET"
}

// Document http://manual.iij.jp/p2/pubapi/59940176.html
func (t StorageResourceGet) Document() string {
	return "http://manual.iij.jp/p2/pubapi/59940176.html"
}

// JPName 追加ストレージリソース状態取得
func (t StorageResourceGet) JPName() string {
	return "追加ストレージリソース状態取得"
}
func init() {
	APIlist = append(APIlist, StorageResourceGet{})
	TypeMap["StorageResourceGet"] = reflect.TypeOf(StorageResourceGet{})
}

// StorageResourceGetResponse 追加ストレージリソース状態取得のレスポンス
type StorageResourceGetResponse struct {
	*CommonResponse
	ResourceStatus string `json:",omitempty"` // ストレージステータス
}

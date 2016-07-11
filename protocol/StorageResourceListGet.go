package protocol

import (
	"reflect"
)

// StorageResourceListGet 追加ストレージリソース状態一覧取得 (同期)
//  http://manual.iij.jp/p2/pubapi/59940135.html
type StorageResourceListGet struct {
	GisServiceCode string `json:"-"`                // P2契約のサービスコード(gis########)
	Item           string `json:"-" p2pub:",query"` // 取得するフィールド("ResourceStatus")
	StartIndex     string `json:"-" p2pub:",query"` // StartIndexだけスキップした位置から情報を取得する(数値)
	Count          string `json:"-" p2pub:",query"` // 取得する情報の最大数(数値)
}

// URI /{{.GisServiceCode}}/additional-storages.json
func (t StorageResourceListGet) URI() string {
	return "/{{.GisServiceCode}}/additional-storages.json"
}

// APIName StorageResourceListGet
func (t StorageResourceListGet) APIName() string {
	return "StorageResourceListGet"
}

// Method GET
func (t StorageResourceListGet) Method() string {
	return "GET"
}

// Document http://manual.iij.jp/p2/pubapi/59940135.html
func (t StorageResourceListGet) Document() string {
	return "http://manual.iij.jp/p2/pubapi/59940135.html"
}

// JPName 追加ストレージリソース状態一覧取得
func (t StorageResourceListGet) JPName() string {
	return "追加ストレージリソース状態一覧取得"
}
func init() {
	APIlist = append(APIlist, StorageResourceListGet{})
	TypeMap["StorageResourceListGet"] = reflect.TypeOf(StorageResourceListGet{})
}

// StorageResourceListGetResponse 追加ストレージリソース状態一覧取得のレスポンス
type StorageResourceListGetResponse struct {
	*CommonResponse
	AdditionalStorageList []struct {
		AdditionalStorageList []struct {
			ResourceStatus string `json:",omitempty"` // ストレージステータス
			ServiceCode    string `json:",omitempty"` // 追加ストレージのサービスコード(ibg########/ibb########)
		} `json:",omitempty"`
	}
}

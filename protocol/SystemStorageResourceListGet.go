package protocol

import (
	"reflect"
)

// SystemStorageResourceListGet システムストレージリソース状態一覧取得 (同期)
//  http://manual.iij.jp/p2/pubapi/59939946.html
type SystemStorageResourceListGet struct {
	GisServiceCode string `json:"-"`                // P2契約のサービスコード(gis########)
	Item           string `json:"-" p2pub:",query"` // 取得するフィールド("ResourceStatus")
	StartIndex     string `json:"-" p2pub:",query"` // StartIndexだけスキップした位置から情報を取得する(数値)
	Count          string `json:"-" p2pub:",query"` // 取得する情報の最大数(数値)
}

// URI /{{.GisServiceCode}}/system-storages.json
func (t SystemStorageResourceListGet) URI() string {
	return "/{{.GisServiceCode}}/system-storages.json"
}

// APIName SystemStorageResourceListGet
func (t SystemStorageResourceListGet) APIName() string {
	return "SystemStorageResourceListGet"
}

// Method GET
func (t SystemStorageResourceListGet) Method() string {
	return "GET"
}

// Document http://manual.iij.jp/p2/pubapi/59939946.html
func (t SystemStorageResourceListGet) Document() string {
	return "http://manual.iij.jp/p2/pubapi/59939946.html"
}

// JPName システムストレージリソース状態一覧取得
func (t SystemStorageResourceListGet) JPName() string {
	return "システムストレージリソース状態一覧取得"
}
func init() {
	APIlist = append(APIlist, SystemStorageResourceListGet{})
	TypeMap["SystemStorageResourceListGet"] = reflect.TypeOf(SystemStorageResourceListGet{})
}

// SystemStorageResourceListGetResponse システムストレージリソース状態一覧取得のレスポンス
type SystemStorageResourceListGetResponse struct {
	*CommonResponse
	SystemStorageList []struct {
		SystemStorageList []struct {
			ResourceStatus string `json:",omitempty"` // ストレージステータス
			ServiceCode    string `json:",omitempty"` // システムストレージのサービスコード(iba########)
		} `json:",omitempty"`
	}
}

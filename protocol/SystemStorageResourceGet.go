package protocol

import (
	"reflect"
)

// SystemStorageResourceGet システムストレージリソース状態取得 (同期)
//  http://manual.iij.jp/p2/pubapi/59939978.html
type SystemStorageResourceGet struct {
	GisServiceCode string `json:"-"`                // P2契約のサービスコード(gis########)
	IbaServiceCode string `json:"-"`                // システムストレージのサービスコード(iba########)
	Item           string `json:"-" p2pub:",query"` // 取得するフィールド("ResourceStatus")
}

// URI /{{.GisServiceCode}}/system-storages/{{.IbaServiceCode}}.json
func (t SystemStorageResourceGet) URI() string {
	return "/{{.GisServiceCode}}/system-storages/{{.IbaServiceCode}}.json"
}

// APIName SystemStorageResourceGet
func (t SystemStorageResourceGet) APIName() string {
	return "SystemStorageResourceGet"
}

// Method GET
func (t SystemStorageResourceGet) Method() string {
	return "GET"
}

// Document http://manual.iij.jp/p2/pubapi/59939978.html
func (t SystemStorageResourceGet) Document() string {
	return "http://manual.iij.jp/p2/pubapi/59939978.html"
}

// JPName システムストレージリソース状態取得
func (t SystemStorageResourceGet) JPName() string {
	return "システムストレージリソース状態取得"
}
func init() {
	APIlist = append(APIlist, SystemStorageResourceGet{})
	TypeMap["SystemStorageResourceGet"] = reflect.TypeOf(SystemStorageResourceGet{})
}

// SystemStorageResourceGetResponse システムストレージリソース状態取得のレスポンス
type SystemStorageResourceGetResponse struct {
	*CommonResponse
	ResourceStatus string `json:",omitempty"` // ストレージステータス
}

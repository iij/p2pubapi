package protocol

import (
	"reflect"
)

// VMResourceGet 仮想サーバリソース状態取得 (同期)
//  http://manual.iij.jp/p2/pubapi/59939450.html
type VMResourceGet struct {
	GisServiceCode string `json:"-"`                // P2契約のサービスコード(gis########)
	IvmServiceCode string `json:"-"`                // 仮想サーバのサービスコード(ivm########, ivd########)
	Item           string `json:"-" p2pub:",query"` // 取得するフィールド("ResourceStatus")
}

// URI /{{.GisServiceCode}}/virtual-servers/{{.IvmServiceCode}}.json
func (t VMResourceGet) URI() string {
	return "/{{.GisServiceCode}}/virtual-servers/{{.IvmServiceCode}}.json"
}

// APIName VMResourceGet
func (t VMResourceGet) APIName() string {
	return "VMResourceGet"
}

// Method GET
func (t VMResourceGet) Method() string {
	return "GET"
}

// Document http://manual.iij.jp/p2/pubapi/59939450.html
func (t VMResourceGet) Document() string {
	return "http://manual.iij.jp/p2/pubapi/59939450.html"
}

// JPName 仮想サーバリソース状態取得
func (t VMResourceGet) JPName() string {
	return "仮想サーバリソース状態取得"
}
func init() {
	APIlist = append(APIlist, VMResourceGet{})
	TypeMap["VMResourceGet"] = reflect.TypeOf(VMResourceGet{})
}

// VMResourceGetResponse 仮想サーバリソース状態取得のレスポンス
type VMResourceGetResponse struct {
	*CommonResponse
	ResourceStatus string `json:",omitempty"` // 仮想サーバステータス
}

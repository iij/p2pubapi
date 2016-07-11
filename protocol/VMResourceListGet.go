package protocol

import (
	"reflect"
)

// VMResourceListGet 仮想サーバリソース状態一覧取得 (同期)
//  http://manual.iij.jp/p2/pubapi/59939418.html
type VMResourceListGet struct {
	GisServiceCode string `json:"-"`                // GisServiceCode
	Item           string `json:"-" p2pub:",query"` // 取得するフィールド("ResourceStatus")
	StartIndex     string `json:"-" p2pub:",query"` // StartIndexだけスキップした位置から情報を取得する(数値)
	Count          string `json:"-" p2pub:",query"` // 取得する情報の最大数(数値)
}

// URI /{{.GisServiceCode}}/virtual-servers.json
func (t VMResourceListGet) URI() string {
	return "/{{.GisServiceCode}}/virtual-servers.json"
}

// APIName VMResourceListGet
func (t VMResourceListGet) APIName() string {
	return "VMResourceListGet"
}

// Method GET
func (t VMResourceListGet) Method() string {
	return "GET"
}

// Document http://manual.iij.jp/p2/pubapi/59939418.html
func (t VMResourceListGet) Document() string {
	return "http://manual.iij.jp/p2/pubapi/59939418.html"
}

// JPName 仮想サーバリソース状態一覧取得
func (t VMResourceListGet) JPName() string {
	return "仮想サーバリソース状態一覧取得"
}
func init() {
	APIlist = append(APIlist, VMResourceListGet{})
	TypeMap["VMResourceListGet"] = reflect.TypeOf(VMResourceListGet{})
}

// VMResourceListGetResponse 仮想サーバリソース状態一覧取得のレスポンス
type VMResourceListGetResponse struct {
	*CommonResponse
	VirtualServerList []struct {
		VirtualServerList []struct {
			ResourceStatus string `json:",omitempty"` // 仮想サーバステータス
			ServiceCode    string `json:",omitempty"` // サービスコード(ivd########/ivm########)
		} `json:",omitempty"`
	}
}

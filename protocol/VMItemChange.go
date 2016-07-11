package protocol

import (
	"reflect"
)

// VMItemChange 仮想サーバ品目変更申込 (同期)
//  http://manual.iij.jp/p2/pubapi/59939458.html
type VMItemChange struct {
	GisServiceCode string `json:"-"` // GisServiceCode
	IvmServiceCode string `json:"-"` // IvmServiceCode
	Type           string // 仮想サーバ品目
}

// URI /{{.GisServiceCode}}/virtual-servers/{{.IvmServiceCode}}.json
func (t VMItemChange) URI() string {
	return "/{{.GisServiceCode}}/virtual-servers/{{.IvmServiceCode}}.json"
}

// APIName VMItemChange
func (t VMItemChange) APIName() string {
	return "VMItemChange"
}

// Method PUT
func (t VMItemChange) Method() string {
	return "PUT"
}

// Document http://manual.iij.jp/p2/pubapi/59939458.html
func (t VMItemChange) Document() string {
	return "http://manual.iij.jp/p2/pubapi/59939458.html"
}

// JPName 仮想サーバ品目変更申込
func (t VMItemChange) JPName() string {
	return "仮想サーバ品目変更申込"
}
func init() {
	APIlist = append(APIlist, VMItemChange{})
	TypeMap["VMItemChange"] = reflect.TypeOf(VMItemChange{})
}

// VMItemChangeResponse 仮想サーバ品目変更申込のレスポンス
type VMItemChangeResponse struct {
	*CommonResponse
	Current struct {
		Type string `json:",omitempty"` // 設定したラベル(文字列)
	} `json:",omitempty"`
	Previous struct {
		Type string `json:",omitempty"` // 設定前のラベル(文字列)
	} `json:",omitempty"`
}

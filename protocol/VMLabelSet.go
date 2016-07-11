package protocol

import (
	"reflect"
)

// VMLabelSet 仮想サーバラベル設定 (同期)
//  http://manual.iij.jp/p2/pubapi/59939498.html
type VMLabelSet struct {
	GisServiceCode string `json:"-"` // P2契約のサービスコード(gis########)
	IvmServiceCode string `json:"-"` // 仮想サーバのサービスコード(ivm########, ivd########)
	Name           string // ラベル(文字列)
}

// URI /{{.GisServiceCode}}/virtual-servers/{{.IvmServiceCode}}/label.json
func (t VMLabelSet) URI() string {
	return "/{{.GisServiceCode}}/virtual-servers/{{.IvmServiceCode}}/label.json"
}

// APIName VMLabelSet
func (t VMLabelSet) APIName() string {
	return "VMLabelSet"
}

// Method PUT
func (t VMLabelSet) Method() string {
	return "PUT"
}

// Document http://manual.iij.jp/p2/pubapi/59939498.html
func (t VMLabelSet) Document() string {
	return "http://manual.iij.jp/p2/pubapi/59939498.html"
}

// JPName 仮想サーバラベル設定
func (t VMLabelSet) JPName() string {
	return "仮想サーバラベル設定"
}
func init() {
	APIlist = append(APIlist, VMLabelSet{})
	TypeMap["VMLabelSet"] = reflect.TypeOf(VMLabelSet{})
}

// VMLabelSetResponse 仮想サーバラベル設定のレスポンス
type VMLabelSetResponse struct {
	*CommonResponse
	Current struct {
		Label string `json:",omitempty"` // 設定したラベル(文字列)
	} `json:",omitempty"`
	Previous struct {
		Label string `json:",omitempty"` // 設定前のラベル(文字列)
	} `json:",omitempty"`
}

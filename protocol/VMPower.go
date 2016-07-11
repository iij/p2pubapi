package protocol

import (
	"reflect"
)

// VMPower 仮想サーバ停止 (非同期)
//  http://manual.iij.jp/p2/pubapi/59939532.html
//  http://manual.iij.jp/p2/pubapi/59939508.html
//  http://manual.iij.jp/p2/pubapi/59939523.html
type VMPower struct {
	GisServiceCode string `json:"-"` // GisServiceCode
	IvmServiceCode string `json:"-"` // IvmServiceCode
	Power          string // 電源状態を設定(On/Off/Reset)
	Force          string `json:",omitempty"` // 強制停止するならばYes, 通常のシャットダウンするならばNo。省略時には通常のシャットダウンを行う(No/Yes)
}

// URI /{{.GisServiceCode}}/virtual-servers/{{.IvmServiceCode}}/action.json
func (t VMPower) URI() string {
	return "/{{.GisServiceCode}}/virtual-servers/{{.IvmServiceCode}}/action.json"
}

// APIName VMPower
func (t VMPower) APIName() string {
	return "VMPower"
}

// Method PUT
func (t VMPower) Method() string {
	return "PUT"
}

// Document http://manual.iij.jp/p2/pubapi/59939532.html
func (t VMPower) Document() string {
	return "http://manual.iij.jp/p2/pubapi/59939532.html"
}

// JPName 仮想サーバ停止
func (t VMPower) JPName() string {
	return "仮想サーバ停止"
}
func init() {
	APIlist = append(APIlist, VMPower{})
	TypeMap["VMPower"] = reflect.TypeOf(VMPower{})
}

// VMPowerResponse 仮想サーバ停止のレスポンス
type VMPowerResponse struct {
	*CommonResponse
	Current struct {
		ResourceStatus string `json:",omitempty"` // 設定した仮想サーバステータス
	} `json:",omitempty"`
	Previous struct {
		ResourceStatus string `json:",omitempty"` // 設定前の仮想サーバステータス
	} `json:",omitempty"`
}

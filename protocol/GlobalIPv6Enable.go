package protocol

import (
	"reflect"
)

// GlobalIPv6Enable グローバルIPv6有効 (同期)
//  http://manual.iij.jp/p2/pubapi/59939671.html
type GlobalIPv6Enable struct {
	GisServiceCode string `json:"-"` // P2契約のサービスコード(gis########)
	IvmServiceCode string `json:"-"` // 仮想サーバのサービスコード(ivm########)
}

// URI /{{.GisServiceCode}}/virtual-servers/{{.IvmServiceCode}}/global-ipaddresses/ipv6
func (t GlobalIPv6Enable) URI() string {
	return "/{{.GisServiceCode}}/virtual-servers/{{.IvmServiceCode}}/global-ipaddresses/ipv6.json"
}

// APIName GlobalIPv6Enable
func (t GlobalIPv6Enable) APIName() string {
	return "GlobalIPv6Enable"
}

// Method PUT
func (t GlobalIPv6Enable) Method() string {
	return "PUT"
}

// Document http://manual.iij.jp/p2/pubapi/59939671.html
func (t GlobalIPv6Enable) Document() string {
	return "http://manual.iij.jp/p2/pubapi/59939671.html"
}

// JPName グローバルIPv6有効
func (t GlobalIPv6Enable) JPName() string {
	return "グローバルIPv6有効"
}
func init() {
	APIlist = append(APIlist, GlobalIPv6Enable{})
	TypeMap["GlobalIPv6Enable"] = reflect.TypeOf(GlobalIPv6Enable{})
}

// GlobalIPv6EnableResponse グローバルIPv6有効のレスポンス
type GlobalIPv6EnableResponse struct {
	*CommonResponse
	Current struct {
		IPv6Enabled string `json:",omitempty"` // 設定したIPv6状態(Yes)
	} `json:",omitempty"`
	Previous struct {
		IPv6Enabled string `json:",omitempty"` // 設定前のIPv6状態(No)
	} `json:",omitempty"`
}

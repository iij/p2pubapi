package protocol

import (
	"reflect"
)

// GlobalIPv6Disable グローバルIPv6無効 (同期)
//  http://manual.iij.jp/p2/pubapi/59939676.html
type GlobalIPv6Disable struct {
	GisServiceCode string `json:"-"` // P2契約のサービスコード(gis########)
	IvmServiceCode string `json:"-"` // 仮想サーバのサービスコード(ivm########)
}

// URI /{{.GisServiceCode}}/virtual-servers/{{.IvmServiceCode}}/global-ipaddresses/ipv6
func (t GlobalIPv6Disable) URI() string {
	return "/{{.GisServiceCode}}/virtual-servers/{{.IvmServiceCode}}/global-ipaddresses/ipv6.json"
}

// APIName GlobalIPv6Disable
func (t GlobalIPv6Disable) APIName() string {
	return "GlobalIPv6Disable"
}

// Method DELETE
func (t GlobalIPv6Disable) Method() string {
	return "DELETE"
}

// Document http://manual.iij.jp/p2/pubapi/59939676.html
func (t GlobalIPv6Disable) Document() string {
	return "http://manual.iij.jp/p2/pubapi/59939676.html"
}

// JPName グローバルIPv6無効
func (t GlobalIPv6Disable) JPName() string {
	return "グローバルIPv6無効"
}
func init() {
	APIlist = append(APIlist, GlobalIPv6Disable{})
	TypeMap["GlobalIPv6Disable"] = reflect.TypeOf(GlobalIPv6Disable{})
}

// GlobalIPv6DisableResponse グローバルIPv6無効のレスポンス
type GlobalIPv6DisableResponse struct {
	*CommonResponse
	Current struct {
		IPv6Enabled string `json:",omitempty"` // 設定したIPv6状態(No)
	} `json:",omitempty"`
	Previous struct {
		IPv6Enabled string `json:",omitempty"` // 設定前のIPv6状態(Yes)
	} `json:",omitempty"`
}

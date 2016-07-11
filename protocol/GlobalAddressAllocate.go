package protocol

import (
	"reflect"
)

// GlobalAddressAllocate グローバルIPアドレス割当 (同期)
//  http://manual.iij.jp/p2/pubapi/59939625.html
type GlobalAddressAllocate struct {
	GisServiceCode string `json:"-"` // P2契約のサービスコード(gis########)
	IvmServiceCode string `json:"-"` // 仮想サーバのサービスコード(ivm########)
}

// URI /{{.GisServiceCode}}/virtual-servers/{{.IvmServiceCode}}/global-ipaddresses.json
func (t GlobalAddressAllocate) URI() string {
	return "/{{.GisServiceCode}}/virtual-servers/{{.IvmServiceCode}}/global-ipaddresses.json"
}

// APIName GlobalAddressAllocate
func (t GlobalAddressAllocate) APIName() string {
	return "GlobalAddressAllocate"
}

// Method POST
func (t GlobalAddressAllocate) Method() string {
	return "POST"
}

// Document http://manual.iij.jp/p2/pubapi/59939625.html
func (t GlobalAddressAllocate) Document() string {
	return "http://manual.iij.jp/p2/pubapi/59939625.html"
}

// JPName グローバルIPアドレス割当
func (t GlobalAddressAllocate) JPName() string {
	return "グローバルIPアドレス割当"
}
func init() {
	APIlist = append(APIlist, GlobalAddressAllocate{})
	TypeMap["GlobalAddressAllocate"] = reflect.TypeOf(GlobalAddressAllocate{})
}

// GlobalAddressAllocateResponse グローバルIPアドレス割当のレスポンス
type GlobalAddressAllocateResponse struct {
	*CommonResponse
	IPv4 struct {
		DomainName string `json:",omitempty"` // 逆引きドメイン名(文字列)
		IpAddress  string `json:",omitempty"` // 割り当てられたIPv4アドレス(IPv4アドレス)
		Type       string `json:",omitempty"` // アドレス管理(Managed)
	} `json:",omitempty"`
	IPv6 struct {
		IpAddress string `json:",omitempty"` // 割り当てられたIPv6アドレス(IPv6アドレス)
		Type      string `json:",omitempty"` // アドレス管理(Managed)
	} `json:",omitempty"`
}

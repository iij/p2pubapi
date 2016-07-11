package protocol

import (
	"reflect"
)

// GlobalAddressGetForVM グローバルIPアドレス情報取得（仮想サーバ） (同期)
//  http://manual.iij.jp/p2/pubapi/59939650.html
type GlobalAddressGetForVM struct {
	GisServiceCode string `json:"-"` // P2契約のサービスコード(gis########)
	IvmServiceCode string `json:"-"` // 仮想サーバのサービスコード(ivm########)
	IPv4Address    string `json:"-"` // IPv4アドレス
	IPv            string `json:"-"` // IPv
}

// URI /{{.GisServiceCode}}/virtual-servers/{{.IvmServiceCode}}/global-ipaddresses/{{.IPv}}4Address.json
func (t GlobalAddressGetForVM) URI() string {
	return "/{{.GisServiceCode}}/virtual-servers/{{.IvmServiceCode}}/global-ipaddresses/{{.IPv}}4Address.json"
}

// APIName GlobalAddressGetForVM
func (t GlobalAddressGetForVM) APIName() string {
	return "GlobalAddressGetForVM"
}

// Method GET
func (t GlobalAddressGetForVM) Method() string {
	return "GET"
}

// Document http://manual.iij.jp/p2/pubapi/59939650.html
func (t GlobalAddressGetForVM) Document() string {
	return "http://manual.iij.jp/p2/pubapi/59939650.html"
}

// JPName グローバルIPアドレス情報取得（仮想サーバ）
func (t GlobalAddressGetForVM) JPName() string {
	return "グローバルIPアドレス情報取得（仮想サーバ）"
}
func init() {
	APIlist = append(APIlist, GlobalAddressGetForVM{})
	TypeMap["GlobalAddressGetForVM"] = reflect.TypeOf(GlobalAddressGetForVM{})
}

// GlobalAddressGetForVMResponse グローバルIPアドレス情報取得（仮想サーバ）のレスポンス
type GlobalAddressGetForVMResponse struct {
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

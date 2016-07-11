package protocol

import (
	"reflect"
)

// GlobalAddressRelease グローバルIPアドレス解放 (同期)
//  http://manual.iij.jp/p2/pubapi/59939664.html
type GlobalAddressRelease struct {
	GisServiceCode string `json:"-"` // P2契約のサービスコード(gis########)
	IvmServiceCode string `json:"-"` // 仮想サーバのサービスコード(ivm########)
	IPv4Address    string `json:"-"` // IPv4アドレス
	IPv            string `json:"-"` // IPv
}

// URI /{{.GisServiceCode}}/virtual-servers/{{.IvmServiceCode}}/global-ipaddresses/{{.IPv}}4Address.json
func (t GlobalAddressRelease) URI() string {
	return "/{{.GisServiceCode}}/virtual-servers/{{.IvmServiceCode}}/global-ipaddresses/{{.IPv}}4Address.json"
}

// APIName GlobalAddressRelease
func (t GlobalAddressRelease) APIName() string {
	return "GlobalAddressRelease"
}

// Method DELETE
func (t GlobalAddressRelease) Method() string {
	return "DELETE"
}

// Document http://manual.iij.jp/p2/pubapi/59939664.html
func (t GlobalAddressRelease) Document() string {
	return "http://manual.iij.jp/p2/pubapi/59939664.html"
}

// JPName グローバルIPアドレス解放
func (t GlobalAddressRelease) JPName() string {
	return "グローバルIPアドレス解放"
}
func init() {
	APIlist = append(APIlist, GlobalAddressRelease{})
	TypeMap["GlobalAddressRelease"] = reflect.TypeOf(GlobalAddressRelease{})
}

// GlobalAddressReleaseResponse グローバルIPアドレス解放のレスポンス
type GlobalAddressReleaseResponse struct {
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

package protocol

import (
	"reflect"
)

// GlobalAddressListGet グローバルIPアドレス情報一覧取得 (同期)
//  http://manual.iij.jp/p2/pubapi/59939639.html
type GlobalAddressListGet struct {
	GisServiceCode string `json:"-"` // P2契約のサービスコード(gis########)
	IvmServiceCode string `json:"-"` // 仮想サーバのサービスコード(ivm########)
}

// URI /{{.GisServiceCode}}/virtual-servers/{{.IvmServiceCode}}/global-ipaddresses.json
func (t GlobalAddressListGet) URI() string {
	return "/{{.GisServiceCode}}/virtual-servers/{{.IvmServiceCode}}/global-ipaddresses.json"
}

// APIName GlobalAddressListGet
func (t GlobalAddressListGet) APIName() string {
	return "GlobalAddressListGet"
}

// Method GET
func (t GlobalAddressListGet) Method() string {
	return "GET"
}

// Document http://manual.iij.jp/p2/pubapi/59939639.html
func (t GlobalAddressListGet) Document() string {
	return "http://manual.iij.jp/p2/pubapi/59939639.html"
}

// JPName グローバルIPアドレス情報一覧取得
func (t GlobalAddressListGet) JPName() string {
	return "グローバルIPアドレス情報一覧取得"
}
func init() {
	APIlist = append(APIlist, GlobalAddressListGet{})
	TypeMap["GlobalAddressListGet"] = reflect.TypeOf(GlobalAddressListGet{})
}

// GlobalAddressListGetResponse グローバルIPアドレス情報一覧取得のレスポンス
type GlobalAddressListGetResponse struct {
	*CommonResponse
	IpAddressList []struct {
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
}

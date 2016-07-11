package protocol

import (
	"reflect"
)

// GlobalAddressVAddIPAddressNumChange グローバルIPアドレス/V追加IPアドレス数変更申込 (同期)
//  http://manual.iij.jp/p2/pubapi/59940441.html
type GlobalAddressVAddIPAddressNumChange struct {
	GisServiceCode string `json:"-"` // P2契約のサービスコード(gis########)
	IgaServiceCode string `json:"-"` // グローバルIPアドレス/Vのサービスコード(iga########)
	AddressNum     string `json:"-"` // グローバルIPアドレス/Vで追加するアドレス数(数値)
}

// URI /{{.GisServiceCode}}/global-ipaddress/{{.IgaServiceCode}}.json
func (t GlobalAddressVAddIPAddressNumChange) URI() string {
	return "/{{.GisServiceCode}}/global-ipaddress/{{.IgaServiceCode}}.json"
}

// APIName GlobalAddressVAddIPAddressNumChange
func (t GlobalAddressVAddIPAddressNumChange) APIName() string {
	return "GlobalAddressVAddIPAddressNumChange"
}

// Method POST
func (t GlobalAddressVAddIPAddressNumChange) Method() string {
	return "POST"
}

// Document http://manual.iij.jp/p2/pubapi/59940441.html
func (t GlobalAddressVAddIPAddressNumChange) Document() string {
	return "http://manual.iij.jp/p2/pubapi/59940441.html"
}

// JPName グローバルIPアドレス/V追加IPアドレス数変更申込
func (t GlobalAddressVAddIPAddressNumChange) JPName() string {
	return "グローバルIPアドレス/V追加IPアドレス数変更申込"
}
func init() {
	APIlist = append(APIlist, GlobalAddressVAddIPAddressNumChange{})
	TypeMap["GlobalAddressVAddIPAddressNumChange"] = reflect.TypeOf(GlobalAddressVAddIPAddressNumChange{})
}

// GlobalAddressVAddIPAddressNumChangeResponse グローバルIPアドレス/V追加IPアドレス数変更申込のレスポンス
type GlobalAddressVAddIPAddressNumChangeResponse struct {
	*CommonResponse
	Current struct {
		AddressNum string `json:",omitempty"` // 設定した追加で利用できるグローバルIPアドレスの数(数値)
	} `json:",omitempty"`
	Previous struct {
		AddressNum string `json:",omitempty"` // 設定前の追加で利用できるグローバルIPアドレスの数(数値)
	} `json:",omitempty"`
}

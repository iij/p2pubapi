package protocol

import (
	"reflect"
)

// GlobalAddressVCancel グローバルIPアドレス/V解約申込 (同期)
//  http://manual.iij.jp/p2/pubapi/59940453.html
type GlobalAddressVCancel struct {
	GisServiceCode string `json:"-"` // P2契約のサービスコード(gis########)
	IgaServiceCode string `json:"-"` // グローバルIPアドレス/Vのサービスコード(iga########)
}

// URI /{{.GisServiceCode}}/global-ipaddress/{{.IgaServiceCode}}.json
func (t GlobalAddressVCancel) URI() string {
	return "/{{.GisServiceCode}}/global-ipaddress/{{.IgaServiceCode}}.json"
}

// APIName GlobalAddressVCancel
func (t GlobalAddressVCancel) APIName() string {
	return "GlobalAddressVCancel"
}

// Method DELETE
func (t GlobalAddressVCancel) Method() string {
	return "DELETE"
}

// Document http://manual.iij.jp/p2/pubapi/59940453.html
func (t GlobalAddressVCancel) Document() string {
	return "http://manual.iij.jp/p2/pubapi/59940453.html"
}

// JPName グローバルIPアドレス/V解約申込
func (t GlobalAddressVCancel) JPName() string {
	return "グローバルIPアドレス/V解約申込"
}
func init() {
	APIlist = append(APIlist, GlobalAddressVCancel{})
	TypeMap["GlobalAddressVCancel"] = reflect.TypeOf(GlobalAddressVCancel{})
}

// GlobalAddressVCancelResponse グローバルIPアドレス/V解約申込のレスポンス
type GlobalAddressVCancelResponse struct {
	*CommonResponse
	Current struct {
		StopDate string `json:",omitempty"` // 設定した解約予定日(YYYYMMDD)
	} `json:",omitempty"`
	Previous struct {
		StopDate string `json:",omitempty"` // 設定前の解約予定日(YYYYMMDD)
	} `json:",omitempty"`
}

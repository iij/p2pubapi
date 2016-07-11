package protocol

import (
	"reflect"
)

// GlobalAddressVAdd グローバルIPアドレス/V追加申込 (同期)
//  http://manual.iij.jp/p2/pubapi/59940392.html
type GlobalAddressVAdd struct {
	GisServiceCode string `json:"-"` // P2契約のサービスコード(gis########)
	AddressNum     string // グローバルIPアドレス/Vで追加するアドレス数(数値)
}

// URI /{{.GisServiceCode}}/global-ipaddress.json
func (t GlobalAddressVAdd) URI() string {
	return "/{{.GisServiceCode}}/global-ipaddress.json"
}

// APIName GlobalAddressVAdd
func (t GlobalAddressVAdd) APIName() string {
	return "GlobalAddressVAdd"
}

// Method POST
func (t GlobalAddressVAdd) Method() string {
	return "POST"
}

// Document http://manual.iij.jp/p2/pubapi/59940392.html
func (t GlobalAddressVAdd) Document() string {
	return "http://manual.iij.jp/p2/pubapi/59940392.html"
}

// JPName グローバルIPアドレス/V追加申込
func (t GlobalAddressVAdd) JPName() string {
	return "グローバルIPアドレス/V追加申込"
}
func init() {
	APIlist = append(APIlist, GlobalAddressVAdd{})
	TypeMap["GlobalAddressVAdd"] = reflect.TypeOf(GlobalAddressVAdd{})
}

// GlobalAddressVAddResponse グローバルIPアドレス/V追加申込のレスポンス
type GlobalAddressVAddResponse struct {
	*CommonResponse
	StartDate      string `json:",omitempty"` // 利用開始日(YYYYMMDD)
	AddressNum     string `json:",omitempty"` // 追加で利用できるグローバルIPアドレスの数(数値)
	ContractStatus string `json:",omitempty"` // 契約状態
	ServiceCode    string `json:",omitempty"` // グローバルIPアドレス/Vのサービスコード(iga########)
	StopDate       string `json:",omitempty"` // 解約予定日(YYYYMMDD)
}

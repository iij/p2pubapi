package protocol

import (
	"reflect"
)

// GlobalAddressVGet グローバルIPアドレス/V情報取得 (同期)
//  http://manual.iij.jp/p2/pubapi/59940431.html
type GlobalAddressVGet struct {
	GisServiceCode string `json:"-"` // P2契約のサービスコード(gis########)
	IgaServiceCode string `json:"-"` // グローバルIPアドレス/Vのサービスコード(iga########)
}

// URI /{{.GisServiceCode}}/global-ipaddress/{{.IgaServiceCode}}.json
func (t GlobalAddressVGet) URI() string {
	return "/{{.GisServiceCode}}/global-ipaddress/{{.IgaServiceCode}}.json"
}

// APIName GlobalAddressVGet
func (t GlobalAddressVGet) APIName() string {
	return "GlobalAddressVGet"
}

// Method GET
func (t GlobalAddressVGet) Method() string {
	return "GET"
}

// Document http://manual.iij.jp/p2/pubapi/59940431.html
func (t GlobalAddressVGet) Document() string {
	return "http://manual.iij.jp/p2/pubapi/59940431.html"
}

// JPName グローバルIPアドレス/V情報取得
func (t GlobalAddressVGet) JPName() string {
	return "グローバルIPアドレス/V情報取得"
}
func init() {
	APIlist = append(APIlist, GlobalAddressVGet{})
	TypeMap["GlobalAddressVGet"] = reflect.TypeOf(GlobalAddressVGet{})
}

// GlobalAddressVGetResponse グローバルIPアドレス/V情報取得のレスポンス
type GlobalAddressVGetResponse struct {
	*CommonResponse
	StartDate      string `json:",omitempty"` // 利用開始日(YYYYMMDD)
	ContractStatus string `json:",omitempty"` // 契約状態
	StopDate       string `json:",omitempty"` // 解約予定日(YYYYMMDD)
	AddressNum     string `json:",omitempty"` // 追加で利用できるグローバルIPアドレスの数(数値)
}

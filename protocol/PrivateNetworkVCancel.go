package protocol

import (
	"reflect"
)

// PrivateNetworkVCancel プライベートネットワーク/V解約申込 (同期)
//  http://manual.iij.jp/p2/pubapi/59940813.html
type PrivateNetworkVCancel struct {
	GisServiceCode string `json:"-"`                // P2契約のサービスコード(gis########)
	IvlServiceCode string `json:"-"`                // プライベートネットワーク/Vのサービスコード(ivl########)
	StopDate       string `json:"-" p2pub:",query"` // 解約予定日。省略した場合は即時(YYYYMMDD)
}

// URI /{{.GisServiceCode}}/private-networks/{{.IvlServiceCode}}.json
func (t PrivateNetworkVCancel) URI() string {
	return "/{{.GisServiceCode}}/private-networks/{{.IvlServiceCode}}.json"
}

// APIName PrivateNetworkVCancel
func (t PrivateNetworkVCancel) APIName() string {
	return "PrivateNetworkVCancel"
}

// Method DELETE
func (t PrivateNetworkVCancel) Method() string {
	return "DELETE"
}

// Document http://manual.iij.jp/p2/pubapi/59940813.html
func (t PrivateNetworkVCancel) Document() string {
	return "http://manual.iij.jp/p2/pubapi/59940813.html"
}

// JPName プライベートネットワーク/V解約申込
func (t PrivateNetworkVCancel) JPName() string {
	return "プライベートネットワーク/V解約申込"
}
func init() {
	APIlist = append(APIlist, PrivateNetworkVCancel{})
	TypeMap["PrivateNetworkVCancel"] = reflect.TypeOf(PrivateNetworkVCancel{})
}

// PrivateNetworkVCancelResponse プライベートネットワーク/V解約申込のレスポンス
type PrivateNetworkVCancelResponse struct {
	*CommonResponse
	Current struct {
		StopDate string `json:",omitempty"` // 設定した解約予定日(YYYYMMDD)
	} `json:",omitempty"`
	Previous struct {
		StopDate string `json:",omitempty"` // 設定前の解約予定日(YYYYMMDD)
	} `json:",omitempty"`
}

package protocol

import (
	"reflect"
)

// PrivateNetworkVAdd プライベートネットワーク/V追加申込 (同期)
//  http://manual.iij.jp/p2/pubapi/59940530.html
type PrivateNetworkVAdd struct {
	GisServiceCode string `json:"-"` // P2契約のサービスコード(gis########)
}

// URI /{{.GisServiceCode}}/private-networks.json
func (t PrivateNetworkVAdd) URI() string {
	return "/{{.GisServiceCode}}/private-networks.json"
}

// APIName PrivateNetworkVAdd
func (t PrivateNetworkVAdd) APIName() string {
	return "PrivateNetworkVAdd"
}

// Method POST
func (t PrivateNetworkVAdd) Method() string {
	return "POST"
}

// Document http://manual.iij.jp/p2/pubapi/59940530.html
func (t PrivateNetworkVAdd) Document() string {
	return "http://manual.iij.jp/p2/pubapi/59940530.html"
}

// JPName プライベートネットワーク/V追加申込
func (t PrivateNetworkVAdd) JPName() string {
	return "プライベートネットワーク/V追加申込"
}
func init() {
	APIlist = append(APIlist, PrivateNetworkVAdd{})
	TypeMap["PrivateNetworkVAdd"] = reflect.TypeOf(PrivateNetworkVAdd{})
}

// PrivateNetworkVAddResponse プライベートネットワーク/V追加申込のレスポンス
type PrivateNetworkVAddResponse struct {
	*CommonResponse
	NetworkAddress string `json:",omitempty"` // プライベートネットワーク/Vで利用されているネットワークアドレス(契約直後は設定されていないため空文字)
	Label          string `json:",omitempty"` // ラベル(文字列)
	ServiceCode    string `json:",omitempty"` // プライベートネットワーク/Vのサービスコード(ivl########)
	StartDate      string `json:",omitempty"` // 利用開始日(YYYYMMDD)
	ContractStatus string `json:",omitempty"` // 契約状態
	StopDate       string `json:",omitempty"` // 解約予定日(YYYYMMDD)
}

package protocol

import (
	"reflect"
)

// PrivateNetworkVContractGet プライベートネットワーク/V契約状態取得 (同期)
//  http://manual.iij.jp/p2/pubapi/59940802.html
type PrivateNetworkVContractGet struct {
	GisServiceCode string `json:"-"`                // P2契約のサービスコード(gis########)
	IvlServiceCode string `json:"-"`                // プライベートネットワーク/Vのサービスコード(ivl########)
	Item           string `json:"-" p2pub:",query"` // 取得するフィールド("ContractStatus")
}

// URI /{{.GisServiceCode}}/private-networks/{{.IvlServiceCode}}.json
func (t PrivateNetworkVContractGet) URI() string {
	return "/{{.GisServiceCode}}/private-networks/{{.IvlServiceCode}}.json"
}

// APIName PrivateNetworkVContractGet
func (t PrivateNetworkVContractGet) APIName() string {
	return "PrivateNetworkVContractGet"
}

// Method GET
func (t PrivateNetworkVContractGet) Method() string {
	return "GET"
}

// Document http://manual.iij.jp/p2/pubapi/59940802.html
func (t PrivateNetworkVContractGet) Document() string {
	return "http://manual.iij.jp/p2/pubapi/59940802.html"
}

// JPName プライベートネットワーク/V契約状態取得
func (t PrivateNetworkVContractGet) JPName() string {
	return "プライベートネットワーク/V契約状態取得"
}
func init() {
	APIlist = append(APIlist, PrivateNetworkVContractGet{})
	TypeMap["PrivateNetworkVContractGet"] = reflect.TypeOf(PrivateNetworkVContractGet{})
}

// PrivateNetworkVContractGetResponse プライベートネットワーク/V契約状態取得のレスポンス
type PrivateNetworkVContractGetResponse struct {
	*CommonResponse
	ContractStatus string `json:",omitempty"` // 契約状態
}

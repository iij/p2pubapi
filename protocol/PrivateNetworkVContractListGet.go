package protocol

import (
	"reflect"
)

// PrivateNetworkVContractListGet プライベートネットワーク/V契約状態一覧取得 (同期)
//  http://manual.iij.jp/p2/pubapi/59940771.html
type PrivateNetworkVContractListGet struct {
	GisServiceCode string `json:"-"`                // P2契約のサービスコード(gis########)
	Item           string `json:"-" p2pub:",query"` // 取得するフィールド("ContractStatus")
	StartIndex     string `json:"-" p2pub:",query"` // StartIndexだけスキップした位置から情報を取得する(数値)
	Count          string `json:"-" p2pub:",query"` // 取得する情報の最大数(数値)
}

// URI /{{.GisServiceCode}}/private-networks.json
func (t PrivateNetworkVContractListGet) URI() string {
	return "/{{.GisServiceCode}}/private-networks.json"
}

// APIName PrivateNetworkVContractListGet
func (t PrivateNetworkVContractListGet) APIName() string {
	return "PrivateNetworkVContractListGet"
}

// Method GET
func (t PrivateNetworkVContractListGet) Method() string {
	return "GET"
}

// Document http://manual.iij.jp/p2/pubapi/59940771.html
func (t PrivateNetworkVContractListGet) Document() string {
	return "http://manual.iij.jp/p2/pubapi/59940771.html"
}

// JPName プライベートネットワーク/V契約状態一覧取得
func (t PrivateNetworkVContractListGet) JPName() string {
	return "プライベートネットワーク/V契約状態一覧取得"
}
func init() {
	APIlist = append(APIlist, PrivateNetworkVContractListGet{})
	TypeMap["PrivateNetworkVContractListGet"] = reflect.TypeOf(PrivateNetworkVContractListGet{})
}

// PrivateNetworkVContractListGetResponse プライベートネットワーク/V契約状態一覧取得のレスポンス
type PrivateNetworkVContractListGetResponse struct {
	*CommonResponse
	PrivateNetworkList []struct {
		ContractStatus string `json:",omitempty"` // 契約状態
	} `json:",omitempty"`
}

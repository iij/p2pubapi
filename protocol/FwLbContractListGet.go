package protocol

import (
	"reflect"
)

// FwLbContractListGet FW+LB契約状態一覧取得 (同期)
//  http://manual.iij.jp/p2/pubapi/59940887.html
type FwLbContractListGet struct {
	GisServiceCode string `json:"-"`                // P2契約のサービスコード(gis########)
	Item           string `json:"-" p2pub:",query"` // 取得するフィールド("ContractStatus")
	StartIndex     string `json:"-" p2pub:",query"` // StartIndexだけスキップした位置から情報を取得する(数値)
	Count          string `json:"-" p2pub:",query"` // 取得する情報の最大数(数値)
}

// URI /{{.GisServiceCode}}/fw-lbs.json
func (t FwLbContractListGet) URI() string {
	return "/{{.GisServiceCode}}/fw-lbs.json"
}

// APIName FwLbContractListGet
func (t FwLbContractListGet) APIName() string {
	return "FwLbContractListGet"
}

// Method GET
func (t FwLbContractListGet) Method() string {
	return "GET"
}

// Document http://manual.iij.jp/p2/pubapi/59940887.html
func (t FwLbContractListGet) Document() string {
	return "http://manual.iij.jp/p2/pubapi/59940887.html"
}

// JPName FW+LB契約状態一覧取得
func (t FwLbContractListGet) JPName() string {
	return "FW+LB契約状態一覧取得"
}
func init() {
	APIlist = append(APIlist, FwLbContractListGet{})
	TypeMap["FwLbContractListGet"] = reflect.TypeOf(FwLbContractListGet{})
}

// FwLbContractListGetResponse FW+LB契約状態一覧取得のレスポンス
type FwLbContractListGetResponse struct {
	*CommonResponse
	FwLbList []struct {
		ContractStatus string `json:",omitempty"` // 契約状態
		ServiceCode    string `json:",omitempty"` // FW+LBのサービスコード(ifl########)
	} `json:",omitempty"`
}

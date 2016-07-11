package protocol

import (
	"reflect"
)

// FwLbContractGet FW+LB契約状態取得 (同期)
//  http://manual.iij.jp/p2/pubapi/59940915.html
type FwLbContractGet struct {
	GisServiceCode string `json:"-"`                // P2契約のサービスコード(gis########)
	IflServiceCode string `json:"-"`                // FW+LB 専有タイプのサービスコード(ifl########)
	Item           string `json:"-" p2pub:",query"` // 取得するフィールド("ContractStatus")
}

// URI /{{.GisServiceCode}}/fw-lbs/{{.IflServiceCode}}.json
func (t FwLbContractGet) URI() string {
	return "/{{.GisServiceCode}}/fw-lbs/{{.IflServiceCode}}.json"
}

// APIName FwLbContractGet
func (t FwLbContractGet) APIName() string {
	return "FwLbContractGet"
}

// Method GET
func (t FwLbContractGet) Method() string {
	return "GET"
}

// Document http://manual.iij.jp/p2/pubapi/59940915.html
func (t FwLbContractGet) Document() string {
	return "http://manual.iij.jp/p2/pubapi/59940915.html"
}

// JPName FW+LB契約状態取得
func (t FwLbContractGet) JPName() string {
	return "FW+LB契約状態取得"
}
func init() {
	APIlist = append(APIlist, FwLbContractGet{})
	TypeMap["FwLbContractGet"] = reflect.TypeOf(FwLbContractGet{})
}

// FwLbContractGetResponse FW+LB契約状態取得のレスポンス
type FwLbContractGetResponse struct {
	*CommonResponse
	ContractStatus string `json:",omitempty"` // 契約状態
}

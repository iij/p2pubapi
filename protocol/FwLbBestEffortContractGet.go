package protocol

import (
	"reflect"
)

// FwLbBestEffortContractGet FW+LBベストエフォートタイプ契約状態取得 (同期)
//  http://manual.iij.jp/p2/pubapi/59940915.html
type FwLbBestEffortContractGet struct {
	GisServiceCode string `json:"-"`                // P2契約のサービスコード(gis########)
	IlbServiceCode string `json:"-"`                // FW+LB ベストエフォートタイプのサービスコード(ilb########)
	Item           string `json:"-" p2pub:",query"` // 取得するフィールド("ContractStatus")
}

// URI /{{.GisServiceCode}}/best-effort-fw-lbs/{{.IlbServiceCode}}.json
func (t FwLbBestEffortContractGet) URI() string {
	return "/{{.GisServiceCode}}/best-effort-fw-lbs/{{.IlbServiceCode}}.json"
}

// APIName FwLbBestEffortContractGet
func (t FwLbBestEffortContractGet) APIName() string {
	return "FwLbBestEffortContractGet"
}

// Method GET
func (t FwLbBestEffortContractGet) Method() string {
	return "GET"
}

// Document http://manual.iij.jp/p2/pubapi/59940915.html
func (t FwLbBestEffortContractGet) Document() string {
	return "http://manual.iij.jp/p2/pubapi/59940915.html"
}

// JPName FW+LBベストエフォートタイプ契約状態取得
func (t FwLbBestEffortContractGet) JPName() string {
	return "FW+LBベストエフォートタイプ契約状態取得"
}
func init() {
	APIlist = append(APIlist, FwLbBestEffortContractGet{})
	TypeMap["FwLbBestEffortContractGet"] = reflect.TypeOf(FwLbBestEffortContractGet{})
}

// FwLbBestEffortContractGetResponse FW+LBベストエフォートタイプ契約状態取得のレスポンス
type FwLbBestEffortContractGetResponse struct {
	*CommonResponse
	ContractStatus string `json:",omitempty"` // 契約状態
}

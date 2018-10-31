package protocol

import (
	"reflect"
)

// FwLbBestEffortContractListGet FW+LBベストエフォートタイプ契約状態一覧取得 (同期)
//  http://manual.iij.jp/p2/pubapi/162300148.html
type FwLbBestEffortContractListGet struct {
	GisServiceCode string `json:"-"`                // P2契約のサービスコード(gis########)
	Item           string `json:"-" p2pub:",query"` // 取得するフィールド("ContractStatus")
	StartIndex     string `json:"-" p2pub:",query"` // StartIndexだけスキップした位置から情報を取得する(数値)
	Count          string `json:"-" p2pub:",query"` // 取得する情報の最大数(数値)
}

// URI /{{.GisServiceCode}}/best-effort-fw-lbs.json
func (t FwLbBestEffortContractListGet) URI() string {
	return "/{{.GisServiceCode}}/best-effort-fw-lbs.json"
}

// APIName FwLbBestEffortContractListGet
func (t FwLbBestEffortContractListGet) APIName() string {
	return "FwLbBestEffortContractListGet"
}

// Method GET
func (t FwLbBestEffortContractListGet) Method() string {
	return "GET"
}

// Document http://manual.iij.jp/p2/pubapi/59940887.html
func (t FwLbBestEffortContractListGet) Document() string {
	return "http://manual.iij.jp/p2/pubapi/59940887.html"
}

// JPName FW+LBベストエフォートタイプ契約状態一覧取得
func (t FwLbBestEffortContractListGet) JPName() string {
	return "FW+LBベストエフォートタイプ契約状態一覧取得"
}
func init() {
	APIlist = append(APIlist, FwLbBestEffortContractListGet{})
	TypeMap["FwLbBestEffortContractListGet"] = reflect.TypeOf(FwLbBestEffortContractListGet{})
}

// FwLbBestEffortContractListGetResponse FW+LBベストエフォートタイプ契約状態一覧取得のレスポンス
type FwLbBestEffortContractListGetResponse struct {
	*CommonResponse
	BestEffortFwLbList []struct {
		ContractStatus string `json:",omitempty"` // 契約状態
		ServiceCode    string `json:",omitempty"` // FW+LBのサービスコード(ilb########)
	} `json:",omitempty"`
}

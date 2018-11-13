package protocol

import (
	"reflect"
)

// FwLbBestEffortCancel FW+LBベストエフォートタイプ解約申込 (同期)
//  http://manual.iij.jp/p2/pubapi/162302634.html
type FwLbBestEffortCancel struct {
	GisServiceCode string `json:"-"`                // P2契約のサービスコード(gis########)
	IlbServiceCode string `json:"-"`                // FW+LB ベストエフォートタイプのサービスコード(ilb########)
	StopDate       string `json:"-" p2pub:",query"` // 解約予定日。省略の場合は即時(YYYYMMDD)
}

// URI /{{.GisServiceCode}}/best-effort-fw-lbs/{{.IlbServiceCode}}.json
func (t FwLbBestEffortCancel) URI() string {
	return "/{{.GisServiceCode}}/best-effort-fw-lbs/{{.IlbServiceCode}}.json"
}

// APIName FwLbBestEffortCancel
func (t FwLbBestEffortCancel) APIName() string {
	return "FwLbBestEffortCancel"
}

// Method DELETE
func (t FwLbBestEffortCancel) Method() string {
	return "DELETE"
}

// Document http://manual.iij.jp/p2/pubapi/162302634.html
func (t FwLbBestEffortCancel) Document() string {
	return "http://manual.iij.jp/p2/pubapi/162302634.html"
}

// JPName FW+LBベストエフォートタイプ解約申込
func (t FwLbBestEffortCancel) JPName() string {
	return "FW+LBベストエフォートタイプ解約申込"
}
func init() {
	APIlist = append(APIlist, FwLbBestEffortCancel{})
	TypeMap["FwLbBestEffortCancel"] = reflect.TypeOf(FwLbBestEffortCancel{})
}

// FwLbBestEffortCancelResponse FW+LBベストエフォートタイプ解約申込のレスポンス
type FwLbBestEffortCancelResponse struct {
	*CommonResponse
	Current struct {
		StopDate string `json:",omitempty"` // 設定した解約予定日(YYYYMMDD)
	} `json:",omitempty"`
	Previous struct {
		StopDate string `json:",omitempty"` // 設定前の解約予定日(YYYYMMDD)
	} `json:",omitempty"`
}

package protocol

import (
	"reflect"
)

// FwLbCancel FW+LB解約申込 (同期)
//  http://manual.iij.jp/p2/pubapi/59940945.html
type FwLbCancel struct {
	GisServiceCode string `json:"-"`                // P2契約のサービスコード(gis########)
	IflServiceCode string `json:"-"`                // FW+LB 専有タイプのサービスコード(ifl########)
	StopDate       string `json:"-" p2pub:",query"` // 解約予定日。省略の場合は即時(YYYYMMDD)
}

// URI /{{.GisServiceCode}}/fw-lbs/{{.IflServiceCode}}.json
func (t FwLbCancel) URI() string {
	return "/{{.GisServiceCode}}/fw-lbs/{{.IflServiceCode}}.json"
}

// APIName FwLbCancel
func (t FwLbCancel) APIName() string {
	return "FwLbCancel"
}

// Method DELETE
func (t FwLbCancel) Method() string {
	return "DELETE"
}

// Document http://manual.iij.jp/p2/pubapi/59940945.html
func (t FwLbCancel) Document() string {
	return "http://manual.iij.jp/p2/pubapi/59940945.html"
}

// JPName FW+LB解約申込
func (t FwLbCancel) JPName() string {
	return "FW+LB解約申込"
}
func init() {
	APIlist = append(APIlist, FwLbCancel{})
	TypeMap["FwLbCancel"] = reflect.TypeOf(FwLbCancel{})
}

// FwLbCancelResponse FW+LB解約申込のレスポンス
type FwLbCancelResponse struct {
	*CommonResponse
	Current struct {
		StopDate string `json:",omitempty"` // 設定した解約予定日(YYYYMMDD)
	} `json:",omitempty"`
	Previous struct {
		StopDate string `json:",omitempty"` // 設定前の解約予定日(YYYYMMDD)
	} `json:",omitempty"`
}

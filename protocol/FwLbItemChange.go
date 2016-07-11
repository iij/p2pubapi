package protocol

import (
	"reflect"
)

// FwLbItemChange FW+LB品目変更申込 (非同期)
//  http://manual.iij.jp/p2/pubapi/59940922.html
type FwLbItemChange struct {
	GisServiceCode string `json:"-"` // P2契約のサービスコード(gis########)
	IflServiceCode string `json:"-"` // FW+LB 専有タイプのサービスコード(ifl########)
	Type           string // FW+LB 専有タイプ品目
}

// URI /{{.GisServiceCode}}/fw-lbs/{{.IflServiceCode}}.json
func (t FwLbItemChange) URI() string {
	return "/{{.GisServiceCode}}/fw-lbs/{{.IflServiceCode}}.json"
}

// APIName FwLbItemChange
func (t FwLbItemChange) APIName() string {
	return "FwLbItemChange"
}

// Method PUT
func (t FwLbItemChange) Method() string {
	return "PUT"
}

// Document http://manual.iij.jp/p2/pubapi/59940922.html
func (t FwLbItemChange) Document() string {
	return "http://manual.iij.jp/p2/pubapi/59940922.html"
}

// JPName FW+LB品目変更申込
func (t FwLbItemChange) JPName() string {
	return "FW+LB品目変更申込"
}
func init() {
	APIlist = append(APIlist, FwLbItemChange{})
	TypeMap["FwLbItemChange"] = reflect.TypeOf(FwLbItemChange{})
}

// FwLbItemChangeResponse FW+LB品目変更申込のレスポンス
type FwLbItemChangeResponse struct {
	*CommonResponse
	Current struct {
		Type string `json:",omitempty"` // 設定前のFW+LB 専有タイプ品目
	} `json:",omitempty"`
	Next struct {
		Type string `json:",omitempty"` // 設定したFW+LB 専有タイプ品目
	} `json:",omitempty"`
}

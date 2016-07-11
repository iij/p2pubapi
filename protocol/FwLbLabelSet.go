package protocol

import (
	"reflect"
)

// FwLbLabelSet FW+LBラベル設定 (同期)
//  http://manual.iij.jp/p2/pubapi/59940968.html
type FwLbLabelSet struct {
	GisServiceCode string `json:"-"` // P2契約のサービスコード(gis########)
	IflServiceCode string `json:"-"` // FW+LB 専有タイプのサービスコード(ifl########)
	Name           string // ラベル(文字列)
}

// URI /{{.GisServiceCode}}/fw-lbs/{{.IflServiceCode}}/label.json
func (t FwLbLabelSet) URI() string {
	return "/{{.GisServiceCode}}/fw-lbs/{{.IflServiceCode}}/label.json"
}

// APIName FwLbLabelSet
func (t FwLbLabelSet) APIName() string {
	return "FwLbLabelSet"
}

// Method GET
func (t FwLbLabelSet) Method() string {
	return "GET"
}

// Document http://manual.iij.jp/p2/pubapi/59940968.html
func (t FwLbLabelSet) Document() string {
	return "http://manual.iij.jp/p2/pubapi/59940968.html"
}

// JPName FW+LBラベル設定
func (t FwLbLabelSet) JPName() string {
	return "FW+LBラベル設定"
}
func init() {
	APIlist = append(APIlist, FwLbLabelSet{})
	TypeMap["FwLbLabelSet"] = reflect.TypeOf(FwLbLabelSet{})
}

// FwLbLabelSetResponse FW+LBラベル設定のレスポンス
type FwLbLabelSetResponse struct {
	*CommonResponse
	Current struct {
		Label string `json:",omitempty"` // 設定したラベル(文字列)
	} `json:",omitempty"`
	Previous struct {
		Label string `json:",omitempty"` // 設定前のラベル(文字列)
	} `json:",omitempty"`
}

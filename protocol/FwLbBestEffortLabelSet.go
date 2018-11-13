package protocol

import (
	"reflect"
)

// FwLbBestEffortLabelSet FW+LBベストエフォートタイプラベル設定 (同期)
//  http://manual.iij.jp/p2/pubapi/162303320.html
type FwLbBestEffortLabelSet struct {
	GisServiceCode string `json:"-"` // P2契約のサービスコード(gis########)
	IlbServiceCode string `json:"-"` // FW+LB ベストエフォートタイプのサービスコード(ilb########)
	Name           string // ラベル(文字列)
}

// URI /{{.GisServiceCode}}/best-effort-fw-lbs/{{.IlbServiceCode}}/label.json
func (t FwLbBestEffortLabelSet) URI() string {
	return "/{{.GisServiceCode}}/best-effort-fw-lbs/{{.IlbServiceCode}}/label.json"
}

// APIName FwLbBestEffortLabelSet
func (t FwLbBestEffortLabelSet) APIName() string {
	return "FwLbBestEffortLabelSet"
}

// Method PUT
func (t FwLbBestEffortLabelSet) Method() string {
	return "PUT"
}

// Document http://manual.iij.jp/p2/pubapi/162303320.html
func (t FwLbBestEffortLabelSet) Document() string {
	return "http://manual.iij.jp/p2/pubapi/162303320.html"
}

// JPName FW+LBベストエフォートタイプラベル設定
func (t FwLbBestEffortLabelSet) JPName() string {
	return "FW+LBベストエフォートタイプラベル設定"
}
func init() {
	APIlist = append(APIlist, FwLbBestEffortLabelSet{})
	TypeMap["FwLbBestEffortLabelSet"] = reflect.TypeOf(FwLbBestEffortLabelSet{})
}

// FwLbBestEffortLabelSetResponse FW+LBベストエフォートタイプラベル設定のレスポンス
type FwLbBestEffortLabelSetResponse struct {
	*CommonResponse
	Current struct {
		Label string `json:",omitempty"` // 設定したラベル(文字列)
	} `json:",omitempty"`
	Previous struct {
		Label string `json:",omitempty"` // 設定前のラベル(文字列)
	} `json:",omitempty"`
}

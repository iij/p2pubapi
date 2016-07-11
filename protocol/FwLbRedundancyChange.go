package protocol

import (
	"reflect"
)

// FwLbRedundancyChange FW+LB冗長構成変更申込 (非同期)
//  http://manual.iij.jp/p2/pubapi/59940935.html
type FwLbRedundancyChange struct {
	GisServiceCode string `json:"-"` // P2契約のサービスコード(gis########)
	IflServiceCode string `json:"-"` // FW+LB 専有タイプのサービスコード(ifl########)
	Redundant      string // 冗長構成有無(Yes" "No)
}

// URI /{{.GisServiceCode}}/fw-lbs/{{.IflServiceCode}}.json
func (t FwLbRedundancyChange) URI() string {
	return "/{{.GisServiceCode}}/fw-lbs/{{.IflServiceCode}}.json"
}

// APIName FwLbRedundancyChange
func (t FwLbRedundancyChange) APIName() string {
	return "FwLbRedundancyChange"
}

// Method PUT
func (t FwLbRedundancyChange) Method() string {
	return "PUT"
}

// Document http://manual.iij.jp/p2/pubapi/59940935.html
func (t FwLbRedundancyChange) Document() string {
	return "http://manual.iij.jp/p2/pubapi/59940935.html"
}

// JPName FW+LB冗長構成変更申込
func (t FwLbRedundancyChange) JPName() string {
	return "FW+LB冗長構成変更申込"
}
func init() {
	APIlist = append(APIlist, FwLbRedundancyChange{})
	TypeMap["FwLbRedundancyChange"] = reflect.TypeOf(FwLbRedundancyChange{})
}

// FwLbRedundancyChangeResponse FW+LB冗長構成変更申込のレスポンス
type FwLbRedundancyChangeResponse struct {
	*CommonResponse
	Current struct {
		Redundant string `json:",omitempty"` // 設定前の冗長構成有無(Yes" "No)
	} `json:",omitempty"`
	Next struct {
		Redundant string `json:",omitempty"` // 設定した冗長構成有無(Yes" "No)
	} `json:",omitempty"`
}

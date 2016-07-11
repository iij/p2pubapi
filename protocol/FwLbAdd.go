package protocol

import (
	"reflect"
)

// FwLbAdd FW+LB追加申込 (非同期)
//  http://manual.iij.jp/p2/pubapi/59940827.html
type FwLbAdd struct {
	GisServiceCode string `json:"-"` // P2契約のサービスコード(gis########)
	Redundant      string // 冗長構成有無(Yes" "No)
	Type           string // FW+LB 専有タイプ品目
}

// URI /{{.GisServiceCode}}/fw-lbs.json
func (t FwLbAdd) URI() string {
	return "/{{.GisServiceCode}}/fw-lbs.json"
}

// APIName FwLbAdd
func (t FwLbAdd) APIName() string {
	return "FwLbAdd"
}

// Method POST
func (t FwLbAdd) Method() string {
	return "POST"
}

// Document http://manual.iij.jp/p2/pubapi/59940827.html
func (t FwLbAdd) Document() string {
	return "http://manual.iij.jp/p2/pubapi/59940827.html"
}

// JPName FW+LB追加申込
func (t FwLbAdd) JPName() string {
	return "FW+LB追加申込"
}
func init() {
	APIlist = append(APIlist, FwLbAdd{})
	TypeMap["FwLbAdd"] = reflect.TypeOf(FwLbAdd{})
}

// FwLbAddResponse FW+LB追加申込のレスポンス
type FwLbAddResponse struct {
	*CommonResponse
	ResourceStatus string `json:",omitempty"` // FW+LB 専有タイプステータス
	HostList       []struct {
		Master                    string `json:",omitempty"` // マスターならばYes, スレーブならばNo(Yes" "No)
		ResourceStatus            string `json:",omitempty"` // FW+LB 専有タイプステータス
		LbSoftwareVersion         string `json:",omitempty"` // LBソフトウェアバージョン(例： "9.9)
		LbAdministrationServerUrl string `json:",omitempty"` // LB管理サーバURL(例： https://lbcNNNNN.lb.pub.p2.iijgio.jp)
	} `json:",omitempty"`
	Label          string `json:",omitempty"` // ラベル(文字列)
	ServiceCode    string `json:",omitempty"` // FW+LB 専有タイプのサービスコード(ifl########)
	Redundant      string `json:",omitempty"` // 冗長構成有無(Yes" "No)
	StartDate      string `json:",omitempty"` // 利用開始日(YYYYMMDD)
	ContractStatus string `json:",omitempty"` // 契約状態
	StopDate       string `json:",omitempty"` // 解約予定日(YYYYMMDD)
	Type           string `json:",omitempty"` // FW+LB 専有タイプ品目
}

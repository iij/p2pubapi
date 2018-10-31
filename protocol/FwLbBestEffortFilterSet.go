package protocol

import (
	"reflect"
)

// FwLbBestEffortFilterSet フィルタリングルール一括設定（FW+LBベストエフォートタイプ） (同期)
//  http://manual.iij.jp/p2/pubapi/162646983.html
type FwLbBestEffortFilterSet struct {
	GisServiceCode string       `json:"-"` // P2契約のサービスコード(gis########)
	IlbServiceCode string       `json:"-"` // FW+LB ベストエフォートタイプのサービスコード(ilb########)
	IpVersion      string       `json:"-"` // IPv4ルール（V4）
	Direction      string       `json:"-"` // 入力方向（ExternalからInternalへ）のルールか出力方向のルールか（in, out）
	FilterRuleList []FilterRule `json:",omitempty"`
}

// URI /:GisServiceCode/best-effort-fw-lbs/:IlbServiceCode/filters/:IpVersion/:Direction.json
func (t FwLbBestEffortFilterSet) URI() string {
	return "/{{.GisServiceCode}}/best-effort-fw-lbs/{{.IlbServiceCode}}/filters/{{.IpVersion}}/{{.Direction}}.json"
}

// APIName FwLbBestEffortFilterSet
func (t FwLbBestEffortFilterSet) APIName() string {
	return "FwLbBestEffortFilterSet"
}

// Method GET
func (t FwLbBestEffortFilterSet) Method() string {
	return "PUT"
}

// http://manual.iij.jp/p2/pubapi/162646983.html
func (t FwLbBestEffortFilterSet) Document() string {
	return "http://manual.iij.jp/p2/pubapi/162646983.html"
}

// JPName フィルタリングルール一括設定（FW+LBベストエフォートタイプ）
func (t FwLbBestEffortFilterSet) JPName() string {
	return "フィルタリングルール一括設定（FW+LBベストエフォートタイプ）"
}
func init() {
	APIlist = append(APIlist, FwLbBestEffortFilterSet{})
	TypeMap["FwLbBestEffortFilterSet"] = reflect.TypeOf(FwLbBestEffortFilterSet{})
}

// FwLbBestEffortFilterSetResponse フィルタリングルール一括設定（FW+LBベストエフォートタイプ）のレスポンス
type FwLbBestEffortFilterSetResponse struct {
	*CommonResponse
	Current struct {
		FilterList []struct {
			FilterId           string `json:",omitempty"`
			SourceNetwork      string `json:",omitempty"`
			DestinationNetwork string `json:",omitempty"`
			DestinationPort    string `json:",omitempty"`
			Protocol           string `json:",omitempty"`
			Action             string `json:",omitempty"`
			Label              string `json:",omitempty"`
		} `json:",omitempty"`
	}
	Previous struct {
		FilterList []struct {
			FilterId           string `json:",omitempty"`
			SourceNetwork      string `json:",omitempty"`
			DestinationNetwork string `json:",omitempty"`
			DestinationPort    string `json:",omitempty"`
			Protocol           string `json:",omitempty"`
			Action             string `json:",omitempty"`
			Label              string `json:",omitempty"`
		} `json:",omitempty"`
	}
}

package protocol

import (
	"reflect"
)

type FilterRule struct {
	SourceNetwork      string
	DestinationNetwork string
	DestinationPort    string
	Protocol           string
	Action             string
	Label              string
}

// FwLbFilterSet フィルタリングルール一括設定 (同期)
//  http://manual.iij.jp/p2/pubapi/73641928.html
type FwLbFilterSet struct {
	GisServiceCode string       `json:"-"` // P2契約のサービスコード(gis########)
	IflServiceCode string       `json:"-"` // FW+LB 専有タイプのサービスコード(ifl########)
	IpVersion      string       `json:"-"` // IPv4ルールかIPv6ルール（V4, V6）
	Direction      string       `json:"-"` // 入力方向（ExternalからInternalへ）のルールか出力方向のルールか。IPv6は"in"のみ指定可能（in, out）
	FilterRuleList []FilterRule
}

// URI /:GisServiceCode/fw-lbs/:IflServiceCode/filters/:IpVersion/:Direction.json
func (t FwLbFilterSet) URI() string {
	return "/{{.GisServiceCode}}/fw-lbs/{{.IflServiceCode}}/filters/{{.IpVersion}}/{{.Direction}}.json"
}

// APIName FwLbFilterSet
func (t FwLbFilterSet) APIName() string {
	return "FwLbFilterSet"
}

// Method GET
func (t FwLbFilterSet) Method() string {
	return "PUT"
}

// http://manual.iij.jp/p2/pubapi/73641928.html
func (t FwLbFilterSet) Document() string {
	return "http://manual.iij.jp/p2/pubapi/73641928.html"
}

// JPName フィルタリングルール一括設定
func (t FwLbFilterSet) JPName() string {
	return "フィルタリングルール一括設定"
}
func init() {
	APIlist = append(APIlist, FwLbFilterSet{})
	TypeMap["FwLbFilterSet"] = reflect.TypeOf(FwLbFilterSet{})
}

// FwLbFilterSetResponse フィルタリングルール一括設定のレスポンス
type FwLbFilterSetResponse struct {
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

package protocol

import (
	"reflect"
)

// FwLbGet FW+LB情報取得 (同期)
//  http://manual.iij.jp/p2/pubapi/59940904.html
type FwLbFilterGet struct {
	GisServiceCode string `json:"-"` // P2契約のサービスコード(gis########)
	IflServiceCode string `json:"-"` // FW+LB 専有タイプのサービスコード(ifl########)
	IpVersion      string `json:"-"` // IPv4ルールかIPv6ルール（V4, V6）
	Direction      string `json:"-"` // 入力方向（ExternalからInternalへ）のルールか出力方向のルールか。IPv6は"in"のみ指定可能（in, out）
}

// URI /:GisServiceCode/fw-lbs/:IflServiceCode/filters/:IpVersion/:Direction.json
func (t FwLbFilterGet) URI() string {
	return "/{{.GisServiceCode}}/fw-lbs/{{.IflServiceCode}}/filters/{{.IpVersion}}/{{.Direction}}.json"
}

// APIName FwLbFilterGet
func (t FwLbFilterGet) APIName() string {
	return "FwLbFilterGet"
}

// Method GET
func (t FwLbFilterGet) Method() string {
	return "GET"
}

// http://manual.iij.jp/p2/pubapi/73641921.html
func (t FwLbFilterGet) Document() string {
	return "http://manual.iij.jp/p2/pubapi/73641921.html"
}

// JPName フィルタリングルール情報取得
func (t FwLbFilterGet) JPName() string {
	return "フィルタリングルール情報取得"
}
func init() {
	APIlist = append(APIlist, FwLbFilterGet{})
	TypeMap["FwLbFilterGet"] = reflect.TypeOf(FwLbFilterGet{})
}

// FwLbFilterGetResponse フィルタリングルール情報取得のレスポンス
type FwLbFilterGetResponse struct {
	*CommonResponse
	FilterRuleList []struct {
		FilterId           string `json:",omitempty"`
		SourceNetwork      string `json:",omitempty"`
		DestinationNetwork string `json:",omitempty"`
		DestinationPort    string `json:",omitempty"`
		Protocol           string `json:",omitempty"`
		Action             string `json:",omitempty"`
		Label              string `json:",omitempty"`
	} `json:",omitempty"`
}

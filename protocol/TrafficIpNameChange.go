package protocol

import (
	"reflect"
)

// TrafficIpNameChange TrafficIp名称変更 (同期)
//  http://manual.iij.jp/p2/pubapi/59941058.html
type TrafficIpNameChange struct {
	GisServiceCode string `json:"-"` // P2契約のサービスコード(gis########)
	IflServiceCode string `json:"-"` // FW+LB 専有タイプのサービスコード(ifl########)
	IpAddress      string `json:"-"` // トラフィックIPアドレス(IPアドレス)
	TrafficIpName  string // トラフィックIP名(文字列)
}

// URI /{{.GisServiceCode}}/fw-lbs/{{.IflServiceCode}}/lb/trafficips/{{.IpAddress}}.json
func (t TrafficIpNameChange) URI() string {
	return "/{{.GisServiceCode}}/fw-lbs/{{.IflServiceCode}}/lb/trafficips/{{.IpAddress}}.json"
}

// APIName TrafficIpNameChange
func (t TrafficIpNameChange) APIName() string {
	return "TrafficIpNameChange"
}

// Method PUT
func (t TrafficIpNameChange) Method() string {
	return "PUT"
}

// Document http://manual.iij.jp/p2/pubapi/59941058.html
func (t TrafficIpNameChange) Document() string {
	return "http://manual.iij.jp/p2/pubapi/59941058.html"
}

// JPName TrafficIp名称変更
func (t TrafficIpNameChange) JPName() string {
	return "TrafficIp名称変更"
}
func init() {
	APIlist = append(APIlist, TrafficIpNameChange{})
	TypeMap["TrafficIpNameChange"] = reflect.TypeOf(TrafficIpNameChange{})
}

// TrafficIpNameChangeResponse TrafficIp名称変更のレスポンス
type TrafficIpNameChangeResponse struct {
	*CommonResponse
	Current struct {
		IPv4 struct {
			TrafficIpName string `json:",omitempty"` // 設定したトラフィックIP名(文字列)
		} `json:",omitempty"`
		IPv6 struct {
			TrafficIpName string `json:",omitempty"` // 設定したトラフィックIP名(文字列)
		} `json:",omitempty"`
	} `json:",omitempty"`
	Previous struct {
		IPv4 struct {
			TrafficIpName string `json:",omitempty"` // 設定前のトラフィックIP名(文字列)
		} `json:",omitempty"`
		IPv6 struct {
			TrafficIpName string `json:",omitempty"` // 設定前のトラフィックIP名(文字列)
		} `json:",omitempty"`
	} `json:",omitempty"`
}

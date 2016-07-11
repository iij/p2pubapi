package protocol

import (
	"reflect"
)

// TrafficIpGet TrafficIp情報取得 (同期)
//  http://manual.iij.jp/p2/pubapi/59941044.html
type TrafficIpGet struct {
	GisServiceCode string `json:"-"` // P2契約のサービスコード(gis########)
	IflServiceCode string `json:"-"` // FW+LB 専有タイプのサービスコード(ifl########)
	IpAddress      string `json:"-"` // トラフィックIPアドレス(IPアドレス)
}

// URI /{{.GisServiceCode}}/fw-lbs/{{.IflServiceCode}}/lb/trafficips/{{.IpAddress}}.json
func (t TrafficIpGet) URI() string {
	return "/{{.GisServiceCode}}/fw-lbs/{{.IflServiceCode}}/lb/trafficips/{{.IpAddress}}.json"
}

// APIName TrafficIpGet
func (t TrafficIpGet) APIName() string {
	return "TrafficIpGet"
}

// Method GET
func (t TrafficIpGet) Method() string {
	return "GET"
}

// Document http://manual.iij.jp/p2/pubapi/59941044.html
func (t TrafficIpGet) Document() string {
	return "http://manual.iij.jp/p2/pubapi/59941044.html"
}

// JPName TrafficIp情報取得
func (t TrafficIpGet) JPName() string {
	return "TrafficIp情報取得"
}
func init() {
	APIlist = append(APIlist, TrafficIpGet{})
	TypeMap["TrafficIpGet"] = reflect.TypeOf(TrafficIpGet{})
}

// TrafficIpGetResponse TrafficIp情報取得のレスポンス
type TrafficIpGetResponse struct {
	*CommonResponse
	NetworkType string `json:",omitempty"` // ネットワーク種別(Gobal/PrivateStandard/Private)
	IPv6        struct {
		TrafficIpName    string `json:",omitempty"` // トラフィックIP名(文字列)
		TrafficIpAddress string `json:",omitempty"` // トラフィックIPアドレス(IPv6アドレス)
	} `json:",omitempty"`
	IPv4 struct {
		TrafficIpName    string `json:",omitempty"` // トラフィックIP名(文字列)
		TrafficIpAddress string `json:",omitempty"` // トラフィックIPアドレス(IPv4アドレス)
		DomainName       string `json:",omitempty"` // 逆引きドメイン名(文字列)
	} `json:",omitempty"`
}

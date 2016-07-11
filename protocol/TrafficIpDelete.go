package protocol

import (
	"reflect"
)

// TrafficIpDelete TrafficIp削除 (同期)
//  http://manual.iij.jp/p2/pubapi/59941073.html
type TrafficIpDelete struct {
	GisServiceCode string `json:"-"` // P2契約のサービスコード(gis########)
	IflServiceCode string `json:"-"` // FW+LB 専有タイプのサービスコード(ifl########)
	IpAddress      string `json:"-"` // トラフィックIPアドレス(IPアドレス)
}

// URI /{{.GisServiceCode}}/fw-lbs/{{.IflServiceCode}}/lb/trafficips/{{.IpAddress}}.json
func (t TrafficIpDelete) URI() string {
	return "/{{.GisServiceCode}}/fw-lbs/{{.IflServiceCode}}/lb/trafficips/{{.IpAddress}}.json"
}

// APIName TrafficIpDelete
func (t TrafficIpDelete) APIName() string {
	return "TrafficIpDelete"
}

// Method DELETE
func (t TrafficIpDelete) Method() string {
	return "DELETE"
}

// Document http://manual.iij.jp/p2/pubapi/59941073.html
func (t TrafficIpDelete) Document() string {
	return "http://manual.iij.jp/p2/pubapi/59941073.html"
}

// JPName TrafficIp削除
func (t TrafficIpDelete) JPName() string {
	return "TrafficIp削除"
}
func init() {
	APIlist = append(APIlist, TrafficIpDelete{})
	TypeMap["TrafficIpDelete"] = reflect.TypeOf(TrafficIpDelete{})
}

// TrafficIpDeleteResponse TrafficIp削除のレスポンス
type TrafficIpDeleteResponse struct {
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

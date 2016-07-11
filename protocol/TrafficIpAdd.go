package protocol

import (
	"reflect"
)

// TrafficIpAdd TrafficIp追加 (同期)
//  http://manual.iij.jp/p2/pubapi/59941022.html
type TrafficIpAdd struct {
	GisServiceCode   string `json:"-"` // P2契約のサービスコード(gis########)
	IflServiceCode   string `json:"-"` // FW+LB 専有タイプのサービスコード(ifl########)
	TrafficIpName    string // トラフィックIP名(文字列)
	TrafficIpAddress string `json:",omitempty"` // トラフィックIPアドレス(IPv4アドレス)
}

// URI /{{.GisServiceCode}}/fw-lbs/{{.IflServiceCode}}/lb/trafficips.json
func (t TrafficIpAdd) URI() string {
	return "/{{.GisServiceCode}}/fw-lbs/{{.IflServiceCode}}/lb/trafficips.json"
}

// APIName TrafficIpAdd
func (t TrafficIpAdd) APIName() string {
	return "TrafficIpAdd"
}

// Method POST
func (t TrafficIpAdd) Method() string {
	return "POST"
}

// Document http://manual.iij.jp/p2/pubapi/59941022.html
func (t TrafficIpAdd) Document() string {
	return "http://manual.iij.jp/p2/pubapi/59941022.html"
}

// JPName TrafficIp追加
func (t TrafficIpAdd) JPName() string {
	return "TrafficIp追加"
}
func init() {
	APIlist = append(APIlist, TrafficIpAdd{})
	TypeMap["TrafficIpAdd"] = reflect.TypeOf(TrafficIpAdd{})
}

// TrafficIpAddResponse TrafficIp追加のレスポンス
type TrafficIpAddResponse struct {
	*CommonResponse
	NetworkType string `json:",omitempty"` // ネットワーク種別(Gobal/Private"（提供準備中）/PrivateStandard)
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

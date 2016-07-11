package protocol

import (
	"reflect"
)

// TrafficIpListGet TrafficIp情報一覧取得 (同期)
//  http://manual.iij.jp/p2/pubapi/59941037.html
type TrafficIpListGet struct {
	GisServiceCode string `json:"-"` // P2契約のサービスコード(gis########)
	IflServiceCode string `json:"-"` // FW+LB 専有タイプのサービスコード(ifl########)
}

// URI /{{.GisServiceCode}}/fw-lbs/{{.IflServiceCode}}/lb/trafficips.json
func (t TrafficIpListGet) URI() string {
	return "/{{.GisServiceCode}}/fw-lbs/{{.IflServiceCode}}/lb/trafficips.json"
}

// APIName TrafficIpListGet
func (t TrafficIpListGet) APIName() string {
	return "TrafficIpListGet"
}

// Method GET
func (t TrafficIpListGet) Method() string {
	return "GET"
}

// Document http://manual.iij.jp/p2/pubapi/59941037.html
func (t TrafficIpListGet) Document() string {
	return "http://manual.iij.jp/p2/pubapi/59941037.html"
}

// JPName TrafficIp情報一覧取得
func (t TrafficIpListGet) JPName() string {
	return "TrafficIp情報一覧取得"
}
func init() {
	APIlist = append(APIlist, TrafficIpListGet{})
	TypeMap["TrafficIpListGet"] = reflect.TypeOf(TrafficIpListGet{})
}

// TrafficIpListGetResponse TrafficIp情報一覧取得のレスポンス
type TrafficIpListGetResponse struct {
	*CommonResponse
	TrafficIpList []struct {
		IPv4 struct {
			TrafficIpName    string `json:",omitempty"` // トラフィックIP名(文字列)
			TrafficIpAddress string `json:",omitempty"` // トラフィックIPアドレス(IPv4アドレス)
			DomainName       string `json:",omitempty"` // 逆引きドメイン名(文字列)
		} `json:",omitempty"`
		IPv6 struct {
			TrafficIpName    string `json:",omitempty"` // トラフィックIP名(文字列)
			TrafficIpAddress string `json:",omitempty"` // トラフィックIPアドレス(IPv6アドレス)
		} `json:",omitempty"`
	} `json:",omitempty"`
	NetworkType string `json:",omitempty"` // ネットワーク種別(Gobal/PrivateStandard/Private)
}

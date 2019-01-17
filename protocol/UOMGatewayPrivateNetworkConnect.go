package protocol

import (
	"reflect"
)

// UOMGatewayPrivateNetworkConnect プライベートネットワーク接続（監視・運用ゲートウェイ）
// http://manual.iij.jp/p2/pubapi/79743118.html
type UOMGatewayPrivateNetworkConnect struct {
	GisServiceCode string `json:"-"` // P2契約のサービスコード(gis########)
	ServiceCode    string            // 接続するネットワークのサービスコード
	IPv4 struct {  // 監視・運用ゲートウェイに割り当てるゲートウェイアドレス
		IpAddress string // ゲートウェイアドレス
		Netmask   string // ゲートウェイのネットマスク長
	}
}

// URI /{{.GisServiceCode}}/uom-gateway/private-networks.json
func (t UOMGatewayPrivateNetworkConnect) URI() string {
	return "/{{.GisServiceCode}}/uom-gateway/private-networks.json"
}

// APIName UOMGatewayPrivateNetworkConnect
func (t UOMGatewayPrivateNetworkConnect) APIName() string {
	return "UOMGatewayPrivateNetworkConnect"
}

// Method POST
func (t UOMGatewayPrivateNetworkConnect) Method() string {
	return "POST"
}

// Document http://manual.iij.jp/p2/pubapi/79743118.html
func (t UOMGatewayPrivateNetworkConnect) Document() string {
	return "http://manual.iij.jp/p2/pubapi/79743118.html"
}

// JPName プライベートネットワーク接続（監視・運用ゲートウェイ）
func (t UOMGatewayPrivateNetworkConnect) JPName() string {
	return "プライベートネットワーク接続（監視・運用ゲートウェイ）"
}
func init() {
	APIlist = append(APIlist, UOMGatewayPrivateNetworkConnect{})
	TypeMap["UOMGatewayPrivateNetworkConnect"] = reflect.TypeOf(UOMGatewayPrivateNetworkConnect{})
}

// UOMGatewayPrivateNetworkConnectResponse プライベートネットワーク接続（監視・運用ゲートウェイ）のレスポンス
type UOMGatewayPrivateNetworkConnectResponse struct {
	*CommonResponse
	ResourceStatus string `json:",omitempty"` // 監視・運用ゲートウェイのステータス
	ServiceCode    string `json:",omitempty"` // 接続されているネットワークのサービスコード
	IPv4 []struct {  // 監視・運用ゲートウェイに割り当てられたゲートウェイアドレス
		IpAddress string `json:",omitempty"`  // ゲートウェイアドレス
		Netmask   string `json:",omitempty"`  // ゲートウェイのネットマスク長
	} `json:",omitempty"`
}

package protocol

import (
	"reflect"
)

// UOMGatewayPrivateNetworkDisconnect プライベートネットワーク切断（監視・運用ゲートウェイ）
// http://manual.iij.jp/p2/pubapi/79743124.html
type UOMGatewayPrivateNetworkDisconnect struct {
	GisServiceCode string `json:"-"` // P2契約のサービスコード(gis########)
	IvlServiceCode string `json:"-"` // プライベートネットワークのサービスコード
}

// URI /{{.GisServiceCode}}/uom-gateway/private-networks.json
func (t UOMGatewayPrivateNetworkDisconnect) URI() string {
	return "/{{.GisServiceCode}}/uom-gateway/private-networks.json"
}

// APIName UOMGatewayPrivateNetworkDisconnect
func (t UOMGatewayPrivateNetworkDisconnect) APIName() string {
	return "UOMGatewayPrivateNetworkDisconnect"
}

// Method DELETE
func (t UOMGatewayPrivateNetworkDisconnect) Method() string {
	return "DELETE"
}

// Document http://manual.iij.jp/p2/pubapi/79743124.html
func (t UOMGatewayPrivateNetworkDisconnect) Document() string {
	return "http://manual.iij.jp/p2/pubapi/79743124.html"
}

// JPName プライベートネットワーク切断（監視・運用ゲートウェイ）
func (t UOMGatewayPrivateNetworkDisconnect) JPName() string {
	return "プライベートネットワーク切断（監視・運用ゲートウェイ）"
}
func init() {
	APIlist = append(APIlist, UOMGatewayPrivateNetworkDisconnect{})
	TypeMap["UOMGatewayPrivateNetworkDisconnect"] = reflect.TypeOf(UOMGatewayPrivateNetworkDisconnect{})
}

// UOMGatewayPrivateNetworkDisconnectResponse プライベートネットワーク切断（監視・運用ゲートウェイ）のレスポンス
type UOMGatewayPrivateNetworkDisconnectResponse struct {
	*CommonResponse
	ResourceStatus string `json:",omitempty"` // 監視・運用ゲートウェイのステータス
	ServiceCode    string `json:",omitempty"` // 接続されているネットワークのサービスコード
	IPv4 []struct {  // 監視・運用ゲートウェイに割り当てられたゲートウェイアドレス
		IpAddress string `json:",omitempty"`  // ゲートウェイアドレス
		Netmask   string `json:",omitempty"`  // ゲートウェイのネットマスク長
	} `json:",omitempty"`
}

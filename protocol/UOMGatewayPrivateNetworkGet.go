package protocol

import (
	"reflect"
)

// UOMGatewayPrivateNetworksGet プライベートネットワーク情報取得（監視・運用ゲートウェイ）
// http://manual.iij.jp/p2/pubapi/79743106.html
type UOMGatewayPrivateNetworkGet struct {
	GisServiceCode string `json:"-"` // P2契約のサービスコード(gis########)
	IvlServiceCode string `json:"-"` // プライベートネットワークのサービスコード
}

// URI /{{.GisServiceCode}}/uom-gateway/private-networks//{{.IvlServiceCode}}.json
func (t UOMGatewayPrivateNetworkGet) URI() string {
	return "/{{.GisServiceCode}}/uom-gateway/private-networks//{{.IvlServiceCode}}.json"
}

// APIName UOMGatewayPrivateNetworkGet
func (t UOMGatewayPrivateNetworkGet) APIName() string {
	return "UOMGatewayPrivateNetworkGet"
}

// Method GET
func (t UOMGatewayPrivateNetworkGet) Method() string {
	return "GET"
}

// Document http://manual.iij.jp/p2/pubapi/79743106.html
func (t UOMGatewayPrivateNetworkGet) Document() string {
	return "http://manual.iij.jp/p2/pubapi/79743106.html"
}

// JPName プライベートネットワーク情報取得（監視・運用ゲートウェイ）
func (t UOMGatewayPrivateNetworkGet) JPName() string {
	return "プライベートネットワーク情報取得（監視・運用ゲートウェイ）"
}
func init() {
	APIlist = append(APIlist, UOMGatewayPrivateNetworkGet{})
	TypeMap["UOMGatewayPrivateNetworkGet"] = reflect.TypeOf(UOMGatewayPrivateNetworkGet{})
}

// UOMGatewayPrivateNetworkGetResponse プライベートネットワーク情報取得（監視・運用ゲートウェイ）のレスポンス
type UOMGatewayPrivateNetworkGetResponse struct {
	*CommonResponse
	URI         string `json:",omitempty"`  // 該当ネットワークの情報へアクセスするためのURI
	IPv4 []struct {  // 監視・運用ゲートウェイに割り当てられたゲートウェイアドレス
		IpAddress string `json:",omitempty"`  // ゲートウェイアドレス
		Netmask   string `json:",omitempty"`  // ゲートウェイのネットマスク長
	} `json:",omitempty"`
}

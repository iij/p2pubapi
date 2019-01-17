package protocol

import (
	"reflect"
)

// UOMGatewayPrivateNetworkListGet プライベートネットワーク情報一覧取得（監視・運用ゲートウェイ）
// http://manual.iij.jp/p2/pubapi/79743102.html
type UOMGatewayPrivateNetworkListGet struct {
	GisServiceCode string `json:"-"` // P2契約のサービスコード(gis########)
}

// URI /{{.GisServiceCode}}/uom-gateway/private-networks.json
func (t UOMGatewayPrivateNetworkListGet) URI() string {
	return "/{{.GisServiceCode}}/uom-gateway/private-networks.json"
}

// APIName UOMGatewayPrivateNetworkListGet
func (t UOMGatewayPrivateNetworkListGet) APIName() string {
	return "UOMGatewayPrivateNetworkListGet"
}

// Method GET
func (t UOMGatewayPrivateNetworkListGet) Method() string {
	return "GET"
}

// Document http://manual.iij.jp/p2/pubapi/79743102.html
func (t UOMGatewayPrivateNetworkListGet) Document() string {
	return "http://manual.iij.jp/p2/pubapi/79743102.html"
}

// JPName プライベートネットワーク情報一覧取得（監視・運用ゲートウェイ）
func (t UOMGatewayPrivateNetworkListGet) JPName() string {
	return "プライベートネットワーク情報一覧取得（監視・運用ゲートウェイ）"
}
func init() {
	APIlist = append(APIlist, UOMGatewayPrivateNetworkListGet{})
	TypeMap["UOMGatewayPrivateNetworkListGet"] = reflect.TypeOf(UOMGatewayPrivateNetworkListGet{})
}

// UOMGatewayPrivateNetworkListGetResponse プライベートネットワーク情報一覧取得（監視・運用ゲートウェイ）のレスポンス
type UOMGatewayPrivateNetworkListGetResponse struct {
	*CommonResponse
	PrivateNetworkList []struct {  // 接続されているネットワークの一覧
		ServiceCode string `json:",omitempty"`  // 接続されているネットワークのサービスコード
		URI         string `json:",omitempty"`  // 該当ネットワークの情報へアクセスするためのURI
		IPv4 []struct {  // 監視・運用ゲートウェイに割り当てられたゲートウェイアドレス
			IpAddress string `json:",omitempty"`  // ゲートウェイアドレス
			Netmask   string `json:",omitempty"`  // ゲートウェイのネットマスク長
		} `json:",omitempty"`
	} `json:",omitempty"`
}

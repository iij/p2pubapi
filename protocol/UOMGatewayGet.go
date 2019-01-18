package protocol

import (
	"reflect"
)

// UOMGatewayGet 監視・運用ゲートウェイ情報取得
// http://manual.iij.jp/p2/pubapi/79743087.html
type UOMGatewayGet struct {
	GisServiceCode string `json:"-"` // P2契約のサービスコード(gis########)
}

// URI /{{.GisServiceCode}}/uom-gateway.json
func (t UOMGatewayGet) URI() string {
	return "/{{.GisServiceCode}}/uom-gateway.json"
}

// APIName UOMGatewayGet
func (t UOMGatewayGet) APIName() string {
	return "UOMGatewayGet"
}

// Method GET
func (t UOMGatewayGet) Method() string {
	return "GET"
}

// Document http://manual.iij.jp/p2/pubapi/79743087.html
func (t UOMGatewayGet) Document() string {
	return "http://manual.iij.jp/p2/pubapi/79743087.html"
}

// JPName データデバイスストレージ情報取得
func (t UOMGatewayGet) JPName() string {
	return "監視・運用ゲートウェイ情報取得"
}
func init() {
	APIlist = append(APIlist, UOMGatewayGet{})
	TypeMap["UOMGatewayGet"] = reflect.TypeOf(UOMGatewayGet{})
}

// UOMGatewayGetResponse 監視・運用ゲートウェイ情報取得
type UOMGatewayGetResponse struct {
	*CommonResponse
	NetworkList []struct {  // 接続されているネットワークの一覧
		NetworkType string `json:",omitempty"`  // 接続されているネットワークの種別 "PrivateStandard", "Private"
		ServiceCode string `json:",omitempty"`  // 接続されているネットワークのサービスコード
		IPv4 []struct {  // 監視・運用ゲートウェイに割り当てられたゲートウェイアドレス
			IpAddress string `json:",omitempty"`  // ゲートウェイアドレス
			Netmask   string `json:",omitempty"`  // ゲートウェイのネットマスク長
		} `json:",omitempty"`
	} `json:",omitempty"`
}


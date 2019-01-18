package protocol

import (
	"reflect"
)

// UOMGatewayStatusGet 監視・運用ゲートウェイリソース状態取得
// http://manual.iij.jp/p2/pubapi/79743115.html
type UOMGatewayStatusGet struct {
	GisServiceCode string `json:"-"` // P2契約のサービスコード(gis########)
}

// URI /{{.GisServiceCode}}/uom-gateway.json?Item=ResourceStatus
func (t UOMGatewayStatusGet) URI() string {
	return "/{{.GisServiceCode}}/uom-gateway.json?Item=ResourceStatus"
}

// APIName UOMGatewayStatusGet
func (t UOMGatewayStatusGet) APIName() string {
	return "UOMGatewayStatusGet"
}

// Method GET
func (t UOMGatewayStatusGet) Method() string {
	return "GET"
}

// Document http://manual.iij.jp/p2/pubapi/79743115.html
func (t UOMGatewayStatusGet) Document() string {
	return "http://manual.iij.jp/p2/pubapi/79743115.html"
}

// JPName 監視・運用ゲートウェイリソース状態取得
func (t UOMGatewayStatusGet) JPName() string {
	return "監視・運用ゲートウェイリソース状態取得"
}
func init() {
	APIlist = append(APIlist, UOMGatewayStatusGet{})
	TypeMap["UOMGatewayStatusGet"] = reflect.TypeOf(UOMGatewayStatusGet{})
}

// UOMGatewayStatusGetResponse 監視・運用ゲートウェイリソース状態取得 レスポンス
type UOMGatewayStatusGetResponse struct {
	*CommonResponse
	ResourceStatus string `json:",omitempty"`  // 監視・運用ゲートウェイのステータス
}


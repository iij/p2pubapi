package protocol

import (
	"reflect"
)

// UOMGatewayMonitoringIPAddressListGet 監視用IPアドレス情報一覧取得
// http://manual.iij.jp/p2/pubapi/79743140.html
type UOMGatewayMonitoringIPAddressListGet struct {
	GisServiceCode string `json:"-"` // P2契約のサービスコード(gis########)
}

// URI /{{.GisServiceCode}}/uom-gateway/monitoring-ip-addresses.json
func (t UOMGatewayMonitoringIPAddressListGet) URI() string {
	return "/{{.GisServiceCode}}/uom-gateway/monitoring-ip-addresses.json"
}

// APIName UOMGatewayMonitoringIPAddressListGet
func (t UOMGatewayMonitoringIPAddressListGet) APIName() string {
	return "UOMGatewayMonitoringIPAddressListGet"
}

// Method GET
func (t UOMGatewayMonitoringIPAddressListGet) Method() string {
	return "GET"
}

// Document http://manual.iij.jp/p2/pubapi/79743140.html
func (t UOMGatewayMonitoringIPAddressListGet) Document() string {
	return "http://manual.iij.jp/p2/pubapi/79743140.html"
}

// JPName 監視用IPアドレス情報一覧取得
func (t UOMGatewayMonitoringIPAddressListGet) JPName() string {
	return "監視用IPアドレス情報一覧取得"
}
func init() {
	APIlist = append(APIlist, UOMGatewayMonitoringIPAddressListGet{})
	TypeMap["UOMGatewayMonitoringIPAddressListGet"] = reflect.TypeOf(UOMGatewayMonitoringIPAddressListGet{})
}

// UOMGatewayMonitoringIPAddressListGetResponse 監視用IPアドレス情報一覧取得のレスポンス
type UOMGatewayMonitoringIPAddressListGetResponse struct {
	*CommonResponse
	MonitoringIpAddressList []struct { // 監視用IPアドレスの一覧
		MonitoringTargetServiceCode string `json:",omitempty"` // 監視対象となるホストのサービスコード
		PrivateNetworkServiceCode   string `json:",omitempty"` // 監視用IPアドレスが所属するプライベートネットワークのサービスコード
		HostIpAddress               string `json:",omitempty"` // 監視用IPアドレスに対応するホストIPアドレス
		MonitoringIpAddress         string `json:",omitempty"` // 割り当てられた監視用IPアドレス
	} `json:",omitempty"`
}

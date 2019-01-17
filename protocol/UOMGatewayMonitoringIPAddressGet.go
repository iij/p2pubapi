package protocol

import (
	"reflect"
)

// UOMGatewayMonitoringIPAddressGet 監視用IPアドレス情報取得
// http://manual.iij.jp/p2/pubapi/79743128.html
type UOMGatewayMonitoringIPAddressGet struct {
	GisServiceCode string `json:"-"` // P2契約のサービスコード(gis########)
	MonitoringIpAddress string `json:"-"` // 監視用IPアドレス
}

// URI /{{.GisServiceCode}}/uom-gateway/monitoring-ip-addresses/{{.MonitoringIpAddress}}.json
func (t UOMGatewayMonitoringIPAddressGet) URI() string {
	return "/{{.GisServiceCode}}/uom-gateway/monitoring-ip-addresses/{{.MonitoringIpAddress}}.json"
}

// APIName UOMGatewayMonitoringIPAddressGet
func (t UOMGatewayMonitoringIPAddressGet) APIName() string {
	return "UOMGatewayMonitoringIPAddressGet"
}

// Method GET
func (t UOMGatewayMonitoringIPAddressGet) Method() string {
	return "GET"
}

// Document http://manual.iij.jp/p2/pubapi/79743128.html
func (t UOMGatewayMonitoringIPAddressGet) Document() string {
	return "http://manual.iij.jp/p2/pubapi/79743128.html"
}

// JPName 監視用IPアドレス情報取得
func (t UOMGatewayMonitoringIPAddressGet) JPName() string {
	return "監視用IPアドレス情報取得"
}
func init() {
	APIlist = append(APIlist, UOMGatewayMonitoringIPAddressGet{})
	TypeMap["UOMGatewayMonitoringIPAddressGet"] = reflect.TypeOf(UOMGatewayMonitoringIPAddressGet{})
}

// UOMGatewayMonitoringIPAddressGetResponse 監視用IPアドレス情報取得のレスポンス
type UOMGatewayMonitoringIPAddressGetResponse struct {
	*CommonResponse
	MonitoringTargetServiceCode string `json:",omitempty"` // 監視対象となるホストのサービスコード
	PrivateNetworkServiceCode   string `json:",omitempty"` // 監視用IPアドレスが所属するプライベートネットワークのサービスコード
	HostIpAddress               string `json:",omitempty"` // 監視用IPアドレスに対応するホストIPアドレス
}

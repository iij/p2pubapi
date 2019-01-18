package protocol

import (
	"reflect"
)

// UOMGatewayMonitoringIPAddressAdd 監視用IPアドレス追加
// http://manual.iij.jp/p2/pubapi/79743131.html
type UOMGatewayMonitoringIPAddressAdd struct {
	GisServiceCode              string `json:"-"` // P2契約のサービスコード(gis########)
	MonitoringTargetServiceCode string            // 監視対象となるホストのサービスコード
	HostIpAddress               string            // 監視用IPアドレスを追加するホストIPアドレス
}

// URI /{{.GisServiceCode}}/uom-gateway/monitoring-ip-addresses.json
func (t UOMGatewayMonitoringIPAddressAdd) URI() string {
	return "/{{.GisServiceCode}}/uom-gateway/monitoring-ip-addresses.json"
}

// APIName UOMGatewayMonitoringIPAddressAdd
func (t UOMGatewayMonitoringIPAddressAdd) APIName() string {
	return "UOMGatewayMonitoringIPAddressAdd"
}

// Method POST
func (t UOMGatewayMonitoringIPAddressAdd) Method() string {
	return "POST"
}

// Document http://manual.iij.jp/p2/pubapi/79743131.html
func (t UOMGatewayMonitoringIPAddressAdd) Document() string {
	return "http://manual.iij.jp/p2/pubapi/79743131.html"
}

// JPName 監視用IPアドレス追加
func (t UOMGatewayMonitoringIPAddressAdd) JPName() string {
	return "監視用IPアドレス追加"
}
func init() {
	APIlist = append(APIlist, UOMGatewayMonitoringIPAddressAdd{})
	TypeMap["UOMGatewayMonitoringIPAddressAdd"] = reflect.TypeOf(UOMGatewayMonitoringIPAddressAdd{})
}

// UOMGatewayMonitoringIPAddressAddResponse 監視用IPアドレス追加のレスポンス
type UOMGatewayMonitoringIPAddressAddResponse struct {
	*CommonResponse
	MonitoringTargetServiceCode string `json:",omitempty"` // 監視対象となるホストのサービスコード
	PrivateNetworkServiceCode   string `json:",omitempty"` // 監視用IPアドレスが所属するプライベートネットワークのサービスコード
	HostIpAddress               string `json:",omitempty"` // ホストIPアドレス
	MonitoringIpAddress         string `json:",omitempty"` // 追加された監視用IPアドレス
}

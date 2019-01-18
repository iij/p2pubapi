package protocol

import (
	"reflect"
)

// UOMGatewayMonitoringIPAddressDelete 監視用IPアドレス削除
// http://manual.iij.jp/p2/pubapi/79743137.html
type UOMGatewayMonitoringIPAddressDelete struct {
	GisServiceCode      string `json:"-"` // P2契約のサービスコード(gis########)
	MonitoringIpAddress string `json:"-"` // 監視用IPアドレス
}

// URI /{{.GisServiceCode}}/uom-gateway/monitoring-ip-addresses/{{.MonitoringIpAddress}}.json
func (t UOMGatewayMonitoringIPAddressDelete) URI() string {
	return "/{{.GisServiceCode}}/uom-gateway/monitoring-ip-addresses/{{.MonitoringIpAddress}}.json"
}

// APIName UOMGatewayMonitoringIPAddressDelete
func (t UOMGatewayMonitoringIPAddressDelete) APIName() string {
	return "UOMGatewayMonitoringIPAddressDelete"
}

// Method DELETE
func (t UOMGatewayMonitoringIPAddressDelete) Method() string {
	return "DELETE"
}

// Document http://manual.iij.jp/p2/pubapi/79743137.html
func (t UOMGatewayMonitoringIPAddressDelete) Document() string {
	return "http://manual.iij.jp/p2/pubapi/79743137.html"
}

// JPName 監視用IPアドレス削除
func (t UOMGatewayMonitoringIPAddressDelete) JPName() string {
	return "監視用IPアドレス削除"
}
func init() {
	APIlist = append(APIlist, UOMGatewayMonitoringIPAddressDelete{})
	TypeMap["UOMGatewayMonitoringIPAddressDelete"] = reflect.TypeOf(UOMGatewayMonitoringIPAddressDelete{})
}

// UOMGatewayMonitoringIPAddressDeleteResponse 監視用IPアドレス削除のレスポンス
type UOMGatewayMonitoringIPAddressDeleteResponse struct {
	*CommonResponse
	MonitoringTargetServiceCode string `json:",omitempty"` // 監視対象となるホストのサービスコード
	PrivateNetworkServiceCode   string `json:",omitempty"` // 監視用IPアドレスが所属するプライベートネットワークのサービスコード
	HostIpAddress               string `json:",omitempty"` // 監視用IPアドレスに対応するホストIPアドレス
}

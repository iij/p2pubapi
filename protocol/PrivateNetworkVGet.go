package protocol

import (
	"reflect"
)

// PrivateNetworkVGet プライベートネットワーク/V情報取得 (同期)
//  http://manual.iij.jp/p2/pubapi/59940783.html
type PrivateNetworkVGet struct {
	GisServiceCode string `json:"-"` // P2契約のサービスコード(gis########)
	IvlServiceCode string `json:"-"` // プライベートネットワーク/Vのサービスコード(ivl########)
}

// URI /{{.GisServiceCode}}/private-networks/{{.IvlServiceCode}}.json
func (t PrivateNetworkVGet) URI() string {
	return "/{{.GisServiceCode}}/private-networks/{{.IvlServiceCode}}.json"
}

// APIName PrivateNetworkVGet
func (t PrivateNetworkVGet) APIName() string {
	return "PrivateNetworkVGet"
}

// Method GET
func (t PrivateNetworkVGet) Method() string {
	return "GET"
}

// Document http://manual.iij.jp/p2/pubapi/59940783.html
func (t PrivateNetworkVGet) Document() string {
	return "http://manual.iij.jp/p2/pubapi/59940783.html"
}

// JPName プライベートネットワーク/V情報取得
func (t PrivateNetworkVGet) JPName() string {
	return "プライベートネットワーク/V情報取得"
}
func init() {
	APIlist = append(APIlist, PrivateNetworkVGet{})
	TypeMap["PrivateNetworkVGet"] = reflect.TypeOf(PrivateNetworkVGet{})
}

// PrivateNetworkVGetResponse プライベートネットワーク/V情報取得のレスポンス
type PrivateNetworkVGetResponse struct {
	*CommonResponse
	NetworkAddress     string `json:",omitempty"` // プライベートネットワーク/Vで利用されているネットワークアドレス(IPv4アドレス/ネットマスク長)
	AttachedUomGateway struct {
		IPv4 struct {
			IpAddress string `json:",omitempty"` // 監視・運用ゲートウェイに設定されたゲートウェイアドレス(IPv4アドレス)
		} `json:",omitempty"`
		ServiceCode string `json:",omitempty"` // IIJ GIO統合運用管理サービスのサービスコード(uom########)
	} `json:",omitempty"`
	Label                    string `json:",omitempty"` // ラベル(文字列)
	ServiceCode              string `json:",omitempty"` // プライベートネットワーク/Vのサービスコード(ivl########)
	MonitorlingIpAddressList []struct {
		HostIpAddress               string `json:",omitempty"` // 監視対象のホストIPアドレス(IPv4アドレス)
		MonitoringTargetServiceCode string `json:",omitempty"` // 監視対象となるホストのサービスコード(ivm########ivd########ifl########)
		MonitoringIpAddress         string `json:",omitempty"` // 監視用IPアドレス(IPv4アドレス)
	} `json:",omitempty"`
	AttachedFwLbList []struct {
		URI              string `json:",omitempty"` // 情報へアクセスするためのURI(URI)
		NetworkInterface string `json:",omitempty"` // インターネルか、エクスターナルか(Internal""External)
		IpAddress        string `json:",omitempty"` // IPアドレス(IPｖ４アドレス)
		ServiceCode      string `json:",omitempty"` // サービスコード(ifl########)
	} `json:",omitempty"`
	StartDate                 string `json:",omitempty"` // 利用開始日(YYYYMMDD)
	ContractStatus            string `json:",omitempty"` // 契約状態
	StopDate                  string `json:",omitempty"` // 解約予定日(YYYYMMDD)
	AttachedVirtualServerList []struct {
		OSType      string `json:",omitempty"` // OS種別(Windows/Linux)
		MacAddress  string `json:",omitempty"` // 接続されたネットワークインターフェイスのMACアドレス(xx-xx-xx-xx-xx-xx)
		URI         string `json:",omitempty"` // 情報へアクセスするためのURI(URI)
		ServiceCode string `json:",omitempty"` // サービスコード(ivd########/ifl########/ivm########)
		Type        string `json:",omitempty"` // 仮想サーバ品目
	} `json:",omitempty"`
}

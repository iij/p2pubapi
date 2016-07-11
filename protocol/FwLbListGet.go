package protocol

import (
	"reflect"
)

// FwLbListGet FW+LB情報一覧取得 (同期)
//  http://manual.iij.jp/p2/pubapi/59940877.html
type FwLbListGet struct {
	GisServiceCode string `json:"-"`                // P2契約のサービスコード(gis########)
	StartIndex     string `json:"-" p2pub:",query"` // StartIndexだけスキップした位置から情報を取得する(数値)
	Count          string `json:"-" p2pub:",query"` // 取得する情報の最大数(数値)
}

// URI /{{.GisServiceCode}}/fw-lbs.json
func (t FwLbListGet) URI() string {
	return "/{{.GisServiceCode}}/fw-lbs.json"
}

// APIName FwLbListGet
func (t FwLbListGet) APIName() string {
	return "FwLbListGet"
}

// Method GET
func (t FwLbListGet) Method() string {
	return "GET"
}

// Document http://manual.iij.jp/p2/pubapi/59940877.html
func (t FwLbListGet) Document() string {
	return "http://manual.iij.jp/p2/pubapi/59940877.html"
}

// JPName FW+LB情報一覧取得
func (t FwLbListGet) JPName() string {
	return "FW+LB情報一覧取得"
}
func init() {
	APIlist = append(APIlist, FwLbListGet{})
	TypeMap["FwLbListGet"] = reflect.TypeOf(FwLbListGet{})
}

// FwLbListGetResponse FW+LB情報一覧取得のレスポンス
type FwLbListGetResponse struct {
	*CommonResponse
	FwLbList []struct {
		ResourceStatus string `json:",omitempty"` // FW+LB 専有タイプステータス
		ServiceCode    string `json:",omitempty"` // FW+LB 専有タイプのサービスコード(ifl########)
		TrafficIpList  []struct {
			IPv4 struct {
				TrafficIpName    string `json:",omitempty"` // トラフィックIP名(文字列)
				TrafficIpAddress string `json:",omitempty"` // トラフィックIPv4アドレス(IPv4アドレス)
				DomainName       string `json:",omitempty"` // 逆引きドメイン名。ネットワーク種別がGlobalに限る(文字列)
			} `json:",omitempty"`
			IPv6 struct {
				TrafficIpName    string `json:",omitempty"` // トラフィックIP名(文字列)
				TrafficIpAddress string `json:",omitempty"` // トラフィックIPv6アドレス(IPv6アドレス)
				DomainName       string `json:",omitempty"` // 逆引きドメイン名。ネットワーク種別がGlobalに限る(文字列)
			} `json:",omitempty"`
		} `json:",omitempty"`
		Internal struct {
			NetworkType      string `json:",omitempty"` // ネットワーク種別
			TrafficIpAddress string `json:",omitempty"` // トラフィックIPアドレス(IPアドレス)
			ServiceCode      string `json:",omitempty"` // ネットワーク種別がPrivateの場合のみ、サービスコード(ivl########)
		} `json:",omitempty"`
		StopDate     string `json:",omitempty"` // 解約予定日(YYYYMMDD)
		Type         string `json:",omitempty"` // FW+LB 専有タイプ品目
		SnatRuleList []struct {
			SourceIpAddress   string `json:",omitempty"` // 変換元IPアドレス(IPv4アドレス)
			ToSourceIpAddress string `json:",omitempty"` // 変換先IPアドレス(IPアドレス)
		} `json:",omitempty"`
		HostList []struct {
			ResourceStatus            string `json:",omitempty"` // FW+LB 専有タイプステータス
			LbAdministrationServerUrl string `json:",omitempty"` // LB管理サーバURL(例： https://lbcNNNNN.lb.pub.p2.iijgio.jp)
			External                  struct {
				IPv4Address string `json:",omitempty"` // ExternalインターフェイスのIPv4アドレス(IPv4アドレス)
				IPv6Address string `json:",omitempty"` // ExternalインターフェイスのIPv6アドレス(IPv6アドレス)
			} `json:",omitempty"`
			Master   string `json:",omitempty"` // マスターならばYes, スレーブならばNo(Yes" "No)
			Internal struct {
				IPv4Address string `json:",omitempty"` // InternalインターフェイスのIPv4アドレス(IPv4アドレス)
			} `json:",omitempty"`
			LbSoftwareVersion string `json:",omitempty"` // LBソフトウェアバージョン(例： 9.9)
		} `json:",omitempty"`
		External struct {
			NetworkType string `json:",omitempty"` // ネットワーク種別
			ServiceCode string `json:",omitempty"` // ネットワーク種別がPrivateの場合のみ、サービスコード(ivl########)
		} `json:",omitempty"`
		StaticRouteList []struct {
			Destination   string `json:",omitempty"` // 設定した宛先ネットワークアドレス(IPv4アドレス/マスク長)
			Gateway       string `json:",omitempty"` // 設定したゲートウェイアドレス(IPv4アドレス)
			ServiceCode   string `json:",omitempty"` // 設定したインターフェイス(ivl########)
			StaticRouteId string `json:",omitempty"` // スタティックルートID(数字)
		} `json:",omitempty"`
		Label          string `json:",omitempty"` // ラベル(文字列)
		Redundant      string `json:",omitempty"` // 冗長構成有無(Yes" "No)
		StartDate      string `json:",omitempty"` // 利用開始日(YYYYMMDD)
		ContractStatus string `json:",omitempty"` // 契約状態
	}
}

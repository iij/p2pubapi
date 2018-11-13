package protocol

import (
	"reflect"
)

// FwLbBestEffortAdd FW+LBベストエフォートタイプ追加申込 (非同期)
//  http://manual.iij.jp/p2/pubapi/161905456.html
type FwLbBestEffortAdd struct {
	GisServiceCode string `json:"-"` // P2契約のサービスコード(gis########)
	Type           string // FW+LB ベストエフォートタイプ品目
}

// URI /{{.GisServiceCode}}/best-effort-fw-lbs.json
func (t FwLbBestEffortAdd) URI() string {
	return "/{{.GisServiceCode}}/best-effort-fw-lbs.json"
}

// APIName FwLbBestEffortAdd
func (t FwLbBestEffortAdd) APIName() string {
	return "FwLbBestEffortAdd"
}

// Method POST
func (t FwLbBestEffortAdd) Method() string {
	return "POST"
}

// Document http://manual.iij.jp/p2/pubapi/161905456.html
func (t FwLbBestEffortAdd) Document() string {
	return "http://manual.iij.jp/p2/pubapi/161905456.html"
}

// JPName FW+LBベストエフォートタイプ追加申込
func (t FwLbBestEffortAdd) JPName() string {
	return "FW+LBベストエフォートタイプ追加申込"
}
func init() {
	APIlist = append(APIlist, FwLbBestEffortAdd{})
	TypeMap["FwLbBestEffortAdd"] = reflect.TypeOf(FwLbBestEffortAdd{})
}

// FwLbBestEffortAddResponse FW+LBベストエフォートタイプ追加申込のレスポンス
type FwLbBestEffortAddResponse struct {
	*CommonResponse
	ServiceCode    string // FW+LBベストエフォートタイプのサービスコード(ilb########)
	Label          string // ラベル(文字列)
	ContractStatus string // 契約状態
	StartDate      string // 利用開始日(YYYYMMDD)
	StopDate       string // 解約予定日(YYYYMMDD)
	Type           string // FW+LBベストエフォートタイプ品目
	Redundant      string // 冗長構成有無(Yes" "No)
	ResourceStatus string // FW+LBベストエフォートタイプステータス
	External       struct {
		NetworkType string `json:",omitempty"` // ネットワーク種別
	}
	Internal struct {
		NetworkType      string `json:",omitempty"` // ネットワーク種別
		TrafficIpAddress string `json:",omitempty"` // トラフィックIPアドレス(IPアドレス)
		ServiceCode      string `json:",omitempty"` // ネットワーク種別がPrivateの場合のみ、サービスコード(ivl########)
	}
	HostList []struct {
		Master         string // マスターならばYes, スレーブならばNo(Yes" "No)
		ResourceStatus string // FW+LBベストエフォートタイプステータス
		External       struct {
			IPv4Address string `json:",omitempty"` // IPv4アドレス
		}
		Internal struct {
			IPv4Address string `json:",omitempty"` // IPv4アドレス
		}
	}
	Lb struct {
		TrafficIpList []struct {
			IPv4 struct {
				TrafficIpName    string // トラフィックIP名(文字列)
				TrafficIpAddress string // トラフィックIPv4アドレス(IPv4アドレス)
				DomainName       string `json:",omitempty"` // 逆引きドメイン名。ネットワーク種別がGlobalに限る(文字列)
			} `json:",omitempty"`
		}
	}
	SnatRuleList []struct {
		SourceIpAddress   string // 変換元IPアドレス(IPv4アドレス)
		ToSourceIpAddress string // 変換先IPアドレス(IPアドレス)
	}
	StaticRouteList []struct {
		Destination   string // 設定した宛先ネットワークアドレス(IPv4アドレス/マスク長)
		Gateway       string // 設定したゲートウェイアドレス(IPv4アドレス)
		ServiceCode   string `json:",omitempty"` // 設定したインターフェイス(ivl########)
		StaticRouteId string // スタティックルートID(数字)
	}
}

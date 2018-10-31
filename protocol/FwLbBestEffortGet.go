package protocol

import (
	"reflect"
)

// FwLbBestEffortGet FW+LBベストエフォートタイプ情報取得 (同期)
//  http://manual.iij.jp/p2/pubapi/162300440.html
type FwLbBestEffortGet struct {
	GisServiceCode string `json:"-"` // P2契約のサービスコード(gis########)
	IlbServiceCode string `json:"-"` // FW+LB ベストエフォートタイプのサービスコード(ilb########)
}

// URI /{{.GisServiceCode}}/best-effort-fw-lbs/{{.IlbServiceCode}}.json
func (t FwLbBestEffortGet) URI() string {
	return "/{{.GisServiceCode}}/best-effort-fw-lbs/{{.IlbServiceCode}}.json"
}

// APIName FwLbBestEffortGet
func (t FwLbBestEffortGet) APIName() string {
	return "FwLbBestEffortGet"
}

// Method GET
func (t FwLbBestEffortGet) Method() string {
	return "GET"
}

// Document http://manual.iij.jp/p2/pubapi/162300440.html
func (t FwLbBestEffortGet) Document() string {
	return "http://manual.iij.jp/p2/pubapi/162300440.html"
}

// JPName FW+LBベストエフォートタイプ情報取得
func (t FwLbBestEffortGet) JPName() string {
	return "FW+LBベストエフォートタイプ情報取得"
}
func init() {
	APIlist = append(APIlist, FwLbBestEffortGet{})
	TypeMap["FwLbBestEffortGet"] = reflect.TypeOf(FwLbBestEffortGet{})
}

// FwLbBestEffortGetResponse FW+LBベストエフォートタイプ情報取得のレスポンス
type FwLbBestEffortGetResponse struct {
	*CommonResponse
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

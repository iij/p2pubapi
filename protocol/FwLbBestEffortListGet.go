package protocol

import (
	"reflect"
)

// FwLbBestEffortListGet FW+LBベストエフォートタイプ情報一覧取得 (同期)
//  http://manual.iij.jp/p2/pubapi/161906402.html
type FwLbBestEffortListGet struct {
	GisServiceCode string `json:"-"`                // P2契約のサービスコード(gis########)
	StartIndex     string `json:"-" p2pub:",query"` // StartIndexだけスキップした位置から情報を取得する(数値)
	Count          string `json:"-" p2pub:",query"` // 取得する情報の最大数(数値)
}

// URI /{{.GisServiceCode}}/best-effort-fw-lbs.json
func (t FwLbBestEffortListGet) URI() string {
	return "/{{.GisServiceCode}}/best-effort-fw-lbs.json"
}

// APIName FwLbBestEffortListGet
func (t FwLbBestEffortListGet) APIName() string {
	return "FwLbBestEffortListGet"
}

// Method GET
func (t FwLbBestEffortListGet) Method() string {
	return "GET"
}

// Document http://manual.iij.jp/p2/pubapi/161906402.html
func (t FwLbBestEffortListGet) Document() string {
	return "http://manual.iij.jp/p2/pubapi/161906402.html"
}

// JPName FW+LBベストエフォートタイプ情報一覧取得
func (t FwLbBestEffortListGet) JPName() string {
	return "FW+LBベストエフォートタイプ情報一覧取得"
}
func init() {
	APIlist = append(APIlist, FwLbBestEffortListGet{})
	TypeMap["FwLbBestEffortListGet"] = reflect.TypeOf(FwLbBestEffortListGet{})
}

// FwLbBestEffortListGetResponse FW+LBベストエフォートタイプ情報一覧取得のレスポンス
type FwLbBestEffortListGetResponse struct {
	*CommonResponse
	BestEffortFwLbList []struct {
		ServiceCode    string // FW+LB ベストエフォートタイプのサービスコード(ilb########)
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
}

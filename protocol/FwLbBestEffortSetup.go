package protocol

import (
	"reflect"
)

// FwLbBestEffortSetup FW+LBベストエフォートタイプセットアップ (非同期)
//  http://manual.iij.jp/p2/pubapi/162303464.html
type FwLbBestEffortSetup struct {
	GisServiceCode string `json:"-"` // P2契約のサービスコード(gis########)
	IlbServiceCode string `json:"-"` // FW+LB ベストエフォートタイプのサービスコード(ifl########)
	External       struct {
		NetworkType   string // ネットワーク種別
		TrafficIpName string // トラフィックIP名(文字列)
	} `json:",omitempty"`
	Internal struct {
		ServiceCode       string `json:",omitempty"` // プライベートネットワーク/Vのサービスコード(ivl########)
		TrafficIpAddress  string `json:",omitempty"` // トラフィックIPアドレス(IPv4アドレス)
		Netmask           string `json:",omitempty"` // ネットマスク(数字)
		NetworkType       string // ネットワーク種別
		MasterHostAddress string `json:",omitempty"` // マスターのホストアドレス(IPv4アドレス)
		TrafficIpName     string `json:",omitempty"` // トラフィックIP名(文字列)
		SlaveHostAddress  string `json:",omitempty"` // スレーブのホストアドレス(IPv4アドレス)
	} `json:",omitempty"`
	ActionType string // "Setup"(リテラル)
}

// URI /{{.GisServiceCode}}/best-effort-fw-lbs/{{.IlbServiceCode}}/action.json
func (t FwLbBestEffortSetup) URI() string {
	return "/{{.GisServiceCode}}/best-effort-fw-lbs/{{.IlbServiceCode}}/action.json"
}

// APIName FwLbBestEffortSetup
func (t FwLbBestEffortSetup) APIName() string {
	return "FwLbBestEffortSetup"
}

// Method PUT
func (t FwLbBestEffortSetup) Method() string {
	return "PUT"
}

// Document http://manual.iij.jp/p2/pubapi/162303464.html
func (t FwLbBestEffortSetup) Document() string {
	return "http://manual.iij.jp/p2/pubapi/162303464.html"
}

// JPName FW+LBベストエフォートタイプセットアップ
func (t FwLbBestEffortSetup) JPName() string {
	return "FW+LBベストエフォートタイプセットアップ"
}
func init() {
	APIlist = append(APIlist, FwLbBestEffortSetup{})
	TypeMap["FwLbBestEffortSetup"] = reflect.TypeOf(FwLbBestEffortSetup{})
}

// FwLbBestEffortSetupResponse FW+LBベストエフォートタイプセットアップのレスポンス
type FwLbBestEffortSetupResponse struct {
	*CommonResponse
	Current struct {
		ResourceStatus string `json:",omitempty"` // FW+LB ベストエフォートタイプステータス
		HostList       []struct {
			Master         string `json:",omitempty"` // マスターならばYes, スレーブならばNo(Yes" "No)
			ResourceStatus string `json:",omitempty"` // FW+LB ベストエフォートタイプステータス
		} `json:",omitempty"`
	} `json:",omitempty"`
	Previous struct {
		ResourceStatus string `json:",omitempty"` // FW+LB ベストエフォートタイプステータス
		HostList       []struct {
			Master         string `json:",omitempty"` // マスターならばYes, スレーブならばNo(Yes" "No)
			ResourceStatus string `json:",omitempty"` // FW+LB ベストエフォートタイプステータス
		} `json:",omitempty"`
	} `json:",omitempty"`
}

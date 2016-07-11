package protocol

import (
	"reflect"
)

// FwLbSetup FW+LBセットアップ (非同期)
//  http://manual.iij.jp/p2/pubapi/59940984.html
type FwLbSetup struct {
	GisServiceCode string `json:"-"` // P2契約のサービスコード(gis########)
	IflServiceCode string `json:"-"` // FW+LB 専有タイプのサービスコード(ifl########)
	External       struct {
		ServiceCode       string `json:",omitempty"` // プライベートネットワーク/Vのサービスコード(ivl########)
		TrafficIpAddress  string `json:",omitempty"` // トラフィックIPアドレス(IPv4アドレス)
		Netmask           string `json:",omitempty"` // ネットマスク(数字)
		NetworkType       string // ネットワーク種別
		MasterHostAddress string `json:",omitempty"` // マスターのホストアドレス(IPv4アドレス)
		TrafficIpName     string // トラフィックIP名(文字列)
		SlaveHostAddress  string `json:",omitempty"` // スレーブのホストアドレス(IPv4アドレス)
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

// URI /{{.GisServiceCode}}/fw-lbs/{{.IflServiceCode}}/action.json
func (t FwLbSetup) URI() string {
	return "/{{.GisServiceCode}}/fw-lbs/{{.IflServiceCode}}/action.json"
}

// APIName FwLbSetup
func (t FwLbSetup) APIName() string {
	return "FwLbSetup"
}

// Method PUT
func (t FwLbSetup) Method() string {
	return "PUT"
}

// Document http://manual.iij.jp/p2/pubapi/59940984.html
func (t FwLbSetup) Document() string {
	return "http://manual.iij.jp/p2/pubapi/59940984.html"
}

// JPName FW+LBセットアップ
func (t FwLbSetup) JPName() string {
	return "FW+LBセットアップ"
}
func init() {
	APIlist = append(APIlist, FwLbSetup{})
	TypeMap["FwLbSetup"] = reflect.TypeOf(FwLbSetup{})
}

// FwLbSetupResponse FW+LBセットアップのレスポンス
type FwLbSetupResponse struct {
	*CommonResponse
	Current struct {
		ResourceStatus string `json:",omitempty"` // FW+LB 専有タイプステータス
		HostList       []struct {
			Master         string `json:",omitempty"` // マスターならばYes, スレーブならばNo(Yes" "No)
			ResourceStatus string `json:",omitempty"` // FW+LB 専有タイプステータス
		} `json:",omitempty"`
	} `json:",omitempty"`
	Previous struct {
		ResourceStatus string `json:",omitempty"` // FW+LB 専有タイプステータス
		HostList       []struct {
			Master         string `json:",omitempty"` // マスターならばYes, スレーブならばNo(Yes" "No)
			ResourceStatus string `json:",omitempty"` // FW+LB 専有タイプステータス
		} `json:",omitempty"`
	} `json:",omitempty"`
}

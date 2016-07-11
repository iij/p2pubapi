package protocol

import (
	"reflect"
)

// GlobalAddressGetForGlobal グローバルIPアドレス情報取得（グローバルIPアドレス） (同期)
//  http://manual.iij.jp/p2/pubapi/59940463.html
type GlobalAddressGetForGlobal struct {
	GisServiceCode string `json:"-"` // P2契約のサービスコード(gis########)
}

// URI /{{.GisServiceCode}}/global-network.json
func (t GlobalAddressGetForGlobal) URI() string {
	return "/{{.GisServiceCode}}/global-network.json"
}

// APIName GlobalAddressGetForGlobal
func (t GlobalAddressGetForGlobal) APIName() string {
	return "GlobalAddressGetForGlobal"
}

// Method GET
func (t GlobalAddressGetForGlobal) Method() string {
	return "GET"
}

// Document http://manual.iij.jp/p2/pubapi/59940463.html
func (t GlobalAddressGetForGlobal) Document() string {
	return "http://manual.iij.jp/p2/pubapi/59940463.html"
}

// JPName グローバルIPアドレス情報取得（グローバルIPアドレス）
func (t GlobalAddressGetForGlobal) JPName() string {
	return "グローバルIPアドレス情報取得（グローバルIPアドレス）"
}
func init() {
	APIlist = append(APIlist, GlobalAddressGetForGlobal{})
	TypeMap["GlobalAddressGetForGlobal"] = reflect.TypeOf(GlobalAddressGetForGlobal{})
}

// GlobalAddressGetForGlobalResponse グローバルIPアドレス情報取得（グローバルIPアドレス）のレスポンス
type GlobalAddressGetForGlobalResponse struct {
	*CommonResponse
	AssignedResourceList []struct {
		IPv4 struct {
			DomainName string `json:",omitempty"` // 逆引き名(ドメイン名)
			IpAddress  string `json:",omitempty"` // 割り当てられたIPv4アドレス(IPv4アドレス)
			Type       string `json:",omitempty"` // アドレス管理(Managed/Unmanaged)
		} `json:",omitempty"`
		IPv6 struct {
			IpAddress string `json:",omitempty"` // 割り当てられたIPv6アドレス(IPv6アドレス)
			Type      string `json:",omitempty"` // アドレス管理(Managed/Unmanaged)
		} `json:",omitempty"`
		URI         string `json:",omitempty"` // 情報へアクセスするためのURI(URI)
		ServiceCode string `json:",omitempty"` // サービスコード(ifl########/ivm########)
	} `json:",omitempty"`
	GlobalIpAddressNum struct {
		Assigned string `json:",omitempty"` // 利用中のグローバルIPアドレス数(数値)
		Limit    string `json:",omitempty"` // 利用できるグローバルIPアドレス数(数値)
	} `json:",omitempty"`
	GlobalIpAddressV struct {
		StartDate      string `json:",omitempty"` // 利用開始日(YYYYMMDD)
		AddressNum     string `json:",omitempty"` // 追加で利用できるグローバルIPアドレスの数(数値)
		ContractStatus string `json:",omitempty"` // 契約状態
		ServiceCode    string `json:",omitempty"` // グローバルIPアドレス/Vのサービスコード(iga########)
		StopDate       string `json:",omitempty"` // 解約予定日(YYYYMMDD)
	} `json:",omitempty"`
}

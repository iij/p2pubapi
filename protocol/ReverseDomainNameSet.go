package protocol

import (
	"reflect"
)

// ReverseDomainNameSet 逆引きドメイン名設定 (同期)
//  https://manual.iij.jp/p2/pubapi/62130378.html
type ReverseDomainNameSet struct {
	GisServiceCode string `json:"-"` // P2契約のサービスコード(gis########)
	IPv4Address    string `json:"-"` // 割り当てられたIPv4アドレス(IPv4アドレス)
	DomainName     string // 逆引きドメイン名
}

// URI /{{.GisServiceCode}}/global-network/{{.IPv4Address}}/name.json
func (t ReverseDomainNameSet) URI() string {
	return "/{{.GisServiceCode}}/global-network/{{.IPv4Address}}/name.json"
}

// APIName ReverseDomainNameSet
func (t ReverseDomainNameSet) APIName() string {
	return "ReverseDomainNameSet"
}

// Method PUT
func (t ReverseDomainNameSet) Method() string {
	return "PUT"
}

// Document http://manual.iij.jp/p2/pubapi/59940525.html
func (t ReverseDomainNameSet) Document() string {
	return "http://manual.iij.jp/p2/pubapi/59940525.html"
}

// JPName 逆引きドメイン名設定
func (t ReverseDomainNameSet) JPName() string {
	return "逆引きドメイン名設定"
}
func init() {
	APIlist = append(APIlist, ReverseDomainNameSet{})
	TypeMap["ReverseDomainNameSet"] = reflect.TypeOf(ReverseDomainNameSet{})
}

// ReverseDomainNameSetResponse 逆引きドメイン名設定のレスポンス
type ReverseDomainNameSetResponse struct {
	*CommonResponse
	Current struct {
		DomainName string `json:",omitempty"` // 設定した逆引きドメイン名
	} `json:",omitempty"`
	Previous struct {
		DomainName string `json:",omitempty"` // 設定前の逆引きドメイン名
	} `json:",omitempty"`
}

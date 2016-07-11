package protocol

import (
	"reflect"
)

// ReverseDomainNameGet 逆引きドメイン名取得 (同期)
//  http://manual.iij.jp/p2/pubapi/59940525.html
type ReverseDomainNameGet struct {
	GisServiceCode string `json:"-"` // P2契約のサービスコード(gis########)
	IPv4Address    string `json:"-"` // 割り当てられたIPv4アドレス(IPv4アドレス)
	IPv            string `json:"-"` // IPv
}

// URI /{{.GisServiceCode}}/global-network/{{.IPv}}4Address/name.json
func (t ReverseDomainNameGet) URI() string {
	return "/{{.GisServiceCode}}/global-network/{{.IPv}}4Address/name.json"
}

// APIName ReverseDomainNameGet
func (t ReverseDomainNameGet) APIName() string {
	return "ReverseDomainNameGet"
}

// Method GET
func (t ReverseDomainNameGet) Method() string {
	return "GET"
}

// Document http://manual.iij.jp/p2/pubapi/59940525.html
func (t ReverseDomainNameGet) Document() string {
	return "http://manual.iij.jp/p2/pubapi/59940525.html"
}

// JPName 逆引きドメイン名取得
func (t ReverseDomainNameGet) JPName() string {
	return "逆引きドメイン名取得"
}
func init() {
	APIlist = append(APIlist, ReverseDomainNameGet{})
	TypeMap["ReverseDomainNameGet"] = reflect.TypeOf(ReverseDomainNameGet{})
}

// ReverseDomainNameGetResponse 逆引きドメイン名取得のレスポンス
type ReverseDomainNameGetResponse struct {
	*CommonResponse
	DomainName string `json:",omitempty"` // 逆引きドメイン名(ドメイン名)
}

package protocol

import (
	"reflect"
)

// P2PUBContractServiceCodeListGetForSA P2（パブリックリソース）契約サービスコード一覧取得（SA向け） (同期)
//  http://manual.iij.jp/p2/pubapi/58838547.html
type P2PUBContractServiceCodeListGetForSA struct {
	Item       string `json:"-" p2pub:",query"` // 取得するフィールド("ServiceCode")
	StartIndex string `json:"-" p2pub:",query"` // StartIndexだけスキップした位置から情報を取得する(数値)
	Count      string `json:"-" p2pub:",query"` // 取得する情報の最大数(数値)
}

// URI /gises.json
func (t P2PUBContractServiceCodeListGetForSA) URI() string {
	return "/gises.json"
}

// APIName P2PUBContractServiceCodeListGetForSA
func (t P2PUBContractServiceCodeListGetForSA) APIName() string {
	return "P2PUBContractServiceCodeListGetForSA"
}

// Method GET
func (t P2PUBContractServiceCodeListGetForSA) Method() string {
	return "GET"
}

// Document http://manual.iij.jp/p2/pubapi/58838547.html
func (t P2PUBContractServiceCodeListGetForSA) Document() string {
	return "http://manual.iij.jp/p2/pubapi/58838547.html"
}

// JPName P2（パブリックリソース）契約サービスコード一覧取得（SA向け）
func (t P2PUBContractServiceCodeListGetForSA) JPName() string {
	return "P2（パブリックリソース）契約サービスコード一覧取得（SA向け）"
}
func init() {
	APIlist = append(APIlist, P2PUBContractServiceCodeListGetForSA{})
	TypeMap["P2PUBContractServiceCodeListGetForSA"] = reflect.TypeOf(P2PUBContractServiceCodeListGetForSA{})
}

// P2PUBContractServiceCodeListGetForSAResponse P2（パブリックリソース）契約サービスコード一覧取得（SA向け）のレスポンス
type P2PUBContractServiceCodeListGetForSAResponse struct {
	*CommonResponse
	GisList []struct {
		ServiceCode string `json:",omitempty"` // P2契約のサービスコード(gis########)
	}
}

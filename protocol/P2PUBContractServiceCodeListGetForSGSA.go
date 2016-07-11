package protocol

import (
	"reflect"
)

// P2PUBContractServiceCodeListGetForSGSA P2（パブリックリソース）契約サービスコード一覧取得（SGSA向け） (同期)
//  http://manual.iij.jp/p2/pubapi/58838099.html
type P2PUBContractServiceCodeListGetForSGSA struct {
	Customer     string `json:"-"`                // カスタマーコード(SG########)
	CustomerCode string `json:"-"`                // CustomerCode
	Item         string `json:"-" p2pub:",query"` // 取得するフィールド("ServiceCode")
	StartIndex   string `json:"-" p2pub:",query"` // StartIndexだけスキップした位置から情報を取得する(数値)
	Count        string `json:"-" p2pub:",query"` // 取得する情報の最大数(数値)
}

// URI /{{.CustomerCode}}/gises.json
func (t P2PUBContractServiceCodeListGetForSGSA) URI() string {
	return "/{{.CustomerCode}}/gises.json"
}

// APIName P2PUBContractServiceCodeListGetForSGSA
func (t P2PUBContractServiceCodeListGetForSGSA) APIName() string {
	return "P2PUBContractServiceCodeListGetForSGSA"
}

// Method GET
func (t P2PUBContractServiceCodeListGetForSGSA) Method() string {
	return "GET"
}

// Document http://manual.iij.jp/p2/pubapi/58838099.html
func (t P2PUBContractServiceCodeListGetForSGSA) Document() string {
	return "http://manual.iij.jp/p2/pubapi/58838099.html"
}

// JPName P2（パブリックリソース）契約サービスコード一覧取得（SGSA向け）
func (t P2PUBContractServiceCodeListGetForSGSA) JPName() string {
	return "P2（パブリックリソース）契約サービスコード一覧取得（SGSA向け）"
}
func init() {
	APIlist = append(APIlist, P2PUBContractServiceCodeListGetForSGSA{})
	TypeMap["P2PUBContractServiceCodeListGetForSGSA"] = reflect.TypeOf(P2PUBContractServiceCodeListGetForSGSA{})
}

// P2PUBContractServiceCodeListGetForSGSAResponse P2（パブリックリソース）契約サービスコード一覧取得（SGSA向け）のレスポンス
type P2PUBContractServiceCodeListGetForSGSAResponse struct {
	*CommonResponse
	GisList []struct {
		ServiceCode string `json:",omitempty"` // P2契約のサービスコード(gis########)
	}
}

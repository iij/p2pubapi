package protocol

import (
	"reflect"
)

// P2PUBContractLabelSet P2（パブリックリソース）契約ラベル設定 (同期)
//  http://manual.iij.jp/p2/pubapi/58838577.html
type P2PUBContractLabelSet struct {
	GisServiceCode string `json:"-"` // P2契約のサービスコード(gis########)
	Name           string // ラベル(文字列)
}

// URI /gises/{{.GisServiceCode}}/label.json
func (t P2PUBContractLabelSet) URI() string {
	return "/gises/{{.GisServiceCode}}/label.json"
}

// APIName P2PUBContractLabelSet
func (t P2PUBContractLabelSet) APIName() string {
	return "P2PUBContractLabelSet"
}

// Method PUT
func (t P2PUBContractLabelSet) Method() string {
	return "PUT"
}

// Document http://manual.iij.jp/p2/pubapi/58838577.html
func (t P2PUBContractLabelSet) Document() string {
	return "http://manual.iij.jp/p2/pubapi/58838577.html"
}

// JPName P2（パブリックリソース）契約ラベル設定
func (t P2PUBContractLabelSet) JPName() string {
	return "P2（パブリックリソース）契約ラベル設定"
}
func init() {
	APIlist = append(APIlist, P2PUBContractLabelSet{})
	TypeMap["P2PUBContractLabelSet"] = reflect.TypeOf(P2PUBContractLabelSet{})
}

// P2PUBContractLabelSetResponse P2（パブリックリソース）契約ラベル設定のレスポンス
type P2PUBContractLabelSetResponse struct {
	*CommonResponse
	Current struct {
		Label string `json:",omitempty"` // 設定したラベル(文字列)
	} `json:",omitempty"`
	Previous struct {
		Label string `json:",omitempty"` // 設定前のラベル(文字列)
	} `json:",omitempty"`
}

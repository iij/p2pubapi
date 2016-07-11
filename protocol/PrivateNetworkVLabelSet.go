package protocol

import (
	"reflect"
)

// PrivateNetworkVLabelSet プライベートネットワーク/Vラベル設定 (同期)
//  http://manual.iij.jp/p2/pubapi/59940824.html
type PrivateNetworkVLabelSet struct {
	GisServiceCode string `json:"-"` // P2契約のサービスコード(gis########)
	IvlServiceCode string `json:"-"` // プライベートネットワーク/Vのサービスコード(ivl########)
	Name           string `json:"-"` // ラベル(文字列)
}

// URI /{{.GisServiceCode}}/private-networks/{{.IvlServiceCode}}/label.json
func (t PrivateNetworkVLabelSet) URI() string {
	return "/{{.GisServiceCode}}/private-networks/{{.IvlServiceCode}}/label.json"
}

// APIName PrivateNetworkVLabelSet
func (t PrivateNetworkVLabelSet) APIName() string {
	return "PrivateNetworkVLabelSet"
}

// Method PUT
func (t PrivateNetworkVLabelSet) Method() string {
	return "PUT"
}

// Document http://manual.iij.jp/p2/pubapi/59940824.html
func (t PrivateNetworkVLabelSet) Document() string {
	return "http://manual.iij.jp/p2/pubapi/59940824.html"
}

// JPName プライベートネットワーク/Vラベル設定
func (t PrivateNetworkVLabelSet) JPName() string {
	return "プライベートネットワーク/Vラベル設定"
}
func init() {
	APIlist = append(APIlist, PrivateNetworkVLabelSet{})
	TypeMap["PrivateNetworkVLabelSet"] = reflect.TypeOf(PrivateNetworkVLabelSet{})
}

// PrivateNetworkVLabelSetResponse プライベートネットワーク/Vラベル設定のレスポンス
type PrivateNetworkVLabelSetResponse struct {
	*CommonResponse
	Current struct {
		Label string `json:",omitempty"` // 設定したラベル(文字列)
	} `json:",omitempty"`
	Previous struct {
		Label string `json:",omitempty"` // 設定前のラベル(文字列)
	} `json:",omitempty"`
}

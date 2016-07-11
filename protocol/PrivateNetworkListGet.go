package protocol

import (
	"reflect"
)

// PrivateNetworkListGet プライベートネットワーク情報一覧取得 (同期)
//  http://manual.iij.jp/p2/pubapi/59939703.html
type PrivateNetworkListGet struct {
	GisServiceCode string `json:"-"` // P2契約のサービスコード(gis########)
	IvmServiceCode string `json:"-"` // 仮想サーバのサービスコード(ivm########, ivd########)
}

// URI /{{.GisServiceCode}}/virtual-servers/{{.IvmServiceCode}}/private-networks.json
func (t PrivateNetworkListGet) URI() string {
	return "/{{.GisServiceCode}}/virtual-servers/{{.IvmServiceCode}}/private-networks.json"
}

// APIName PrivateNetworkListGet
func (t PrivateNetworkListGet) APIName() string {
	return "PrivateNetworkListGet"
}

// Method GET
func (t PrivateNetworkListGet) Method() string {
	return "GET"
}

// Document http://manual.iij.jp/p2/pubapi/59939703.html
func (t PrivateNetworkListGet) Document() string {
	return "http://manual.iij.jp/p2/pubapi/59939703.html"
}

// JPName プライベートネットワーク情報一覧取得
func (t PrivateNetworkListGet) JPName() string {
	return "プライベートネットワーク情報一覧取得"
}
func init() {
	APIlist = append(APIlist, PrivateNetworkListGet{})
	TypeMap["PrivateNetworkListGet"] = reflect.TypeOf(PrivateNetworkListGet{})
}

// PrivateNetworkListGetResponse プライベートネットワーク情報一覧取得のレスポンス
type PrivateNetworkListGetResponse struct {
	*CommonResponse
	NetworkList []struct {
		MacAddress  string `json:",omitempty"` // MACアドレス(xx-xx-xx-xx-xx-xx)
		URI         string `json:",omitempty"` // プライベートネットワーク/Vの情報へアクセスするためのURI(URI)
		Label       string `json:",omitempty"` // ラベル(文字列)
		ServiceCode string `json:",omitempty"` // プライベートネットワーク/Vのサービスコード(ivl########)
	}
}

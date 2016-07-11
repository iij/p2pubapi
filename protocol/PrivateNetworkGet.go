package protocol

import (
	"reflect"
)

// PrivateNetworkGet プライベートネットワーク情報取得 (同期)
//  http://manual.iij.jp/p2/pubapi/59939722.html
type PrivateNetworkGet struct {
	GisServiceCode string `json:"-"` // P2契約のサービスコード(gis########)
	IvmServiceCode string `json:"-"` // 仮想サーバのサービスコード(ivm########, ivd########)
	MacAddress     string `json:"-"` // MACアドレス(xx-xx-xx-xx-xx-xx)
}

// URI /{{.GisServiceCode}}/virtual-servers/{{.IvmServiceCode}}/private-networks/{{.MacAddress}}.json
func (t PrivateNetworkGet) URI() string {
	return "/{{.GisServiceCode}}/virtual-servers/{{.IvmServiceCode}}/private-networks/{{.MacAddress}}.json"
}

// APIName PrivateNetworkGet
func (t PrivateNetworkGet) APIName() string {
	return "PrivateNetworkGet"
}

// Method GET
func (t PrivateNetworkGet) Method() string {
	return "GET"
}

// Document http://manual.iij.jp/p2/pubapi/59939722.html
func (t PrivateNetworkGet) Document() string {
	return "http://manual.iij.jp/p2/pubapi/59939722.html"
}

// JPName プライベートネットワーク情報取得
func (t PrivateNetworkGet) JPName() string {
	return "プライベートネットワーク情報取得"
}
func init() {
	APIlist = append(APIlist, PrivateNetworkGet{})
	TypeMap["PrivateNetworkGet"] = reflect.TypeOf(PrivateNetworkGet{})
}

// PrivateNetworkGetResponse プライベートネットワーク情報取得のレスポンス
type PrivateNetworkGetResponse struct {
	*CommonResponse
	NetworkList []struct {
		URI         string `json:",omitempty"` // プライベートネットワーク/Vの情報へアクセスするためのURI(URI)
		Label       string `json:",omitempty"` // ラベル(文字列)
		ServiceCode string `json:",omitempty"` // プライベートネットワーク/Vのサービスコード(ivl########)
	}
}

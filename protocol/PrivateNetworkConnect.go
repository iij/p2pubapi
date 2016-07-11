package protocol

import (
	"reflect"
)

// PrivateNetworkConnect プライベートネットワーク接続 (同期)
//  http://manual.iij.jp/p2/pubapi/59939688.html
type PrivateNetworkConnect struct {
	GisServiceCode string `json:"-"` // P2契約のサービスコード(gis########)
	IvmServiceCode string `json:"-"` // 仮想サーバのサービスコード(ivm########, ivd########)
	IvlServiceCode string // プライベートネットワーク/Vのサービスコード(ivl########)
}

// URI /{{.GisServiceCode}}/virtual-servers/{{.IvmServiceCode}}/private-networks.json
func (t PrivateNetworkConnect) URI() string {
	return "/{{.GisServiceCode}}/virtual-servers/{{.IvmServiceCode}}/private-networks.json"
}

// APIName PrivateNetworkConnect
func (t PrivateNetworkConnect) APIName() string {
	return "PrivateNetworkConnect"
}

// Method POST
func (t PrivateNetworkConnect) Method() string {
	return "POST"
}

// Document http://manual.iij.jp/p2/pubapi/59939688.html
func (t PrivateNetworkConnect) Document() string {
	return "http://manual.iij.jp/p2/pubapi/59939688.html"
}

// JPName プライベートネットワーク接続
func (t PrivateNetworkConnect) JPName() string {
	return "プライベートネットワーク接続"
}
func init() {
	APIlist = append(APIlist, PrivateNetworkConnect{})
	TypeMap["PrivateNetworkConnect"] = reflect.TypeOf(PrivateNetworkConnect{})
}

// PrivateNetworkConnectResponse プライベートネットワーク接続のレスポンス
type PrivateNetworkConnectResponse struct {
	*CommonResponse
	MacAddress  string `json:",omitempty"` // MACアドレス(xx-xx-xx-xx-xx-xx)
	URI         string `json:",omitempty"` // プライベートネットワーク/Vの情報へアクセスするためのURI(URI)
	Label       string `json:",omitempty"` // ラベル(文字列)
	ServiceCode string `json:",omitempty"` // プライベートネットワーク/Vのサービスコード(ivl########)
}

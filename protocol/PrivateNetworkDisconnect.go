package protocol

import (
	"reflect"
)

// PrivateNetworkDisconnect プライベートネットワーク切断 (同期)
//  http://manual.iij.jp/p2/pubapi/59939734.html
type PrivateNetworkDisconnect struct {
	GisServiceCode string `json:"-"` // P2契約のサービスコード(gis########)
	IvmServiceCode string `json:"-"` // 仮想サーバのサービスコード(ivm########, ivd########)
	MacAddress     string `json:"-"` // MACアドレス(xx-xx-xx-xx-xx-xx)
}

// URI /{{.GisServiceCode}}/virtual-servers/{{.IvmServiceCode}}/private-networks/{{.MacAddress}}.json
func (t PrivateNetworkDisconnect) URI() string {
	return "/{{.GisServiceCode}}/virtual-servers/{{.IvmServiceCode}}/private-networks/{{.MacAddress}}.json"
}

// APIName PrivateNetworkDisconnect
func (t PrivateNetworkDisconnect) APIName() string {
	return "PrivateNetworkDisconnect"
}

// Method DELETE
func (t PrivateNetworkDisconnect) Method() string {
	return "DELETE"
}

// Document http://manual.iij.jp/p2/pubapi/59939734.html
func (t PrivateNetworkDisconnect) Document() string {
	return "http://manual.iij.jp/p2/pubapi/59939734.html"
}

// JPName プライベートネットワーク切断
func (t PrivateNetworkDisconnect) JPName() string {
	return "プライベートネットワーク切断"
}
func init() {
	APIlist = append(APIlist, PrivateNetworkDisconnect{})
	TypeMap["PrivateNetworkDisconnect"] = reflect.TypeOf(PrivateNetworkDisconnect{})
}

// PrivateNetworkDisconnectResponse プライベートネットワーク切断のレスポンス
type PrivateNetworkDisconnectResponse struct {
	*CommonResponse
	NetworkList []struct {
		URI         string `json:",omitempty"` // プライベートネットワーク/Vの情報へアクセスするためのURI(URI)
		Label       string `json:",omitempty"` // ラベル(文字列)
		ServiceCode string `json:",omitempty"` // プライベートネットワーク/Vのサービスコード(ivl########)
	}
}

package protocol

import (
	"reflect"
)

// VMContractGet 仮想サーバ契約状態取得 (同期)
//  http://manual.iij.jp/p2/pubapi/59939439.html
type VMContractGet struct {
	GisServiceCode string `json:"-"`                // P2契約のサービスコード(gis########)
	IvmServiceCode string `json:"-"`                // 仮想サーバのサービスコード(ivm########, ivd########)
	Item           string `json:"-" p2pub:",query"` // 取得するフィールド("ContractStatus")
}

// URI /{{.GisServiceCode}}/virtual-servers/{{.IvmServiceCode}}.json
func (t VMContractGet) URI() string {
	return "/{{.GisServiceCode}}/virtual-servers/{{.IvmServiceCode}}.json"
}

// APIName VMContractGet
func (t VMContractGet) APIName() string {
	return "VMContractGet"
}

// Method GET
func (t VMContractGet) Method() string {
	return "GET"
}

// Document http://manual.iij.jp/p2/pubapi/59939439.html
func (t VMContractGet) Document() string {
	return "http://manual.iij.jp/p2/pubapi/59939439.html"
}

// JPName 仮想サーバ契約状態取得
func (t VMContractGet) JPName() string {
	return "仮想サーバ契約状態取得"
}
func init() {
	APIlist = append(APIlist, VMContractGet{})
	TypeMap["VMContractGet"] = reflect.TypeOf(VMContractGet{})
}

// VMContractGetResponse 仮想サーバ契約状態取得のレスポンス
type VMContractGetResponse struct {
	*CommonResponse
	ContractStatus string `json:",omitempty"` // 契約状態
}

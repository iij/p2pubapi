package protocol

import (
	"reflect"
)

// VMContractListGet 仮想サーバ契約状態一覧取得 (同期)
//  http://manual.iij.jp/p2/pubapi/59939414.html
type VMContractListGet struct {
	GisServiceCode string `json:"-"`                // GisServiceCode
	Item           string `json:"-" p2pub:",query"` // 取得するフィールド("ContractStatus")
	StartIndex     string `json:"-" p2pub:",query"` // StartIndexだけスキップした位置から情報を取得する(数値)
	Count          string `json:"-" p2pub:",query"` // 取得する情報の最大数(数値)
}

// URI /{{.GisServiceCode}}/virtual-servers.json
func (t VMContractListGet) URI() string {
	return "/{{.GisServiceCode}}/virtual-servers.json"
}

// APIName VMContractListGet
func (t VMContractListGet) APIName() string {
	return "VMContractListGet"
}

// Method GET
func (t VMContractListGet) Method() string {
	return "GET"
}

// Document http://manual.iij.jp/p2/pubapi/59939414.html
func (t VMContractListGet) Document() string {
	return "http://manual.iij.jp/p2/pubapi/59939414.html"
}

// JPName 仮想サーバ契約状態一覧取得
func (t VMContractListGet) JPName() string {
	return "仮想サーバ契約状態一覧取得"
}
func init() {
	APIlist = append(APIlist, VMContractListGet{})
	TypeMap["VMContractListGet"] = reflect.TypeOf(VMContractListGet{})
}

// VMContractListGetResponse 仮想サーバ契約状態一覧取得のレスポンス
type VMContractListGetResponse struct {
	*CommonResponse
	VirtualServerList []struct {
		VirtualServerList []struct {
			ContractStatus string `json:",omitempty"` // 契約状態
			ServiceCode    string `json:",omitempty"` // サービスコード(ivd########/ivm########)
		} `json:",omitempty"`
	}
}

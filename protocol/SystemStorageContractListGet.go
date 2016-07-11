package protocol

import (
	"reflect"
)

// SystemStorageContractListGet システムストレージ契約状態一覧取得 (同期)
//  http://manual.iij.jp/p2/pubapi/59939934.html
type SystemStorageContractListGet struct {
	GisServiceCode string `json:"-"`                // P2契約のサービスコード(gis########)
	Item           string `json:"-" p2pub:",query"` // 取得するフィールド("ContractStatus")
	StartIndex     string `json:"-" p2pub:",query"` // StartIndexだけスキップした位置から情報を取得する(数値)
	Count          string `json:"-" p2pub:",query"` // 取得する情報の最大数(数値)
}

// URI /{{.GisServiceCode}}/system-storages.json
func (t SystemStorageContractListGet) URI() string {
	return "/{{.GisServiceCode}}/system-storages.json"
}

// APIName SystemStorageContractListGet
func (t SystemStorageContractListGet) APIName() string {
	return "SystemStorageContractListGet"
}

// Method GET
func (t SystemStorageContractListGet) Method() string {
	return "GET"
}

// Document http://manual.iij.jp/p2/pubapi/59939934.html
func (t SystemStorageContractListGet) Document() string {
	return "http://manual.iij.jp/p2/pubapi/59939934.html"
}

// JPName システムストレージ契約状態一覧取得
func (t SystemStorageContractListGet) JPName() string {
	return "システムストレージ契約状態一覧取得"
}
func init() {
	APIlist = append(APIlist, SystemStorageContractListGet{})
	TypeMap["SystemStorageContractListGet"] = reflect.TypeOf(SystemStorageContractListGet{})
}

// SystemStorageContractListGetResponse システムストレージ契約状態一覧取得のレスポンス
type SystemStorageContractListGetResponse struct {
	*CommonResponse
	SystemStorageList []struct {
		SystemStorageList []struct {
			ContractStatus string `json:",omitempty"` // 契約状態
			ServiceCode    string `json:",omitempty"` // システムストレージのサービスコード(iba########)
		} `json:",omitempty"`
	}
}

package protocol

import (
	"reflect"
)

// SystemStorageContractGet システムストレージ契約状態取得 (同期)
//  http://manual.iij.jp/p2/pubapi/59939962.html
type SystemStorageContractGet struct {
	GisServiceCode string `json:"-"`                // P2契約のサービスコード(gis########)
	IbaServiceCode string `json:"-"`                // システムストレージのサービスコード(iba########)
	Item           string `json:"-" p2pub:",query"` // 取得するフィールド("ContractStatus")
}

// URI /{{.GisServiceCode}}/system-storages/{{.IbaServiceCode}}.json
func (t SystemStorageContractGet) URI() string {
	return "/{{.GisServiceCode}}/system-storages/{{.IbaServiceCode}}.json"
}

// APIName SystemStorageContractGet
func (t SystemStorageContractGet) APIName() string {
	return "SystemStorageContractGet"
}

// Method GET
func (t SystemStorageContractGet) Method() string {
	return "GET"
}

// Document http://manual.iij.jp/p2/pubapi/59939962.html
func (t SystemStorageContractGet) Document() string {
	return "http://manual.iij.jp/p2/pubapi/59939962.html"
}

// JPName システムストレージ契約状態取得
func (t SystemStorageContractGet) JPName() string {
	return "システムストレージ契約状態取得"
}
func init() {
	APIlist = append(APIlist, SystemStorageContractGet{})
	TypeMap["SystemStorageContractGet"] = reflect.TypeOf(SystemStorageContractGet{})
}

// SystemStorageContractGetResponse システムストレージ契約状態取得のレスポンス
type SystemStorageContractGetResponse struct {
	*CommonResponse
	ContractStatus string `json:",omitempty"` // 契約状態
}

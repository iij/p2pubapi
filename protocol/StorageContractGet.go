package protocol

import (
	"reflect"
)

// StorageContractGet 追加ストレージ契約状態取得 (同期)
//  http://manual.iij.jp/p2/pubapi/59940154.html
type StorageContractGet struct {
	GisServiceCode string `json:"-"`                // P2契約のサービスコード(gis########)
	IbgServiceCode string `json:"-"`                // 追加ストレージのサービスコード(ibb########, ibg########)
	Item           string `json:"-" p2pub:",query"` // 取得するフィールド("ContractStatus")
}

// URI /{{.GisServiceCode}}/additional-storages/{{.IbgServiceCode}}.json
func (t StorageContractGet) URI() string {
	return "/{{.GisServiceCode}}/additional-storages/{{.IbgServiceCode}}.json"
}

// APIName StorageContractGet
func (t StorageContractGet) APIName() string {
	return "StorageContractGet"
}

// Method GET
func (t StorageContractGet) Method() string {
	return "GET"
}

// Document http://manual.iij.jp/p2/pubapi/59940154.html
func (t StorageContractGet) Document() string {
	return "http://manual.iij.jp/p2/pubapi/59940154.html"
}

// JPName 追加ストレージ契約状態取得
func (t StorageContractGet) JPName() string {
	return "追加ストレージ契約状態取得"
}
func init() {
	APIlist = append(APIlist, StorageContractGet{})
	TypeMap["StorageContractGet"] = reflect.TypeOf(StorageContractGet{})
}

// StorageContractGetResponse 追加ストレージ契約状態取得のレスポンス
type StorageContractGetResponse struct {
	*CommonResponse
	ContractStatus string `json:",omitempty"` // 契約状態
}

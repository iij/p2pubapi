package protocol

import (
	"reflect"
)

// StorageContractListGet 追加ストレージ契約状態一覧取得 (同期)
//  http://manual.iij.jp/p2/pubapi/59940130.html
type StorageContractListGet struct {
	GisServiceCode string `json:"-"`                // P2契約のサービスコード(gis########)
	Item           string `json:"-" p2pub:",query"` // 取得するフィールド("ContractStatus")
	StartIndex     string `json:"-" p2pub:",query"` // StartIndexだけスキップした位置から情報を取得する(数値)
	Count          string `json:"-" p2pub:",query"` // 取得する情報の最大数(数値)
}

// URI /{{.GisServiceCode}}/additional-storages.json
func (t StorageContractListGet) URI() string {
	return "/{{.GisServiceCode}}/additional-storages.json"
}

// APIName StorageContractListGet
func (t StorageContractListGet) APIName() string {
	return "StorageContractListGet"
}

// Method GET
func (t StorageContractListGet) Method() string {
	return "GET"
}

// Document http://manual.iij.jp/p2/pubapi/59940130.html
func (t StorageContractListGet) Document() string {
	return "http://manual.iij.jp/p2/pubapi/59940130.html"
}

// JPName 追加ストレージ契約状態一覧取得
func (t StorageContractListGet) JPName() string {
	return "追加ストレージ契約状態一覧取得"
}
func init() {
	APIlist = append(APIlist, StorageContractListGet{})
	TypeMap["StorageContractListGet"] = reflect.TypeOf(StorageContractListGet{})
}

// StorageContractListGetResponse 追加ストレージ契約状態一覧取得のレスポンス
type StorageContractListGetResponse struct {
	*CommonResponse
	AdditionalStorageList []struct {
		AdditionalStorageList []struct {
			ContractStatus string `json:",omitempty"` // 契約状態
			ServiceCode    string `json:",omitempty"` // 追加ストレージのサービスコード(ibg########/ibb########/icg########/icb########)
		} `json:",omitempty"`
	}
}

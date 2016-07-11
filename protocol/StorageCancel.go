package protocol

import (
	"reflect"
)

// StorageCancel 追加ストレージ解約申込 (同期)
//  http://manual.iij.jp/p2/pubapi/59940162.html
type StorageCancel struct {
	GisServiceCode string `json:"-"`                // P2契約のサービスコード(gis########)
	IbgServiceCode string `json:"-"`                // 追加ストレージのサービスコード(ibb########, ibg########)
	StopDate       string `json:"-" p2pub:",query"` // 解約予定日。省略した場合は即時(YYYYMMDD)
}

// URI /{{.GisServiceCode}}/additional-storages/{{.IbgServiceCode}}.json
func (t StorageCancel) URI() string {
	return "/{{.GisServiceCode}}/additional-storages/{{.IbgServiceCode}}.json"
}

// APIName StorageCancel
func (t StorageCancel) APIName() string {
	return "StorageCancel"
}

// Method DELETE
func (t StorageCancel) Method() string {
	return "DELETE"
}

// Document http://manual.iij.jp/p2/pubapi/59940162.html
func (t StorageCancel) Document() string {
	return "http://manual.iij.jp/p2/pubapi/59940162.html"
}

// JPName 追加ストレージ解約申込
func (t StorageCancel) JPName() string {
	return "追加ストレージ解約申込"
}
func init() {
	APIlist = append(APIlist, StorageCancel{})
	TypeMap["StorageCancel"] = reflect.TypeOf(StorageCancel{})
}

// StorageCancelResponse 追加ストレージ解約申込のレスポンス
type StorageCancelResponse struct {
	*CommonResponse
	Current struct {
		StopDate string `json:",omitempty"` // 設定した解約予定日(YYYYMMDD)
	} `json:",omitempty"`
	Previous struct {
		StopDate string `json:",omitempty"` // 設定前の解約予定日(YYYYMMDD。未設定ならば空文字列)
	} `json:",omitempty"`
}

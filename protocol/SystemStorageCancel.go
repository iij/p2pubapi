package protocol

import (
	"reflect"
)

// SystemStorageCancel システムストレージ解約申込 (同期)
//  http://manual.iij.jp/p2/pubapi/59939969.html
type SystemStorageCancel struct {
	GisServiceCode     string `json:"-"`                // P2契約のサービスコード(gis########)
	StorageServiceCode string `json:"-"`                // システムストレージのサービスコード(iba########/ica########)
	StopDate           string `json:"-" p2pub:",query"` // 解約予定日(YYYYMMDD)
}

// URI /{{.GisServiceCode}}/system-storages/{{.StorageServiceCode}}.json
func (t SystemStorageCancel) URI() string {
	return "/{{.GisServiceCode}}/system-storages/{{.StorageServiceCode}}.json"
}

// APIName SystemStorageCancel
func (t SystemStorageCancel) APIName() string {
	return "SystemStorageCancel"
}

// Method DELETE
func (t SystemStorageCancel) Method() string {
	return "DELETE"
}

// Document http://manual.iij.jp/p2/pubapi/59939969.html
func (t SystemStorageCancel) Document() string {
	return "http://manual.iij.jp/p2/pubapi/59939969.html"
}

// JPName システムストレージ解約申込
func (t SystemStorageCancel) JPName() string {
	return "システムストレージ解約申込"
}
func init() {
	APIlist = append(APIlist, SystemStorageCancel{})
	TypeMap["SystemStorageCancel"] = reflect.TypeOf(SystemStorageCancel{})
}

// SystemStorageCancelResponse システムストレージ解約申込のレスポンス
type SystemStorageCancelResponse struct {
	*CommonResponse
	Current struct {
		StopDate string `json:",omitempty"` // 設定した解約予定日(YYYYMMDD)
	} `json:",omitempty"`
	Previous struct {
		StopDate string `json:",omitempty"` // 設定前の解約予定日(YYYYMMDD)
	} `json:",omitempty"`
}

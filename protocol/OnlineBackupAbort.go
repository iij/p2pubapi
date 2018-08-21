package protocol

import (
	"reflect"
)

// OnlineBackupAbort オンラインバックアップ中止 (非同期)
//  http://manual.iij.jp/p2/pubapi/152684075.html
type OnlineBackupAbort struct {
	GisServiceCode     string `json:"-"`          // P2契約のサービスコード(gis########)
	StorageServiceCode string `json:"-"`          // ストレージのサービスコード(ica########)
}

// URI /{{.GisServiceCode}}/system-storages/{{.StorageServiceCode}}/backup.json
func (t OnlineBackupAbort) URI() string {
	return "/{{.GisServiceCode}}/system-storages/{{.StorageServiceCode}}/backup.json"
}

// APIName OnlineBackupAbort
func (t OnlineBackupAbort) APIName() string {
	return "OnlineBackupAbort"
}

// Method DELETE
func (t OnlineBackupAbort) Method() string {
	return "DELETE"
}

// Document http://manual.iij.jp/p2/pubapi/152684075.html
func (t OnlineBackupAbort) Document() string {
	return "http://manual.iij.jp/p2/pubapi/152684075.html"
}

// JPName オンラインバックアップ中止
func (t OnlineBackupAbort) JPName() string {
	return "オンラインバックアップ中止"
}
func init() {
	APIlist = append(APIlist, OnlineBackupAbort{})
	TypeMap["OnlineBackupAbort"] = reflect.TypeOf(OnlineBackupAbort{})
}

// OnlineBackupAbortResponse オンラインバックアップ中止のレスポンス
type OnlineBackupAbortResponse struct {
	*CommonResponse
	Current struct {
		ResourceStatus string `json:",omitempty"`
	} `json:",omitempty"`
	Previous struct {
		ResourceStatus string `json:",omitempty"`
	} `json:",omitempty"`
}

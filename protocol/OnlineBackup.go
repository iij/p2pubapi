package protocol

import (
	"reflect"
)

// OnlineBackup オンラインバックアップ (非同期)
//  http://manual.iij.jp/p2/pubapi/152684064.html
type OnlineBackup struct {
	GisServiceCode     string `json:"-"`          // P2契約のサービスコード(gis########)
	StorageServiceCode string `json:"-"`          // ストレージのサービスコード(ica########)
	Label              string `json:",omitempty"` // ラベル
}

// URI /{{.GisServiceCode}}/system-storages/{{.StorageServiceCode}}/backup.json
func (t OnlineBackup) URI() string {
	return "/{{.GisServiceCode}}/system-storages/{{.StorageServiceCode}}/backup.json"
}

// APIName OnlineBackup
func (t OnlineBackup) APIName() string {
	return "OnlineBackup"
}

// Method PUT
func (t OnlineBackup) Method() string {
	return "PUT"
}

// Document http://manual.iij.jp/p2/pubapi/152684064.html
func (t OnlineBackup) Document() string {
	return "http://manual.iij.jp/p2/pubapi/152684064.html"
}

// JPName オンラインバックアップ
func (t OnlineBackup) JPName() string {
	return "オンラインバックアップ"
}
func init() {
	APIlist = append(APIlist, OnlineBackup{})
	TypeMap["OnlineBackup"] = reflect.TypeOf(OnlineBackup{})
}

// OnlineBackupResponse オンラインバックアップのレスポンス
type OnlineBackupResponse struct {
	*CommonResponse
	Current struct {
		ResourceStatus string `json:",omitempty"`
	} `json:",omitempty"`
	Previous struct {
		ResourceStatus string `json:",omitempty"`
	} `json:",omitempty"`
}

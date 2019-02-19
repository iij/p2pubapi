package protocol

import (
	"reflect"
)

// PasswordSet パスワード設定 (非同期)
//  http://manual.iij.jp/p2/pubapi/59940023.html
type PasswordSet struct {
	GisServiceCode     string `json:"-"` // P2契約のサービスコード(gis########)
	StorageServiceCode string `json:"-"` // システムストレージのサービスコード(iba########/ica########)
	Password           string // パスワード(8～32文字)
}

// URI /{{.GisServiceCode}}/system-storages/{{.StorageServiceCode}}/password.json
func (t PasswordSet) URI() string {
	return "/{{.GisServiceCode}}/system-storages/{{.StorageServiceCode}}/password.json"
}

// APIName PasswordSet
func (t PasswordSet) APIName() string {
	return "PasswordSet"
}

// Method PUT
func (t PasswordSet) Method() string {
	return "PUT"
}

// Document http://manual.iij.jp/p2/pubapi/59940023.html
func (t PasswordSet) Document() string {
	return "http://manual.iij.jp/p2/pubapi/59940023.html"
}

// JPName パスワード設定
func (t PasswordSet) JPName() string {
	return "パスワード設定"
}
func init() {
	APIlist = append(APIlist, PasswordSet{})
	TypeMap["PasswordSet"] = reflect.TypeOf(PasswordSet{})
}

// PasswordSetResponse パスワード設定のレスポンス
type PasswordSetResponse struct {
	*CommonResponse
	Current struct {
		ResourceStatus string `json:",omitempty"` // 設定後のストレージステータス
	} `json:",omitempty"`
	Previous struct {
		ResourceStatus string `json:",omitempty"` // 設定前のストレージステータス
	} `json:",omitempty"`
}

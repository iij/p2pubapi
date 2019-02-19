package protocol

import (
	"reflect"
)

// PublicKeyAdd 公開鍵追加 (非同期)
//  http://manual.iij.jp/p2/pubapi/59940004.html
type PublicKeyAdd struct {
	GisServiceCode     string `json:"-"` // P2契約のサービスコード(gis########)
	StorageServiceCode string `json:"-"` // システムストレージのサービスコード(iba########/ica########)
	PublicKey          string // 公開鍵(文字列。例： "ssh-rsa XXXXXXXXXXXXXXXXXXXX comment)
}

// URI /{{.GisServiceCode}}/system-storages/{{.StorageServiceCode}}/publickeys.json
func (t PublicKeyAdd) URI() string {
	return "/{{.GisServiceCode}}/system-storages/{{.StorageServiceCode}}/publickeys.json"
}

// APIName PublicKeyAdd
func (t PublicKeyAdd) APIName() string {
	return "PublicKeyAdd"
}

// Method POST
func (t PublicKeyAdd) Method() string {
	return "POST"
}

// Document http://manual.iij.jp/p2/pubapi/59940004.html
func (t PublicKeyAdd) Document() string {
	return "http://manual.iij.jp/p2/pubapi/59940004.html"
}

// JPName 公開鍵追加
func (t PublicKeyAdd) JPName() string {
	return "公開鍵追加"
}
func init() {
	APIlist = append(APIlist, PublicKeyAdd{})
	TypeMap["PublicKeyAdd"] = reflect.TypeOf(PublicKeyAdd{})
}

// PublicKeyAddResponse 公開鍵追加のレスポンス
type PublicKeyAddResponse struct {
	*CommonResponse
	Current struct {
		ResourceStatus string `json:",omitempty"` // 設定後のストレージステータス
	} `json:",omitempty"`
	Previous struct {
		ResourceStatus string `json:",omitempty"` // 設定前のストレージステータス
	} `json:",omitempty"`
}

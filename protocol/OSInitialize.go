package protocol

import (
	"reflect"
)

// OSInitialize OS初期化 (非同期)
//  http://manual.iij.jp/p2/pubapi/59940041.html
type OSInitialize struct {
	GisServiceCode     string `json:"-"`          // P2契約のサービスコード(gis########)
	StorageServiceCode string `json:"-"`          // システムストレージのサービスコード(iba########, ica########)
	Version            string `json:",omitempty"` // OSのバージョン。省略時は最新バージョンが選択される(文字列。指定可能な値はコントロールパネルでOS初期化を実行するときに、ドロップダウンボックスに表示されるものと同じ)
	Image              string // 操作内容(Initialize)
}

// URI /{{.GisServiceCode}}/system-storages/{{.StorageServiceCode}}/action.json
func (t OSInitialize) URI() string {
	return "/{{.GisServiceCode}}/system-storages/{{.StorageServiceCode}}/action.json"
}

// APIName OSInitialize
func (t OSInitialize) APIName() string {
	return "OSInitialize"
}

// Method PUT
func (t OSInitialize) Method() string {
	return "PUT"
}

// Document http://manual.iij.jp/p2/pubapi/59940041.html
func (t OSInitialize) Document() string {
	return "http://manual.iij.jp/p2/pubapi/59940041.html"
}

// JPName OS初期化
func (t OSInitialize) JPName() string {
	return "OS初期化"
}
func init() {
	APIlist = append(APIlist, OSInitialize{})
	TypeMap["OSInitialize"] = reflect.TypeOf(OSInitialize{})
}

// OSInitializeResponse OS初期化のレスポンス
type OSInitializeResponse struct {
	*CommonResponse
	Current struct {
		ResourceStatus string `json:",omitempty"` // 設定後のストレージステータス
	} `json:",omitempty"`
	Previous struct {
		ResourceStatus string `json:",omitempty"` // 設定前のストレージステータス
	} `json:",omitempty"`
}

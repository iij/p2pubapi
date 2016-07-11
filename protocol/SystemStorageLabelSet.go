package protocol

import (
	"reflect"
)

// SystemStorageLabelSet システムストレージラベル設定 (同期)
//  http://manual.iij.jp/p2/pubapi/59939994.html
type SystemStorageLabelSet struct {
	GisServiceCode string `json:"-"` // P2契約のサービスコード(gis########)
	IbaServiceCode string `json:"-"` // システムストレージのサービスコード(iba########)
	Name           string // ラベル(文字列)
}

// URI /{{.GisServiceCode}}/system-storages/{{.IbaServiceCode}}/label.json
func (t SystemStorageLabelSet) URI() string {
	return "/{{.GisServiceCode}}/system-storages/{{.IbaServiceCode}}/label.json"
}

// APIName SystemStorageLabelSet
func (t SystemStorageLabelSet) APIName() string {
	return "SystemStorageLabelSet"
}

// Method PUT
func (t SystemStorageLabelSet) Method() string {
	return "PUT"
}

// Document http://manual.iij.jp/p2/pubapi/59939994.html
func (t SystemStorageLabelSet) Document() string {
	return "http://manual.iij.jp/p2/pubapi/59939994.html"
}

// JPName システムストレージラベル設定
func (t SystemStorageLabelSet) JPName() string {
	return "システムストレージラベル設定"
}
func init() {
	APIlist = append(APIlist, SystemStorageLabelSet{})
	TypeMap["SystemStorageLabelSet"] = reflect.TypeOf(SystemStorageLabelSet{})
}

// SystemStorageLabelSetResponse システムストレージラベル設定のレスポンス
type SystemStorageLabelSetResponse struct {
	*CommonResponse
	Current struct {
		Label string `json:",omitempty"` // 設定したラベル(文字列)
	} `json:",omitempty"`
	Previous struct {
		Label string `json:",omitempty"` // 設定前のラベル(文字列)
	} `json:",omitempty"`
}

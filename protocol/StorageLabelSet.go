package protocol

import (
	"reflect"
)

// StorageLabelSet 追加ストレージラベル設定 (同期)
//  http://manual.iij.jp/p2/pubapi/59940185.html
type StorageLabelSet struct {
	GisServiceCode string `json:"-"` // P2契約のサービスコード(gis########)
	IbgServiceCode string `json:"-"` // 追加ストレージのサービスコード(ibb########, ibg########)
	Name           string // ラベル(文字列)
}

// URI /{{.GisServiceCode}}/additional-storages/{{.IbgServiceCode}}/label.json
func (t StorageLabelSet) URI() string {
	return "/{{.GisServiceCode}}/additional-storages/{{.IbgServiceCode}}/label.json"
}

// APIName StorageLabelSet
func (t StorageLabelSet) APIName() string {
	return "StorageLabelSet"
}

// Method PUT
func (t StorageLabelSet) Method() string {
	return "PUT"
}

// Document http://manual.iij.jp/p2/pubapi/59940185.html
func (t StorageLabelSet) Document() string {
	return "http://manual.iij.jp/p2/pubapi/59940185.html"
}

// JPName 追加ストレージラベル設定
func (t StorageLabelSet) JPName() string {
	return "追加ストレージラベル設定"
}
func init() {
	APIlist = append(APIlist, StorageLabelSet{})
	TypeMap["StorageLabelSet"] = reflect.TypeOf(StorageLabelSet{})
}

// StorageLabelSetResponse 追加ストレージラベル設定のレスポンス
type StorageLabelSetResponse struct {
	*CommonResponse
	Current struct {
		Label string `json:",omitempty"` // 設定したラベル(文字列)
	} `json:",omitempty"`
	Previous struct {
		Label string `json:",omitempty"` // 設定前のラベル(文字列)
	} `json:",omitempty"`
}

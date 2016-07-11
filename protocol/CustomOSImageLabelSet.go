package protocol

import (
	"reflect"
)

// CustomOSImageLabelSet カスタムOSイメージラベル設定 (同期)
//  http://manual.iij.jp/p2/pubapi/59940388.html
type CustomOSImageLabelSet struct {
	GisServiceCode string `json:"-"` // P2契約のサービスコード(gis########)
	IarServiceCode string `json:"-"` // ストレージアーカイブのサービスコード(iar########)
	ImageId        string `json:"-"` // カスタムOSイメージのId(数値)
	Name           string // ラベル(文字列)
}

// URI /{{.GisServiceCode}}/storage-archive/{{.IarServiceCode}}/images/{{.ImageId}}/name.json
func (t CustomOSImageLabelSet) URI() string {
	return "/{{.GisServiceCode}}/storage-archive/{{.IarServiceCode}}/images/{{.ImageId}}/name.json"
}

// APIName CustomOSImageLabelSet
func (t CustomOSImageLabelSet) APIName() string {
	return "CustomOSImageLabelSet"
}

// Method PUT
func (t CustomOSImageLabelSet) Method() string {
	return "PUT"
}

// Document http://manual.iij.jp/p2/pubapi/59940388.html
func (t CustomOSImageLabelSet) Document() string {
	return "http://manual.iij.jp/p2/pubapi/59940388.html"
}

// JPName カスタムOSイメージラベル設定
func (t CustomOSImageLabelSet) JPName() string {
	return "カスタムOSイメージラベル設定"
}
func init() {
	APIlist = append(APIlist, CustomOSImageLabelSet{})
	TypeMap["CustomOSImageLabelSet"] = reflect.TypeOf(CustomOSImageLabelSet{})
}

// CustomOSImageLabelSetResponse カスタムOSイメージラベル設定のレスポンス
type CustomOSImageLabelSetResponse struct {
	*CommonResponse
	Current struct {
		Label string `json:",omitempty"` // 設定したラベル(文字列)
	} `json:",omitempty"`
	Previous struct {
		Label string `json:",omitempty"` // 設定前のラベル(文字列)
	} `json:",omitempty"`
}

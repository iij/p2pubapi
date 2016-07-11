package protocol

import (
	"reflect"
)

// CustomOSImageCreate カスタムOSイメージ作成 (非同期)
//  http://manual.iij.jp/p2/pubapi/59940344.html
type CustomOSImageCreate struct {
	GisServiceCode string `json:"-"` // P2契約のサービスコード(gis########)
	IarServiceCode string `json:"-"` // ストレージアーカイブのサービスコード(iar########)
	Name           string // カスタムOSイメージの名前(文字列)
	IbaServiceCode string // システムストレージのサービスコード(iba########)
}

// URI /{{.GisServiceCode}}/storage-archive/{{.IarServiceCode}}/images.json
func (t CustomOSImageCreate) URI() string {
	return "/{{.GisServiceCode}}/storage-archive/{{.IarServiceCode}}/images.json"
}

// APIName CustomOSImageCreate
func (t CustomOSImageCreate) APIName() string {
	return "CustomOSImageCreate"
}

// Method POST
func (t CustomOSImageCreate) Method() string {
	return "POST"
}

// Document http://manual.iij.jp/p2/pubapi/59940344.html
func (t CustomOSImageCreate) Document() string {
	return "http://manual.iij.jp/p2/pubapi/59940344.html"
}

// JPName カスタムOSイメージ作成
func (t CustomOSImageCreate) JPName() string {
	return "カスタムOSイメージ作成"
}
func init() {
	APIlist = append(APIlist, CustomOSImageCreate{})
	TypeMap["CustomOSImageCreate"] = reflect.TypeOf(CustomOSImageCreate{})
}

// CustomOSImageCreateResponse カスタムOSイメージ作成のレスポンス
type CustomOSImageCreateResponse struct {
	*CommonResponse
	Status           string `json:",omitempty"` // ステータス(Creating/Created)
	OSType           string `json:",omitempty"` // OS種別(Windows/Linux)
	ArchivedDateTime string `json:",omitempty"` // カスタムOSイメージの作成日時(YYYY-MM-DDThh:mm:ssZ)
	Label            string `json:",omitempty"` // ラベル(文字列)
	SrcServiceCode   string `json:",omitempty"` // システムストレージのサービスコード(iba########)
	ImageId          string `json:",omitempty"` // カスタムOSイメージのId(数値)
	ImageSize        string `json:",omitempty"` // カスタムOSイメージのサイズ（MB）(文字列)
	Type             string `json:",omitempty"` // システムストレージ品目
}

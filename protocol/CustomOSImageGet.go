package protocol

import (
	"reflect"
)

// CustomOSImageGet カスタムOSイメージ情報取得 (同期)
//  http://manual.iij.jp/p2/pubapi/59940366.html
type CustomOSImageGet struct {
	GisServiceCode string `json:"-"` // P2契約のサービスコード(gis########)
	IarServiceCode string `json:"-"` // ストレージアーカイブのサービスコード(iar########)
	ImageId        string `json:"-"` // カスタムOSイメージのId(数値)
}

// URI /{{.GisServiceCode}}/storage-archive/{{.IarServiceCode}}/images/{{.ImageId}}.json
func (t CustomOSImageGet) URI() string {
	return "/{{.GisServiceCode}}/storage-archive/{{.IarServiceCode}}/images/{{.ImageId}}.json"
}

// APIName CustomOSImageGet
func (t CustomOSImageGet) APIName() string {
	return "CustomOSImageGet"
}

// Method GET
func (t CustomOSImageGet) Method() string {
	return "GET"
}

// Document http://manual.iij.jp/p2/pubapi/59940366.html
func (t CustomOSImageGet) Document() string {
	return "http://manual.iij.jp/p2/pubapi/59940366.html"
}

// JPName カスタムOSイメージ情報取得
func (t CustomOSImageGet) JPName() string {
	return "カスタムOSイメージ情報取得"
}
func init() {
	APIlist = append(APIlist, CustomOSImageGet{})
	TypeMap["CustomOSImageGet"] = reflect.TypeOf(CustomOSImageGet{})
}

// CustomOSImageGetResponse カスタムOSイメージ情報取得のレスポンス
type CustomOSImageGetResponse struct {
	*CommonResponse
	Status           string `json:",omitempty"` // ステータス(Creating、Created)
	OSType           string `json:",omitempty"` // OS種別(Windows/Linux)
	ArchivedDateTime string `json:",omitempty"` // カスタムOSイメージの作成日時(YYYY-MM-DDThh:mm:ssZ)
	Label            string `json:",omitempty"` // ラベル(文字列)
	SrcServiceCode   string `json:",omitempty"` // システムストレージのサービスコード(iba########)
	ImageId          string `json:",omitempty"` // カスタムOSイメージのId(数値)
	ImageSize        string `json:",omitempty"` // カスタムOSイメージのサイズ（MB）(文字列)
	Type             string `json:",omitempty"` // システムストレージ品目
}

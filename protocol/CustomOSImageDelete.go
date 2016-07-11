package protocol

import (
	"reflect"
)

// CustomOSImageDelete カスタムOSイメージ削除 (同期)
//  http://manual.iij.jp/p2/pubapi/59940375.html
type CustomOSImageDelete struct {
	GisServiceCode string `json:"-"` // P2契約のサービスコード(gis########)
	IarServiceCode string `json:"-"` // ストレージアーカイブのサービスコード(iar########)
	ImageId        string `json:"-"` // カスタムOSイメージのId(数値)
}

// URI /{{.GisServiceCode}}/storage-archive/{{.IarServiceCode}}/images/{{.ImageId}}.json
func (t CustomOSImageDelete) URI() string {
	return "/{{.GisServiceCode}}/storage-archive/{{.IarServiceCode}}/images/{{.ImageId}}.json"
}

// APIName CustomOSImageDelete
func (t CustomOSImageDelete) APIName() string {
	return "CustomOSImageDelete"
}

// Method DELETE
func (t CustomOSImageDelete) Method() string {
	return "DELETE"
}

// Document http://manual.iij.jp/p2/pubapi/59940375.html
func (t CustomOSImageDelete) Document() string {
	return "http://manual.iij.jp/p2/pubapi/59940375.html"
}

// JPName カスタムOSイメージ削除
func (t CustomOSImageDelete) JPName() string {
	return "カスタムOSイメージ削除"
}
func init() {
	APIlist = append(APIlist, CustomOSImageDelete{})
	TypeMap["CustomOSImageDelete"] = reflect.TypeOf(CustomOSImageDelete{})
}

// CustomOSImageDeleteResponse カスタムOSイメージ削除のレスポンス
type CustomOSImageDeleteResponse struct {
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

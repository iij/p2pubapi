package protocol

import (
	"reflect"
)

// CustomOSImageListGet カスタムOSイメージ情報一覧取得 (同期)
//  http://manual.iij.jp/p2/pubapi/59940352.html
type CustomOSImageListGet struct {
	GisServiceCode string `json:"-"` // P2契約のサービスコード(gis########)
	IarServiceCode string `json:"-"` // ストレージアーカイブのサービスコード(iar########)
}

// URI /{{.GisServiceCode}}/storage-archive/{{.IarServiceCode}}/images.json
func (t CustomOSImageListGet) URI() string {
	return "/{{.GisServiceCode}}/storage-archive/{{.IarServiceCode}}/images.json"
}

// APIName CustomOSImageListGet
func (t CustomOSImageListGet) APIName() string {
	return "CustomOSImageListGet"
}

// Method GET
func (t CustomOSImageListGet) Method() string {
	return "GET"
}

// Document http://manual.iij.jp/p2/pubapi/59940352.html
func (t CustomOSImageListGet) Document() string {
	return "http://manual.iij.jp/p2/pubapi/59940352.html"
}

// JPName カスタムOSイメージ情報一覧取得
func (t CustomOSImageListGet) JPName() string {
	return "カスタムOSイメージ情報一覧取得"
}
func init() {
	APIlist = append(APIlist, CustomOSImageListGet{})
	TypeMap["CustomOSImageListGet"] = reflect.TypeOf(CustomOSImageListGet{})
}

// CustomOSImageListGetResponse カスタムOSイメージ情報一覧取得のレスポンス
type CustomOSImageListGetResponse struct {
	*CommonResponse
	ImageList []struct {
		Status           string `json:",omitempty"` // ステータス(Creating/Created)
		OSType           string `json:",omitempty"` // OS種別(Windows/Linux)
		ArchivedDateTime string `json:",omitempty"` // カスタムOSイメージの作成日時(YYYY-MM-DDThh:mm:ssZ)
		Label            string `json:",omitempty"` // ラベル(文字列)
		SrcServiceCode   string `json:",omitempty"` // システムストレージのサービスコード(iba########/ica########)
		ImageId          string `json:",omitempty"` // カスタムOSイメージのId(数値)
		ImageSize        string `json:",omitempty"` // カスタムOSイメージのサイズ（MB）(文字列。作成中は"0MB)
		Type             string `json:",omitempty"` // システムストレージ品目
	}
}

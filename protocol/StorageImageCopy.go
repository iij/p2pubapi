package protocol

import (
	"reflect"
)

// StorageImageCopy ストレージイメージコピー
//  http://manual.iij.jp/p2/pubapi/137048215.html
type StorageImageCopy struct {
	SrcGisServiceCode string `json:"-"` // 操作対象が所属するgisのサービスコード(gis########)
	SrcIarServiceCode string `json:"-"` // 操作対象となるイメージが所属するストレージアーカイブのサービスコード(iar########)
	SrcImageId        string `json:"-"` // 操作対象となるイメージのID
	DstGisServiceCode string `json:"GisServiceCode"` // コピー先のgisのサービスコード(gis########)
	DstIarServiceCode string `json:"IarServiceCode"` // コピー先のiarのサービスコード(iar########)
	Image             string  // コピー ("Copy")
}

// URI /{{.GisServiceCode}}/storage-archive/{{.IarServiceCode}}/images/{{.ImageId}}/action.json
func (t StorageImageCopy) URI() string {
	return "/{{.SrcGisServiceCode}}/storage-archive/{{.SrcIarServiceCode}}/images/{{.SrcImageId}}/action.json"
}

// APIName StorageImageCopy
func (t StorageImageCopy) APIName() string {
	return "StorageImageCopy"
}

// Method POST
func (t StorageImageCopy) Method() string {
	return "POST"
}

// Document http://manual.iij.jp/p2/pubapi/137048215.html
func (t StorageImageCopy) Document() string {
	return "http://manual.iij.jp/p2/pubapi/137048215.html"
}

// JPName ストレージイメージコピー
func (t StorageImageCopy) JPName() string {
	return "ストレージイメージコピー"
}
func init() {
	APIlist = append(APIlist, StorageImageCopy{})
	TypeMap["StorageImageCopy"] = reflect.TypeOf(StorageImageCopy{})
}

// StorageImageCopyResponse
type StorageImageCopyResponse struct {
	*CommonResponse
	GisServiceCode string `json:",omitempty"` // コピー先のgisのサービスコード
	IarServiceCode string `json:",omitempty"` // コピー先のiarのサービスコード
	ImageId        string `json:",omitempty"` // コピー先でのイメージのID
}

package protocol

import (
	"reflect"
)

// StorageImageCopy ストレージイメージコピー
//  http://manual.iij.jp/p2/pubapi/137048215.html
type StorageImageCopy struct {
	SrcGisServiceCode string `json:"-"` // P2契約のサービスコード(gis########)
	SrcIarServiceCode string `json:"-"`
	SrcImageId        string `json:"-"`
	DstGisServiceCode string `json:"GisServiceCode"`
	DstIarServiceCode string `json:"IarServiceCode"`
	Image             string 
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
	GisServiceCode string `json:",omitempty"`
	IarServiceCode string `json:",omitempty"`
	ImageId        string `json:",omitempty"`
}

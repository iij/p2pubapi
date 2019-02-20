package protocol

import (
	"reflect"
)

// Restore リストア (非同期)
//  http://manual.iij.jp/p2/pubapi/59940054.html
type Restore struct {
	GisServiceCode     string `json:"-"` // P2契約のサービスコード(gis########)
	StorageServiceCode string `json:"-"` // システムストレージのサービスコード(iba########/ica########)
	ImageId            string // カスタムOSイメージID(数値)
	IarServiceCode     string // ストレージアーカイブのサービスコード(iar########)
	Image              string // 操作内容(Archive)
}

// URI /{{.GisServiceCode}}/system-storages/{{.StorageServiceCode}}/action.json
func (t Restore) URI() string {
	return "/{{.GisServiceCode}}/system-storages/{{.StorageServiceCode}}/action.json"
}

// APIName Restore
func (t Restore) APIName() string {
	return "Restore"
}

// Method PUT
func (t Restore) Method() string {
	return "PUT"
}

// Document http://manual.iij.jp/p2/pubapi/59940054.html
func (t Restore) Document() string {
	return "http://manual.iij.jp/p2/pubapi/59940054.html"
}

// JPName リストア
func (t Restore) JPName() string {
	return "リストア"
}
func init() {
	APIlist = append(APIlist, Restore{})
	TypeMap["Restore"] = reflect.TypeOf(Restore{})
}

// RestoreResponse リストアのレスポンス
type RestoreResponse struct {
	*CommonResponse
	Current struct {
		ResourceStatus string `json:",omitempty"` // 設定後のストレージステータス
	} `json:",omitempty"`
	Previous struct {
		ResourceStatus string `json:",omitempty"` // 設定前のストレージステータス
	} `json:",omitempty"`
}

package protocol

import (
	"reflect"
)

// UserDataSet システムストレージ ユーザーデータ設定
// http://manual.iij.jp/p2/pubapi/137048177.html
type UserDataSet struct {
	GisServiceCode     string `json:"-"` // ユーザデータを設定する仮想サーバが所属するgisのサービスコード(gis########)
	StorageServiceCode string `json:"-"` // ユーザデータを設定するシステムストレージのサービスコード(iba########/ica########)
	UserData           string // base64エンコードされたユーザデータ
}

// URI /{{.GisServiceCode}}/system-storages/{{.StorageServiceCode}}/user-data.json
func (t UserDataSet) URI() string {
	return "/{{.GisServiceCode}}/system-storages/{{.StorageServiceCode}}/user-data.json"
}

// APIName UserDataSet
func (t UserDataSet) APIName() string {
	return "UserDataSet"
}

// Method PUT
func (t UserDataSet) Method() string {
	return "PUT"
}

// Document http://manual.iij.jp/p2/pubapi/137048177.html
func (t UserDataSet) Document() string {
	return "http://manual.iij.jp/p2/pubapi/137048177.html"
}

// JPName システムストレージユーザデータ設定
func (t UserDataSet) JPName() string {
	return "システムストレージユーザデータ設定"
}
func init() {
	APIlist = append(APIlist, UserDataSet{})
	TypeMap["UserDataSet"] = reflect.TypeOf(UserDataSet{})
}

// UserDataSetResponse
type UserDataSetResponse struct {
	*CommonResponse
}

package protocol

import (
	"reflect"
)

// UserDataGet システムストレージ ユーザーデータ設定
// http://manual.iij.jp/p2/pubapi/137048177.html
type UserDataGet struct {
	GisServiceCode string `json:"-"` // P2契約のサービスコード(gis########)
	IbaServiceCode string `json:"-"`
}

// URI /{{.GisServiceCode}}/system-storages/{{.IbaServiceCode}}/user-data.json
func (t UserDataGet) URI() string {
	return "/{{.GisServiceCode}}/system-storages/{{.IbaServiceCode}}/user-data.json"
}

// APIName UserDataGet
func (t UserDataGet) APIName() string {
	return "UserDataGet"
}

// Method PUT
func (t UserDataGet) Method() string {
	return "GET"
}

// Document http://manual.iij.jp/p2/pubapi/137048129.html
func (t UserDataGet) Document() string {
	return "http://manual.iij.jp/p2/pubapi/137048129.html"
}

// JPName システムストレージユーザデータ取得
func (t UserDataGet) JPName() string {
	return "システムストレージユーザデータ取得"
}
func init() {
	APIlist = append(APIlist, UserDataGet{})
	TypeMap["UserDataGet"] = reflect.TypeOf(UserDataGet{})
}

// UserDataGetResponse
type UserDataGetResponse struct {
	*CommonResponse
	UserData  string `json:",omitempty"`
}


package protocol

import (
	"reflect"
)

// FwLbResourceGet FW+LBリソース状態取得 (同期)
//  http://manual.iij.jp/p2/pubapi/59940962.html
type FwLbResourceGet struct {
	GisServiceCode string `json:"-"`                // P2契約のサービスコード(gis########)
	IflServiceCode string `json:"-"`                // IflServiceCode
	Item           string `json:"-" p2pub:",query"` // 取得するフィールド("ResourceStatus")
}

// URI /{{.GisServiceCode}}/fw-lbs/{{.IflServiceCode}}.json
func (t FwLbResourceGet) URI() string {
	return "/{{.GisServiceCode}}/fw-lbs/{{.IflServiceCode}}.json"
}

// APIName FwLbResourceGet
func (t FwLbResourceGet) APIName() string {
	return "FwLbResourceGet"
}

// Method GET
func (t FwLbResourceGet) Method() string {
	return "GET"
}

// Document http://manual.iij.jp/p2/pubapi/59940962.html
func (t FwLbResourceGet) Document() string {
	return "http://manual.iij.jp/p2/pubapi/59940962.html"
}

// JPName FW+LBリソース状態取得
func (t FwLbResourceGet) JPName() string {
	return "FW+LBリソース状態取得"
}
func init() {
	APIlist = append(APIlist, FwLbResourceGet{})
	TypeMap["FwLbResourceGet"] = reflect.TypeOf(FwLbResourceGet{})
}

// FwLbResourceGetResponse FW+LBリソース状態取得のレスポンス
type FwLbResourceGetResponse struct {
	*CommonResponse
	ResourceStatus string `json:",omitempty"` // FW+LB 専有タイプステータス
	HostList       []struct {
		Master         string `json:",omitempty"` // マスターならばYes, スレーブならばNo(Yes" "No)
		ResourceStatus string `json:",omitempty"` // FW+LB 専有タイプステータス
	} `json:",omitempty"`
}

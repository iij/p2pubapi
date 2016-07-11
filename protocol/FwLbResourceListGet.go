package protocol

import (
	"reflect"
)

// FwLbResourceListGet FW+LBリソース状態一覧取得 (同期)
//  http://manual.iij.jp/p2/pubapi/59940955.html
type FwLbResourceListGet struct {
	GisServiceCode string `json:"-"`                // P2契約のサービスコード(gis########)
	Item           string `json:"-" p2pub:",query"` // 取得するフィールド("ResourceStatus")
	StartIndex     string `json:"-" p2pub:",query"` // StartIndexだけスキップした位置から情報を取得する(数値)
	Count          string `json:"-" p2pub:",query"` // 取得する情報の最大数(数値)
}

// URI /{{.GisServiceCode}}/fw-lbs.json
func (t FwLbResourceListGet) URI() string {
	return "/{{.GisServiceCode}}/fw-lbs.json"
}

// APIName FwLbResourceListGet
func (t FwLbResourceListGet) APIName() string {
	return "FwLbResourceListGet"
}

// Method GET
func (t FwLbResourceListGet) Method() string {
	return "GET"
}

// Document http://manual.iij.jp/p2/pubapi/59940955.html
func (t FwLbResourceListGet) Document() string {
	return "http://manual.iij.jp/p2/pubapi/59940955.html"
}

// JPName FW+LBリソース状態一覧取得
func (t FwLbResourceListGet) JPName() string {
	return "FW+LBリソース状態一覧取得"
}
func init() {
	APIlist = append(APIlist, FwLbResourceListGet{})
	TypeMap["FwLbResourceListGet"] = reflect.TypeOf(FwLbResourceListGet{})
}

// FwLbResourceListGetResponse FW+LBリソース状態一覧取得のレスポンス
type FwLbResourceListGetResponse struct {
	*CommonResponse
	FwLbList []struct {
		ResourceStatus string `json:",omitempty"` // FW+LB 専有タイプステータス
		HostList       []struct {
			Master         string `json:",omitempty"` // マスターならばYes, スレーブならばNo(Yes" "No)
			ResourceStatus string `json:",omitempty"` // FW+LB 専有タイプステータス
		} `json:",omitempty"`
		ServiceCode string `json:",omitempty"` // FW+LB 専有タイプのサービスコード(ifl########)
	} `json:",omitempty"`
}

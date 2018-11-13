package protocol

import (
	"reflect"
)

// FwLbBestEffortResourceGet FW+LBベストエフォートタイプリソース状態取得 (同期)
//  http://manual.iij.jp/p2/pubapi/162303028.html
type FwLbBestEffortResourceGet struct {
	GisServiceCode string `json:"-"`                // P2契約のサービスコード(gis########)
	IlbServiceCode string `json:"-"`                // IlbServiceCode
	Item           string `json:"-" p2pub:",query"` // 取得するフィールド("ResourceStatus")
}

// URI /{{.GisServiceCode}}/best-effort-fw-lbs/{{.IlbServiceCode}}.json
func (t FwLbBestEffortResourceGet) URI() string {
	return "/{{.GisServiceCode}}/best-effort-fw-lbs/{{.IlbServiceCode}}.json"
}

// APIName FwLbBestEffortResourceGet
func (t FwLbBestEffortResourceGet) APIName() string {
	return "FwLbBestEffortResourceGet"
}

// Method GET
func (t FwLbBestEffortResourceGet) Method() string {
	return "GET"
}

// Document http://manual.iij.jp/p2/pubapi/162303028.html
func (t FwLbBestEffortResourceGet) Document() string {
	return "http://manual.iij.jp/p2/pubapi/162303028.html"
}

// JPName FW+LBベストエフォートタイプリソース状態取得
func (t FwLbBestEffortResourceGet) JPName() string {
	return "FW+LBベストエフォートタイプリソース状態取得"
}
func init() {
	APIlist = append(APIlist, FwLbBestEffortResourceGet{})
	TypeMap["FwLbBestEffortResourceGet"] = reflect.TypeOf(FwLbBestEffortResourceGet{})
}

// FwLbBestEffortResourceGetResponse FW+LBベストエフォートタイプリソース状態取得のレスポンス
type FwLbBestEffortResourceGetResponse struct {
	*CommonResponse
	ResourceStatus string `json:",omitempty"` // FW+LB ベストエフォートタイプステータス
	HostList       []struct {
		Master         string `json:",omitempty"` // マスターならばYes, スレーブならばNo(Yes" "No)
		ResourceStatus string `json:",omitempty"` // FW+LB ベストエフォートタイプステータス
	} `json:",omitempty"`
}

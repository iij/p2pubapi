package protocol

import (
	"reflect"
)

// FwLbBestEffortResourceListGet FW+LBベストエフォートタイプリソース状態一覧取得 (同期)
//  http://manual.iij.jp/p2/pubapi/162303028.html
type FwLbBestEffortResourceListGet struct {
	GisServiceCode string `json:"-"`                // P2契約のサービスコード(gis########)
	Item           string `json:"-" p2pub:",query"` // 取得するフィールド("ResourceStatus")
	StartIndex     string `json:"-" p2pub:",query"` // StartIndexだけスキップした位置から情報を取得する(数値)
	Count          string `json:"-" p2pub:",query"` // 取得する情報の最大数(数値)
}

// URI /{{.GisServiceCode}}/best-effort-fw-lbs.json
func (t FwLbBestEffortResourceListGet) URI() string {
	return "/{{.GisServiceCode}}/best-effort-fw-lbs.json"
}

// APIName FwLbBestEffortResourceListGet
func (t FwLbBestEffortResourceListGet) APIName() string {
	return "FwLbBestEffortResourceListGet"
}

// Method GET
func (t FwLbBestEffortResourceListGet) Method() string {
	return "GET"
}

// Document http://manual.iij.jp/p2/pubapi/162303028.html
func (t FwLbBestEffortResourceListGet) Document() string {
	return "http://manual.iij.jp/p2/pubapi/162303028.html"
}

// JPName FW+LBベストエフォートタイプリソース状態一覧取得
func (t FwLbBestEffortResourceListGet) JPName() string {
	return "FW+LBベストエフォートタイプリソース状態一覧取得"
}
func init() {
	APIlist = append(APIlist, FwLbBestEffortResourceListGet{})
	TypeMap["FwLbBestEffortResourceListGet"] = reflect.TypeOf(FwLbBestEffortResourceListGet{})
}

// FwLbBestEffortResourceListGetResponse FW+LBベストエフォートタイプリソース状態一覧取得のレスポンス
type FwLbBestEffortResourceListGetResponse struct {
	*CommonResponse
	BestEffortFwLbList []struct {
		ResourceStatus string `json:",omitempty"` // FW+LB ベストエフォートタイプステータス
		HostList       []struct {
			Master         string `json:",omitempty"` // マスターならばYes, スレーブならばNo(Yes" "No)
			ResourceStatus string `json:",omitempty"` // FW+LB ベストエフォートタイプステータス
		} `json:",omitempty"`
		ServiceCode string `json:",omitempty"` // FW+LB ベストエフォートタイプのサービスコード(ifl########)
	} `json:",omitempty"`
}

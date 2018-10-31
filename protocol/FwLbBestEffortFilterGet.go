package protocol

import (
	"reflect"
)

// FwLbBestEffortFilterGet フィルタリングルール情報取得（FW+LBベストエフォートタイプ） (同期)
//  http://manual.iij.jp/p2/pubapi/162646975.html
type FwLbBestEffortFilterGet struct {
	GisServiceCode string `json:"-"` // P2契約のサービスコード(gis########)
	IlbServiceCode string `json:"-"` // FW+LB ベストエフォートタイプのサービスコード(ilb########)
	IpVersion      string `json:"-"` // IPv4ルール（v4）
	Direction      string `json:"-"` // 入力方向（ExternalからInternalへ）のルールか出力方向のルールか（in, out）
}

// URI /:GisServiceCode/best-effort-fw-lbs/:IlbServiceCode/filters/:IpVersion/:Direction.json
func (t FwLbBestEffortFilterGet) URI() string {
	return "/{{.GisServiceCode}}/best-effort-fw-lbs/{{.IlbServiceCode}}/filters/{{.IpVersion}}/{{.Direction}}.json"
}

// APIName FwLbBestEffortFilterGet
func (t FwLbBestEffortFilterGet) APIName() string {
	return "FwLbBestEffortFilterGet"
}

// Method GET
func (t FwLbBestEffortFilterGet) Method() string {
	return "GET"
}

// http://manual.iij.jp/p2/pubapi/162646975.html
func (t FwLbBestEffortFilterGet) Document() string {
	return "http://manual.iij.jp/p2/pubapi/162646975.html"
}

// JPName フィルタリングルール情報取得（FW+LBベストエフォートタイプ）
func (t FwLbBestEffortFilterGet) JPName() string {
	return "フィルタリングルール情報取得（FW+LBベストエフォートタイプ）"
}
func init() {
	APIlist = append(APIlist, FwLbBestEffortFilterGet{})
	TypeMap["FwLbBestEffortFilterGet"] = reflect.TypeOf(FwLbBestEffortFilterGet{})
}

// FwLbBestEffortFilterGetResponse フィルタリングルール情報取得（FW+LBベストエフォートタイプ）のレスポンス
type FwLbBestEffortFilterGetResponse struct {
	*CommonResponse
	FilterRuleList []struct {
		FilterId           string // フィルタID
		SourceNetwork      string // ソースネットワーク
		DestinationNetwork string // デスティネーションネットワーク
		DestinationPort    string // デスティネーションポート番号
		Protocol           string // プロトコル
		Action             string // ルールにマッチしたパケットに対する処理
		Label              string // ラベル
	}
}

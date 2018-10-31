package protocol

import (
	"reflect"
)

// FwLbBestEffortDestinationGet ロードバランス先情報取得
//  http://manual.iij.jp/p2/pubapi/162315120.html
type FwLbBestEffortDestinationGet struct {
	GisServiceCode string `json:"-"` // P2契約のサービスコード(gis########)
	IlbServiceCode string `json:"-"` // FW+LB ベストエフォートタイプのサービスコード(ilb########)
	Id             string `json:"-"` // ロードバランス先を一意に識別するID
}

// URI /:GisServiceCode/best-effort-fw-lbs/:IlbServiceCode/lb/destinations/:Id.json
func (t FwLbBestEffortDestinationGet) URI() string {
	return "/{{.GisServiceCode}}/best-effort-fw-lbs/{{.IlbServiceCode}}/lb/destinations/{{.Id}}.json"
}

// APIName FwLbBestEffortDestinationGet
func (t FwLbBestEffortDestinationGet) APIName() string {
	return "FwLbBestEffortDestinationGet"
}

// Method GET
func (t FwLbBestEffortDestinationGet) Method() string {
	return "GET"
}

// http://manual.iij.jp/p2/pubapi/162315120.html
func (t FwLbBestEffortDestinationGet) Document() string {
	return "http://manual.iij.jp/p2/pubapi/162315120.html"
}

// JPName ロードバランス先情報取得
func (t FwLbBestEffortDestinationGet) JPName() string {
	return "ロードバランス先情報取得"
}

func init() {
	APIlist = append(APIlist, FwLbBestEffortDestinationGet{})
	TypeMap["FwLbBestEffortDestinationGet"] = reflect.TypeOf(FwLbBestEffortDestinationGet{})
}

// FwLbBestEffortDestinationGetResponse ロードバランス先情報取得 のレスポンス
type FwLbBestEffortDestinationGetResponse struct {
	*CommonResponse
	LbConfigId        string // ロードバランシング設定のID
	IpAddress         string // ロードバランス先アドレス
	Port              string // ロードバランス先ポート番号
	Weight            string // ロードバランス先重み
	Failover          string // フェールオーバー先
	Enabled           string // 有効・無効
	HealthCheckStatus string // ヘルスチェックステータス
}

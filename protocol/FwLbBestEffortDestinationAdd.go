package protocol

import (
	"reflect"
)

// FwLbBestEffortDestinationAdd ロードバランス先追加
//  http://manual.iij.jp/p2/pubapi/162315113.html
type FwLbBestEffortDestinationAdd struct {
	GisServiceCode string `json:"-"` // P2契約のサービスコード(gis########)
	IlbServiceCode string `json:"-"` // FW+LB ベストエフォートタイプのサービスコード(ilb########)
	LbConfigId     string // ロードバランシング設定のID
	IpAddress      string // ロードバランス先アドレス
	Port           string // ロードバランス先ポート番号
	Weight         string `json:",omitempty"` // ロードバランス先重み
	Failover       string // フェールオーバー先
	Enabled        string // 有効・無効
}

// URI /:GisServiceCode/best-effort-fw-lbs/:IlbServiceCode/lb/destinations.json
func (t FwLbBestEffortDestinationAdd) URI() string {
	return "/{{.GisServiceCode}}/best-effort-fw-lbs/{{.IlbServiceCode}}/lb/destinations.json"
}

// APIName FwLbBestEffortDestinationAdd
func (t FwLbBestEffortDestinationAdd) APIName() string {
	return "FwLbBestEffortDestinationAdd"
}

// Method POST
func (t FwLbBestEffortDestinationAdd) Method() string {
	return "POST"
}

// http://manual.iij.jp/p2/pubapi/162315113.html
func (t FwLbBestEffortDestinationAdd) Document() string {
	return "http://manual.iij.jp/p2/pubapi/162315113.html"
}

// JPName ロードバランス先追加
func (t FwLbBestEffortDestinationAdd) JPName() string {
	return "ロードバランス先追加"
}

func init() {
	APIlist = append(APIlist, FwLbBestEffortDestinationAdd{})
	TypeMap["FwLbBestEffortDestinationAdd"] = reflect.TypeOf(FwLbBestEffortDestinationAdd{})
}

// FwLbBestEffortDestinationAddResponse ロードバランス先追加 のレスポンス
type FwLbBestEffortDestinationAddResponse struct {
	*CommonResponse
	Id                string // ロードバランス先を一意に識別するID
	LbConfigId        string // ロードバランシング設定のID
	IpAddress         string // ロードバランス先アドレス
	Port              string // ロードバランス先ポート番号
	Weight            string // ロードバランス先重み
	Failover          string // フェールオーバー先
	Enabled           string // 有効・無効
	HealthCheckStatus string // ヘルスチェックステータス
}

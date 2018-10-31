package protocol

import (
	"reflect"
)

// FwLbBestEffortDestinationDelete ロードバランス先削除
//  http://manual.iij.jp/p2/pubapi/162315135.html
type FwLbBestEffortDestinationDelete struct {
	GisServiceCode string `json:"-"` // P2契約のサービスコード(gis########)
	IlbServiceCode string `json:"-"` // FW+LB ベストエフォートタイプのサービスコード(ilb########)
	Id             string `json:"-"` // ロードバランス先を一意に識別するID
}

// URI /:GisServiceCode/best-effort-fw-lbs/:IlbServiceCode/lb/destinations/:Id.json
func (t FwLbBestEffortDestinationDelete) URI() string {
	return "/{{.GisServiceCode}}/best-effort-fw-lbs/{{.IlbServiceCode}}/lb/destinations/{{.Id}}.json"
}

// APIName FwLbBestEffortDestinationDelete
func (t FwLbBestEffortDestinationDelete) APIName() string {
	return "FwLbBestEffortDestinationDelete"
}

// Method DELETE
func (t FwLbBestEffortDestinationDelete) Method() string {
	return "DELETE"
}

// http://manual.iij.jp/p2/pubapi/162315135.html
func (t FwLbBestEffortDestinationDelete) Document() string {
	return "http://manual.iij.jp/p2/pubapi/162315135.html"
}

// JPName ロードバランス先削除
func (t FwLbBestEffortDestinationDelete) JPName() string {
	return "ロードバランス先削除"
}

func init() {
	APIlist = append(APIlist, FwLbBestEffortDestinationDelete{})
	TypeMap["FwLbBestEffortDestinationDelete"] = reflect.TypeOf(FwLbBestEffortDestinationDelete{})
}

// FwLbBestEffortDestinationDeleteResponse ロードバランス先削除 のレスポンス
type FwLbBestEffortDestinationDeleteResponse struct {
	*CommonResponse
	LbConfigId        string // ロードバランシング設定のID
	IpAddress         string // ロードバランス先アドレス
	Port              string // ロードバランス先ポート番号
	Weight            string // ロードバランス先重み
	Failover          string // フェールオーバー先
	Enabled           string // 有効・無効
	HealthCheckStatus string // ヘルスチェックステータス
}

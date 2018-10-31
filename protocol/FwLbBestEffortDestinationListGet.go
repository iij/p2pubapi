package protocol

import (
	"reflect"
)

// FwLbBestEffortDestinationListGet ロードバランス先情報一覧取得
//  http://manual.iij.jp/p2/pubapi/162315096.html
type FwLbBestEffortDestinationListGet struct {
	GisServiceCode string `json:"-"` // P2契約のサービスコード(gis########)
	IlbServiceCode string `json:"-"` // FW+LBベストエフォートタイプのサービスコード
}

// URI /:GisServiceCode/best-effort-fw-lbs/:IlbServiceCode/lb/destinations.json
func (t FwLbBestEffortDestinationListGet) URI() string {
	return "/{{.GisServiceCode}}/best-effort-fw-lbs/{{.IlbServiceCode}}/lb/destinations.json"
}

// APIName FwLbBestEffortDestinationListGet
func (t FwLbBestEffortDestinationListGet) APIName() string {
	return "FwLbBestEffortDestinationListGet"
}

// Method GET
func (t FwLbBestEffortDestinationListGet) Method() string {
	return "GET"
}

// http://manual.iij.jp/p2/pubapi/162315096.html
func (t FwLbBestEffortDestinationListGet) Document() string {
	return "http://manual.iij.jp/p2/pubapi/162315096.html"
}

// JPName ロードバランス先情報一覧取得
func (t FwLbBestEffortDestinationListGet) JPName() string {
	return "ロードバランス先情報一覧取得"
}

func init() {
	APIlist = append(APIlist, FwLbBestEffortDestinationListGet{})
	TypeMap["FwLbBestEffortDestinationListGet"] = reflect.TypeOf(FwLbBestEffortDestinationListGet{})
}

// FwLbBestEffortDestinationListGetResponse ロードバランス先情報一覧取得 のレスポンス
type FwLbBestEffortDestinationListGetResponse struct {
	*CommonResponse
	DestinationList []struct {
		Id                string // ロードバランス先ID
		LbConfigId        string // ロードバランシング設定のID
		IpAddress         string // ロードバランス先アドレス
		Port              string // ロードバランス先ポート番号
		Weight            string // ロードバランス先重み
		Failover          string // フェールオーバー先
		Enabled           string // 有効・無効
		HealthCheckStatus string // ヘルスチェックステータス
	}
}

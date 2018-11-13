package protocol

import (
	"reflect"
)

// FwLbBestEffortConfigDelete ロードバランシング設定削除
//  http://manual.iij.jp/p2/pubapi/162312797.html
type FwLbBestEffortConfigDelete struct {
	GisServiceCode string `json:"-"` // P2契約のサービスコード(gis########)
	IlbServiceCode string `json:"-"` // FW+LB ベストエフォートタイプのサービスコード(ilb########)
	Id             string `json:"-"` // ロードバランシング設定のID
}

// URI /:GisServiceCode/best-effort-fw-lbs/:IlbServiceCode/lb/configs/:Id.json
func (t FwLbBestEffortConfigDelete) URI() string {
	return "/{{.GisServiceCode}}/best-effort-fw-lbs/{{.IlbServiceCode}}/lb/configs/{{.Id}}.json"
}

// APIName FwLbBestEffortConfigDelete
func (t FwLbBestEffortConfigDelete) APIName() string {
	return "FwLbBestEffortConfigDelete"
}

// Method GET
func (t FwLbBestEffortConfigDelete) Method() string {
	return "DELETE"
}

// http://manual.iij.jp/p2/pubapi/162312797.html
func (t FwLbBestEffortConfigDelete) Document() string {
	return "http://manual.iij.jp/p2/pubapi/162312797.html"
}

// JPName ロードバランシング設定削除
func (t FwLbBestEffortConfigDelete) JPName() string {
	return "ロードバランシング設定削除"
}

func init() {
	APIlist = append(APIlist, FwLbBestEffortConfigDelete{})
	TypeMap["FwLbBestEffortConfigDelete"] = reflect.TypeOf(FwLbBestEffortConfigDelete{})
}

// FwLbBestEffortConfigDeleteResponse ロードバランシング設定削除 のレスポンス
type FwLbBestEffortConfigDeleteResponse struct {
	*CommonResponse
	Id                string // ロードバランシング設定のID
	Name              string // 設定名
	TrafficIpName     string // TrafficIp の名前
	Port              string // ポート番号
	Protocol          string // プロトコル
	Algorithm         string // 負荷分散方式
	IpTransparent     string // IPアドレス透過
	HealthCheckMethod string // ヘルスチェック方式
}

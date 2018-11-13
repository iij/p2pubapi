package protocol

import (
	"reflect"
)

// FwLbBestEffortConfigAdd ロードバランシング設定追加
//  http://manual.iij.jp/p2/pubapi/162312646.html
type FwLbBestEffortConfigAdd struct {
	GisServiceCode    string `json:"-"` // P2契約のサービスコード(gis########)
	IlbServiceCode    string `json:"-"` // FW+LB ベストエフォートタイプのサービスコード(ilb########)
	Name              string // 設定名
	TrafficIpName     string // TrafficIp の名前
	Port              string // ポート番号
	Protocol          string // プロトコル
	Algorithm         string // 負荷分散方式
	IpTransparent     string // IPアドレス透過
	HealthCheckMethod string // ヘルスチェック方式
}

// URI /:GisServiceCode/best-effort-fw-lbs/:IlbServiceCode/lb/configs.json
func (t FwLbBestEffortConfigAdd) URI() string {
	return "/{{.GisServiceCode}}/best-effort-fw-lbs/{{.IlbServiceCode}}/lb/configs.json"
}

// APIName FwLbBestEffortConfigAdd
func (t FwLbBestEffortConfigAdd) APIName() string {
	return "FwLbBestEffortConfigAdd"
}

// Method GET
func (t FwLbBestEffortConfigAdd) Method() string {
	return "POST"
}

// http://manual.iij.jp/p2/pubapi/162312646.html
func (t FwLbBestEffortConfigAdd) Document() string {
	return "http://manual.iij.jp/p2/pubapi/162312646.html"
}

// JPName ロードバランシング設定追加
func (t FwLbBestEffortConfigAdd) JPName() string {
	return "ロードバランシング設定追加"
}

func init() {
	APIlist = append(APIlist, FwLbBestEffortConfigAdd{})
	TypeMap["FwLbBestEffortConfigAdd"] = reflect.TypeOf(FwLbBestEffortConfigAdd{})
}

// FwLbBestEffortConfigAddResponse ロードバランシング設定追加 のレスポンス
type FwLbBestEffortConfigAddResponse struct {
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

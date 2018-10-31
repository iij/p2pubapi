package protocol

import (
	"reflect"
)

// FwLbBestEffortConfigGet ロードバランシング設定情報取得
//  http://manual.iij.jp/p2/pubapi/162312652.html
type FwLbBestEffortConfigGet struct {
	GisServiceCode string `json:"-"` // P2契約のサービスコード(gis########)
	IlbServiceCode string `json:"-"` // FW+LB ベストエフォートタイプのサービスコード(ilb########)
	Id             string `json:"-"` // ロードバランシング設定のID
}

// URI /:GisServiceCode/best-effort-fw-lbs/:IlbServiceCode/lb/configs/Id.json
func (t FwLbBestEffortConfigGet) URI() string {
	return "/{{.GisServiceCode}}/best-effort-fw-lbs/{{.IlbServiceCode}}/lb/configs/{{.Id}}.json"
}

// APIName FwLbBestEffortConfigGet
func (t FwLbBestEffortConfigGet) APIName() string {
	return "FwLbBestEffortConfigGet"
}

// Method GET
func (t FwLbBestEffortConfigGet) Method() string {
	return "GET"
}

// http://manual.iij.jp/p2/pubapi/162312652.html
func (t FwLbBestEffortConfigGet) Document() string {
	return "http://manual.iij.jp/p2/pubapi/162312652.html"
}

// JPName ロードバランシング設定情報取得
func (t FwLbBestEffortConfigGet) JPName() string {
	return "ロードバランシング設定情報取得"
}

func init() {
	APIlist = append(APIlist, FwLbBestEffortConfigGet{})
	TypeMap["FwLbBestEffortConfigGet"] = reflect.TypeOf(FwLbBestEffortConfigGet{})
}

// FwLbBestEffortConfigGetResponse ロードバランシング設定情報取得 のレスポンス
type FwLbBestEffortConfigGetResponse struct {
	*CommonResponse
	Name              string // 設定名
	TrafficIpName     string // TrafficIp の名前
	Port              string // ポート番号
	Protocol          string // プロトコル
	Algorithm         string // 負荷分散方式
	IpTransparent     string // IPアドレス透過
	HealthCheckMethod string // ヘルスチェック方式
}

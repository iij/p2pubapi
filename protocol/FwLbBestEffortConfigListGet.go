package protocol

import (
	"reflect"
)

// FwLbBestEffortConfigListGet ロードバランシング設定情報一覧取得
//  http://manual.iij.jp/p2/pubapi/162312601.html
type FwLbBestEffortConfigListGet struct {
	GisServiceCode string `json:"-"` // P2契約のサービスコード(gis########)
	IlbServiceCode string `json:"-"` // FW+LBベストエフォートタイプのサービスコード
}

// URI /:GisServiceCode/best-effort-fw-lbs/:IlbServiceCode/lb/configs.json
func (t FwLbBestEffortConfigListGet) URI() string {
	return "/{{.GisServiceCode}}/best-effort-fw-lbs/{{.IlbServiceCode}}/lb/configs.json"
}

// APIName FwLbBestEffortConfigListGet
func (t FwLbBestEffortConfigListGet) APIName() string {
	return "FwLbBestEffortConfigListGet"
}

// Method GET
func (t FwLbBestEffortConfigListGet) Method() string {
	return "GET"
}

// http://manual.iij.jp/p2/pubapi/162312601.html
func (t FwLbBestEffortConfigListGet) Document() string {
	return "http://manual.iij.jp/p2/pubapi/162312601.html"
}

// JPName ロードバランシング設定情報一覧取得
func (t FwLbBestEffortConfigListGet) JPName() string {
	return "ロードバランシング設定情報一覧取得"
}

func init() {
	APIlist = append(APIlist, FwLbBestEffortConfigListGet{})
	TypeMap["FwLbBestEffortConfigListGet"] = reflect.TypeOf(FwLbBestEffortConfigListGet{})
}

// FwLbBestEffortConfigListGetResponse ロードバランシング設定情報一覧取得 のレスポンス
type FwLbBestEffortConfigListGetResponse struct {
	*CommonResponse
	LbConfigList []struct {
		Id                string // ロードバランシング設定を一意に識別するID
		Name              string // 設定名
		TrafficIpName     string // TrafficIp の名前
		Port              string // ポート番号
		Protocol          string // プロトコル
		Algorithm         string // 負荷分散方式
		IpTransparent     string // IPアドレス透過
		HealthCheckMethod string // ヘルスチェック方式
	}
}

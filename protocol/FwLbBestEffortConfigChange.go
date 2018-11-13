package protocol

import (
	"reflect"
)

// FwLbBestEffortConfigChange ロードバランシング設定変更
//  http://manual.iij.jp/p2/pubapi/162312792.html
type FwLbBestEffortConfigChange struct {
	GisServiceCode    string `json:"-"` // P2契約のサービスコード(gis########)
	IlbServiceCode    string `json:"-"` // FW+LB ベストエフォートタイプのサービスコード(ilb########)
	Id                string `json:"-"` // ロードバランシング設定のID
	Name              string // 設定名
	Protocol          string // プロトコル
	Algorithm         string // 負荷分散方式
	IpTransparent     string // IPアドレス透過
	HealthCheckMethod string // ヘルスチェック方式
}

// URI /:GisServiceCode/best-effort-fw-lbs/:IlbServiceCode/lb/configs/Id.json
func (t FwLbBestEffortConfigChange) URI() string {
	return "/{{.GisServiceCode}}/best-effort-fw-lbs/{{.IlbServiceCode}}/lb/configs/{{.Id}}.json"
}

// APIName FwLbBestEffortConfigChange
func (t FwLbBestEffortConfigChange) APIName() string {
	return "FwLbBestEffortConfigChange"
}

// Method GET
func (t FwLbBestEffortConfigChange) Method() string {
	return "PUT"
}

// http://manual.iij.jp/p2/pubapi/162312792.html
func (t FwLbBestEffortConfigChange) Document() string {
	return "http://manual.iij.jp/p2/pubapi/162312792.html"
}

// JPName ロードバランシング設定変更
func (t FwLbBestEffortConfigChange) JPName() string {
	return "ロードバランシング設定変更"
}

func init() {
	APIlist = append(APIlist, FwLbBestEffortConfigChange{})
	TypeMap["FwLbBestEffortConfigChange"] = reflect.TypeOf(FwLbBestEffortConfigChange{})
}

// FwLbBestEffortConfigChangeResponse ロードバランシング設定変更 のレスポンス
type FwLbBestEffortConfigChangeResponse struct {
	*CommonResponse
	Current struct {
		Name              string // 設定名
		Protocol          string // プロトコル
		Algorithm         string // 負荷分散方式
		IpTransparent     string // IPアドレス透過
		HealthCheckMethod string // ヘルスチェック方式
	}
	Previous struct {
		Name              string // 設定名
		Protocol          string // プロトコル
		Algorithm         string // 負荷分散方式
		IpTransparent     string // IPアドレス透過
		HealthCheckMethod string // ヘルスチェック方式
	}
}

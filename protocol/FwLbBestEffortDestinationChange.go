package protocol

import (
	"reflect"
)

// FwLbBestEffortDestinationChange ロードバランス先設定変更
//  http://manual.iij.jp/p2/pubapi/162315127.html
type FwLbBestEffortDestinationChange struct {
	GisServiceCode string `json:"-"`          // P2契約のサービスコード(gis########)
	IlbServiceCode string `json:"-"`          // FW+LB ベストエフォートタイプのサービスコード(ilb########)
	Id             string `json:"-"`          // ロードバランス先を一意に識別するID
	Weight         string `json:",omitempty"` // ロードバランス先重み
	Enabled        string // 有効・無効
}

// URI /:GisServiceCode/best-effort-fw-lbs/:IlbServiceCode/lb/destinations/:Id.json
func (t FwLbBestEffortDestinationChange) URI() string {
	return "/{{.GisServiceCode}}/best-effort-fw-lbs/{{.IlbServiceCode}}/lb/destinations/{{.Id}}.json"
}

// APIName FwLbBestEffortDestinationChange
func (t FwLbBestEffortDestinationChange) APIName() string {
	return "FwLbBestEffortDestinationChange"
}

// Method PUT
func (t FwLbBestEffortDestinationChange) Method() string {
	return "PUT"
}

// http://manual.iij.jp/p2/pubapi/162315127.html
func (t FwLbBestEffortDestinationChange) Document() string {
	return "http://manual.iij.jp/p2/pubapi/162315127.html"
}

// JPName ロードバランス先設定変更
func (t FwLbBestEffortDestinationChange) JPName() string {
	return "ロードバランス先設定変更"
}

func init() {
	APIlist = append(APIlist, FwLbBestEffortDestinationChange{})
	TypeMap["FwLbBestEffortDestinationChange"] = reflect.TypeOf(FwLbBestEffortDestinationChange{})
}

// FwLbBestEffortDestinationChangeResponse ロードバランス先設定変更 のレスポンス
type FwLbBestEffortDestinationChangeResponse struct {
	*CommonResponse
	Current struct {
		Weight  string // ロードバランス先重み
		Enabled string // 有効・無効
	}
	Previous struct {
		Weight  string // ロードバランス先重み
		Enabled string // 有効・無効
	}
}

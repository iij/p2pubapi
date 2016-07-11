package protocol

import (
	"reflect"
)

// LBControlPanelACLGet LB管理画面アクセス制限情報取得 (同期)
//  http://manual.iij.jp/p2/pubapi/59940994.html
type LBControlPanelACLGet struct {
	GisServiceCode string `json:"-"` // P2契約のサービスコード(gis########)
	IflServiceCode string `json:"-"` // FW+LB 専有タイプのサービスコード(ifl########)
}

// URI /{{.GisServiceCode}}/fw-lbs/{{.IflServiceCode}}/lb/adminserver/allows.json
func (t LBControlPanelACLGet) URI() string {
	return "/{{.GisServiceCode}}/fw-lbs/{{.IflServiceCode}}/lb/adminserver/allows.json"
}

// APIName LBControlPanelACLGet
func (t LBControlPanelACLGet) APIName() string {
	return "LBControlPanelACLGet"
}

// Method GET
func (t LBControlPanelACLGet) Method() string {
	return "GET"
}

// Document http://manual.iij.jp/p2/pubapi/59940994.html
func (t LBControlPanelACLGet) Document() string {
	return "http://manual.iij.jp/p2/pubapi/59940994.html"
}

// JPName LB管理画面アクセス制限情報取得
func (t LBControlPanelACLGet) JPName() string {
	return "LB管理画面アクセス制限情報取得"
}
func init() {
	APIlist = append(APIlist, LBControlPanelACLGet{})
	TypeMap["LBControlPanelACLGet"] = reflect.TypeOf(LBControlPanelACLGet{})
}

// LBControlPanelACLGetResponse LB管理画面アクセス制限情報取得のレスポンス
type LBControlPanelACLGetResponse struct {
	*CommonResponse
	AdministrationServerAllowNetworkList []struct {
	}
}

package protocol

import (
	"reflect"
)

// LBControlPanelACLSet LB管理画面アクセス制限設定 (同期)
//  http://manual.iij.jp/p2/pubapi/59941002.html
type LBControlPanelACLSet struct {
	GisServiceCode                       string `json:"-"` // P2契約のサービスコード(gis########)
	IflServiceCode                       string `json:"-"` // FW+LB 専有タイプのサービスコード(ifl########)
	AdministrationServerAllowNetworkList []struct {
	}
}

// URI /{{.GisServiceCode}}/fw-lbs/{{.IflServiceCode}}/lb/adminserver/allows.json
func (t LBControlPanelACLSet) URI() string {
	return "/{{.GisServiceCode}}/fw-lbs/{{.IflServiceCode}}/lb/adminserver/allows.json"
}

// APIName LBControlPanelACLSet
func (t LBControlPanelACLSet) APIName() string {
	return "LBControlPanelACLSet"
}

// Method PUT
func (t LBControlPanelACLSet) Method() string {
	return "PUT"
}

// Document http://manual.iij.jp/p2/pubapi/59941002.html
func (t LBControlPanelACLSet) Document() string {
	return "http://manual.iij.jp/p2/pubapi/59941002.html"
}

// JPName LB管理画面アクセス制限設定
func (t LBControlPanelACLSet) JPName() string {
	return "LB管理画面アクセス制限設定"
}
func init() {
	APIlist = append(APIlist, LBControlPanelACLSet{})
	TypeMap["LBControlPanelACLSet"] = reflect.TypeOf(LBControlPanelACLSet{})
}

// LBControlPanelACLSetResponse LB管理画面アクセス制限設定のレスポンス
type LBControlPanelACLSetResponse struct {
	*CommonResponse
}

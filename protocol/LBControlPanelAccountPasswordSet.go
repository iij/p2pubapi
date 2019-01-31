package protocol

import (
	"reflect"
)

// LBControlPanelAccountPasswordSet LB管理画面アカウントパスワード設定 (同期)
//  http://manual.iij.jp/p2/pubapi/59941013.html
type LBControlPanelAccountPasswordSet struct {
	GisServiceCode string `json:"-"` // P2契約のサービスコード(gis########)
	IflServiceCode string `json:"-"` // FW+LB 専有タイプのサービスコード(ifl########)
	AccountName    string `json:"-"` // パスワードを設定するアカウント("customer" "monitoring")
	Password       string // パスワード(8文字以上、32文字以下  半角英字・半角数字・半角記号をそれぞれ１文字以上含む  半角記号は「"#$%&'*+,.-/:;=<>()[]{}?@\^_`|~!」を利用可能)
}

// URI /{{.GisServiceCode}}/fw-lbs/{{.IflServiceCode}}/lb/accounts/{{.AccountName}}.json
func (t LBControlPanelAccountPasswordSet) URI() string {
	return "/{{.GisServiceCode}}/fw-lbs/{{.IflServiceCode}}/lb/accounts/{{.AccountName}}.json"
}

// APIName LBControlPanelAccountPasswordSet
func (t LBControlPanelAccountPasswordSet) APIName() string {
	return "LBControlPanelAccountPasswordSet"
}

// Method PUT
func (t LBControlPanelAccountPasswordSet) Method() string {
	return "PUT"
}

// Document http://manual.iij.jp/p2/pubapi/59941013.html
func (t LBControlPanelAccountPasswordSet) Document() string {
	return "http://manual.iij.jp/p2/pubapi/59941013.html"
}

// JPName LB管理画面アカウントパスワード設定
func (t LBControlPanelAccountPasswordSet) JPName() string {
	return "LB管理画面アカウントパスワード設定"
}
func init() {
	APIlist = append(APIlist, LBControlPanelAccountPasswordSet{})
	TypeMap["LBControlPanelAccountPasswordSet"] = reflect.TypeOf(LBControlPanelAccountPasswordSet{})
}

// LBControlPanelAccountPasswordSetResponse LB管理画面アカウントパスワード設定のレスポンス
type LBControlPanelAccountPasswordSetResponse struct {
	*CommonResponse
	なし string //
}

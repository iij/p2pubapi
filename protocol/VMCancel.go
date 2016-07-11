package protocol

import (
	"reflect"
)

// VMCancel 仮想サーバ解約申込 (同期)
//  http://manual.iij.jp/p2/pubapi/59939466.html
type VMCancel struct {
	GisServiceCode string `json:"-"`                // GisServiceCode
	IvmServiceCode string `json:"-"`                // IvmServiceCode
	StopDate       string `json:"-" p2pub:",query"` // 解約予定日。省略した場合は即時(YYYYMMDD。当日以降の日付を指定可能)
}

// URI /{{.GisServiceCode}}/virtual-servers/{{.IvmServiceCode}}.json
func (t VMCancel) URI() string {
	return "/{{.GisServiceCode}}/virtual-servers/{{.IvmServiceCode}}.json"
}

// APIName VMCancel
func (t VMCancel) APIName() string {
	return "VMCancel"
}

// Method DELETE
func (t VMCancel) Method() string {
	return "DELETE"
}

// Document http://manual.iij.jp/p2/pubapi/59939466.html
func (t VMCancel) Document() string {
	return "http://manual.iij.jp/p2/pubapi/59939466.html"
}

// JPName 仮想サーバ解約申込
func (t VMCancel) JPName() string {
	return "仮想サーバ解約申込"
}
func init() {
	APIlist = append(APIlist, VMCancel{})
	TypeMap["VMCancel"] = reflect.TypeOf(VMCancel{})
}

// VMCancelResponse 仮想サーバ解約申込のレスポンス
type VMCancelResponse struct {
	*CommonResponse
	Current struct {
		StopDate string `json:",omitempty"` // 設定した解約予定日(YYYYMMDD)
	} `json:",omitempty"`
	Previous struct {
		StopDate string `json:",omitempty"` // 設定前の解約予定日。未設定ならば空文字(YYYYMMDD)
	} `json:",omitempty"`
}

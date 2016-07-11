package protocol

import (
	"reflect"
)

// VMAdd 仮想サーバ追加申込 (同期)
//  http://manual.iij.jp/p2/pubapi/59939383.html
type VMAdd struct {
	GisServiceCode string `json:"-"` // P2契約のサービスコード(gis########)
	OSType         string // OS種別(Windows/Linux)
	ServerGroup    string `json:",omitempty"` // サーバグループ(A/B"。省略した場合はどちらかが自動的に選択される)
	Type           string // 仮想サーバ品目
}

// URI /{{.GisServiceCode}}/virtual-servers.json
func (t VMAdd) URI() string {
	return "/{{.GisServiceCode}}/virtual-servers.json"
}

// APIName VMAdd
func (t VMAdd) APIName() string {
	return "VMAdd"
}

// Method POST
func (t VMAdd) Method() string {
	return "POST"
}

// Document http://manual.iij.jp/p2/pubapi/59939383.html
func (t VMAdd) Document() string {
	return "http://manual.iij.jp/p2/pubapi/59939383.html"
}

// JPName 仮想サーバ追加申込
func (t VMAdd) JPName() string {
	return "仮想サーバ追加申込"
}
func init() {
	APIlist = append(APIlist, VMAdd{})
	TypeMap["VMAdd"] = reflect.TypeOf(VMAdd{})
}

// VMAddResponse 仮想サーバ追加申込のレスポンス
type VMAddResponse struct {
	*CommonResponse
	ResourceStatus string `json:",omitempty"` // 仮想サーバステータス
	ServiceCode    string `json:",omitempty"` // 仮想サーバのサービスコード(ivd########/ivm########)
	ServerSpec     struct {
		Memory string `json:",omitempty"` // メモリ量（GB）(例： 1GB)
		CPU    string `json:",omitempty"` // コア数(例： 1vCore)
	} `json:",omitempty"`
	StopDate string `json:",omitempty"` // 解約予定日(YYYYMMDD)
	Type     string `json:",omitempty"` // 仮想サーバ品目
	Network  struct {
		IncludingGlobalSide string `json:",omitempty"` // 帯域上限にグローバルネットワークを含む品目ならばYes, そうでなければNo(No/Yes)
		MaxTrafficSpeed     string `json:",omitempty"` // 帯域上限（Mbps）(例： 100Mbps)
	} `json:",omitempty"`
	NetworkList []struct {
		MacAddress    string `json:",omitempty"` // MACアドレス(例： 12-34-56-78-9a-bc)
		IpAddressList []struct {
			IPv4 struct {
				IpAddress string `json:",omitempty"` // IPv4アドレス
				Type      string `json:",omitempty"` // アドレス管理(Managed/Unmanaged)
			} `json:",omitempty"`
			IPv6 struct {
				IpAddress string `json:",omitempty"` // IPv6アドレス
				Type      string `json:",omitempty"` // アドレス管理(Managed/Unmanaged)
			} `json:",omitempty"`
		} `json:",omitempty"`
		Label       string `json:",omitempty"` // ラベル(文字列)
		ServiceCode string `json:",omitempty"` // プライベートネットワーク/Vのサービスコード(ivl########)
		NetworkType string `json:",omitempty"` // ネットワークタイプ(Private"。初期状態では"PrivateStandard"のみが接続されています/Gobal/PrivateStandard)
		URI         string `json:",omitempty"` // プライベートネットワーク/Vの情報へアクセスするためのURI(URI)
		IPv6Enabled string `json:",omitempty"` // IPv6有効または無効(Disabled/Enabled)
	} `json:",omitempty"`
	ServerGroup    string `json:",omitempty"` // サーバグループ(A/B)
	Label          string `json:",omitempty"` // ラベル(文字列)
	OSType         string `json:",omitempty"` // OS種別(Windows/Linux)
	StartDate      string `json:",omitempty"` // 利用開始日 or 課金開始日(YYYYMMDD)
	Category       string `json:",omitempty"` // 仮想サーバ種別(BestEffort/Gurantee/Dedicated)
	ContractStatus string `json:",omitempty"` // 契約状態
}

package protocol

import (
	"reflect"
)

// VMListGet 仮想サーバ情報一覧取得 (同期)
//  http://manual.iij.jp/p2/pubapi/59939407.html
type VMListGet struct {
	GisServiceCode string `json:"-"`                // GisServiceCode
	StartIndex     string `json:"-" p2pub:",query"` // StartIndexだけスキップした位置から情報を取得する(数値)
	Count          string `json:"-" p2pub:",query"` // 取得する情報の最大数(数値)
}

// URI /{{.GisServiceCode}}/virtual-servers.json
func (t VMListGet) URI() string {
	return "/{{.GisServiceCode}}/virtual-servers.json"
}

// APIName VMListGet
func (t VMListGet) APIName() string {
	return "VMListGet"
}

// Method GET
func (t VMListGet) Method() string {
	return "GET"
}

// Document http://manual.iij.jp/p2/pubapi/59939407.html
func (t VMListGet) Document() string {
	return "http://manual.iij.jp/p2/pubapi/59939407.html"
}

// JPName 仮想サーバ情報一覧取得
func (t VMListGet) JPName() string {
	return "仮想サーバ情報一覧取得"
}
func init() {
	APIlist = append(APIlist, VMListGet{})
	TypeMap["VMListGet"] = reflect.TypeOf(VMListGet{})
}

// VMListGetResponse 仮想サーバ情報一覧取得のレスポンス
type VMListGetResponse struct {
	*CommonResponse
	VirtualServerList []struct {
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
		StorageList []struct {
			Boot        string `json:",omitempty"` // ブート可否(No/Yes)
			PciSlot     string `json:",omitempty"` // PCIスロット(16進数。例： 0x10)
			ServiceCode string `json:",omitempty"` // ストレージのサービスコード(ibg########/iba########/ibb########)
			OSType      string `json:",omitempty"` // OS種別(None/Windows/Linux)
			URI         string `json:",omitempty"` // ストレージの情報へアクセスするためのURI(URI)
			Type        string `json:",omitempty"` // ストレージ品目
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
			NetworkType string `json:",omitempty"` // ネットワークタイプ(Gobal/PrivateStandard/Private)
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
}

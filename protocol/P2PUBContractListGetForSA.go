package protocol

import (
	"reflect"
)

// P2PUBContractListGetForSA P2（パブリックリソース）契約情報一覧取得（SA向け） (同期)
//  http://manual.iij.jp/p2/pubapi/58838523.html
type P2PUBContractListGetForSA struct {
	StartIndex string `json:"-" p2pub:",query"` // StartIndexだけスキップした位置から情報を取得する(数値)
	Count      string `json:"-" p2pub:",query"` // 取得する情報の最大数(数値)
}

// URI /gises.json
func (t P2PUBContractListGetForSA) URI() string {
	return "/gises.json"
}

// APIName P2PUBContractListGetForSA
func (t P2PUBContractListGetForSA) APIName() string {
	return "P2PUBContractListGetForSA"
}

// Method GET
func (t P2PUBContractListGetForSA) Method() string {
	return "GET"
}

// Document http://manual.iij.jp/p2/pubapi/58838523.html
func (t P2PUBContractListGetForSA) Document() string {
	return "http://manual.iij.jp/p2/pubapi/58838523.html"
}

// JPName P2（パブリックリソース）契約情報一覧取得（SA向け）
func (t P2PUBContractListGetForSA) JPName() string {
	return "P2（パブリックリソース）契約情報一覧取得（SA向け）"
}
func init() {
	APIlist = append(APIlist, P2PUBContractListGetForSA{})
	TypeMap["P2PUBContractListGetForSA"] = reflect.TypeOf(P2PUBContractListGetForSA{})
}

// P2PUBContractListGetForSAResponse P2（パブリックリソース）契約情報一覧取得（SA向け）のレスポンス
type P2PUBContractListGetForSAResponse struct {
	*CommonResponse
	GisList []struct {
		FwLbList []struct {
			URI         string `json:",omitempty"` // FW+LB 専有タイプの情報へアクセスするためのURI(URI)
			ServiceCode string `json:",omitempty"` // FW+LB 専有タイプのサービスコード(ifl########)
		} `json:",omitempty"`
		Region                string `json:",omitempty"` // リージョン(例：JP-EAST)
		AdditionalStorageList []struct {
			URI         string `json:",omitempty"` // 追加ストレージの情報へアクセスするためのURI(URI)
			Category    string `json:",omitempty"` // 追加ストレージ種別(BestEffort/Gurantee)
			ServiceCode string `json:",omitempty"` // 追加ストレージのサービスコード(ibg########/ibb########)
		} `json:",omitempty"`
		GlobalIpAddressNum struct {
			Assigned string `json:",omitempty"` // 利用中のグローバルアドレス数(数値)
			Limit    string `json:",omitempty"` // 利用可能なグローバルアドレスの上限数(数値)
		} `json:",omitempty"`
		ServiceCode            string `json:",omitempty"` // P2契約のサービスコード(gis########)
		PrivateNetworkStandard struct {
			NetworkAddress string `json:",omitempty"` // 標準プライベートネットワークのアドレスレンジ(IPv4アドレス/mask)
		} `json:",omitempty"`
		GlobalIpAddress struct {
			URI         string `json:",omitempty"` // グローバルIPアドレス/Vの情報へアクセスするためのURI(URI)
			ServiceCode string `json:",omitempty"` // グローバルIPアドレス/Vのサービスコード(iga########)
		} `json:",omitempty"`
		StopDate           string `json:",omitempty"` // 解約予定日(YYYYMMDD)
		Uom                string `json:",omitempty"` // IIJ GIO統合運用管理サービスのサービスコード。連携していなければ空文字(uom########)
		Site               string `json:",omitempty"` // サイト(例：S11)
		PrivateNetworkList []struct {
			URI         string `json:",omitempty"` // プライベートネットワーク/Vの情報へアクセスするためのURI(URI)
			ServiceCode string `json:",omitempty"` // プライベートネットワーク/Vのサービスコード(ivl########)
		} `json:",omitempty"`
		VirtualServerList []struct {
			URI         string `json:",omitempty"` // 仮想サーバの情報へアクセスするためのURI(URI)
			Category    string `json:",omitempty"` // 仮想サーバ種別(BestEffort/Gurantee/Dedicated)
			ServiceCode string `json:",omitempty"` // 仮想サーバのサービスコード(ivm########)
		} `json:",omitempty"`
		StorageArchive struct {
			URI         string `json:",omitempty"` // ストレージアーカイブの情報へアクセスするためのURI(URI)
			ServiceCode string `json:",omitempty"` // ストレージアーカイブのサービスコード(iar########)
		} `json:",omitempty"`
		Label             string `json:",omitempty"` // ラベル(文字列)
		SystemStorageList []struct {
			URI         string `json:",omitempty"` // システムストレージの情報へアクセスするためのURI(URI)
			ServiceCode string `json:",omitempty"` // システムストレージのサービスコード(iba########)
		} `json:",omitempty"`
		StartDate      string `json:",omitempty"` // 利用開始日(YYYYMMDD)
		ContractStatus string `json:",omitempty"` // 契約状態(詳細は別ページへ)
	}
}

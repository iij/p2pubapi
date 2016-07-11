package protocol

import (
	"reflect"
)

// P2PUBContractListGetForSGSA P2（パブリックリソース）契約情報一覧取得（SGSA向け） (同期)
//  http://manual.iij.jp/p2/pubapi/58449151.html
type P2PUBContractListGetForSGSA struct {
	CustomerCode string `json:"-"`                // カスタマーコード(SG########)
	StartIndex   string `json:"-" p2pub:",query"` // StartIndexだけスキップした位置から情報を取得する。省略した場合は先頭から取得(数字)
	Count        string `json:"-" p2pub:",query"` // 取得する情報の最大数。省略した場合は終わりまですべて取得(数字)
}

// URI /{{.CustomerCode}}/gises.json
func (t P2PUBContractListGetForSGSA) URI() string {
	return "/{{.CustomerCode}}/gises.json"
}

// APIName P2PUBContractListGetForSGSA
func (t P2PUBContractListGetForSGSA) APIName() string {
	return "P2PUBContractListGetForSGSA"
}

// Method GET
func (t P2PUBContractListGetForSGSA) Method() string {
	return "GET"
}

// Document http://manual.iij.jp/p2/pubapi/58449151.html
func (t P2PUBContractListGetForSGSA) Document() string {
	return "http://manual.iij.jp/p2/pubapi/58449151.html"
}

// JPName P2（パブリックリソース）契約情報一覧取得（SGSA向け）
func (t P2PUBContractListGetForSGSA) JPName() string {
	return "P2（パブリックリソース）契約情報一覧取得（SGSA向け）"
}
func init() {
	APIlist = append(APIlist, P2PUBContractListGetForSGSA{})
	TypeMap["P2PUBContractListGetForSGSA"] = reflect.TypeOf(P2PUBContractListGetForSGSA{})
}

// P2PUBContractListGetForSGSAResponse P2（パブリックリソース）契約情報一覧取得（SGSA向け）のレスポンス
type P2PUBContractListGetForSGSAResponse struct {
	*CommonResponse
	GisList []struct {
		FwLbList []struct {
			URI         string `json:",omitempty"` // FW+LB 専有タイプの情報へアクセスするためのURI(URI)
			ServiceCode string `json:",omitempty"` // FW+LB 専有タイプのサービスコード(ifl########)
		} `json:",omitempty"`
		Region                string `json:",omitempty"` // リージョン(例：JP-EAST)
		AdditionalStorageList []struct {
			URI         string `json:",omitempty"` // 追加ストレージの情報へアクセスするためのURI(URI)
			Category    string `json:",omitempty"` // 追加ストレージ種別(BestEffort/Gurantee"のいずれか)
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
			Category    string `json:",omitempty"` // 仮想サーバ種別(BestEffort/Gurantee/Dedicated"のいずれか)
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
		ContractStatus string `json:",omitempty"` // 契約状態(詳細は＜契約状態＞へ)
	}
}

package protocol

import (
	"reflect"
)

// StorageGet 追加ストレージ情報取得 (同期)
//  http://manual.iij.jp/p2/pubapi/59940143.html
type StorageGet struct {
	GisServiceCode string `json:"-"` // P2契約のサービスコード(gis########)
	IbgServiceCode string `json:"-"` // 追加ストレージのサービスコード(ibb########, ibg########)
}

// URI /{{.GisServiceCode}}/additional-storages/{{.IbgServiceCode}}.json
func (t StorageGet) URI() string {
	return "/{{.GisServiceCode}}/additional-storages/{{.IbgServiceCode}}.json"
}

// APIName StorageGet
func (t StorageGet) APIName() string {
	return "StorageGet"
}

// Method GET
func (t StorageGet) Method() string {
	return "GET"
}

// Document http://manual.iij.jp/p2/pubapi/59940143.html
func (t StorageGet) Document() string {
	return "http://manual.iij.jp/p2/pubapi/59940143.html"
}

// JPName 追加ストレージ情報取得
func (t StorageGet) JPName() string {
	return "追加ストレージ情報取得"
}
func init() {
	APIlist = append(APIlist, StorageGet{})
	TypeMap["StorageGet"] = reflect.TypeOf(StorageGet{})
}

// StorageGetResponse 追加ストレージ情報取得のレスポンス
type StorageGetResponse struct {
	*CommonResponse
	ResourceStatus        string `json:",omitempty"` // ストレージステータス
	AttachedVirtualServer struct {
		OSType      string `json:",omitempty"` // OS種別(Windows/Linux)
		URI         string `json:",omitempty"` // 仮想サーバの情報へアクセスするためのURI(URI)
		PciSlot     string `json:",omitempty"` // PCIスロット(16進数)
		ServiceCode string `json:",omitempty"` // 接続された仮想サーバのサービスコード(ivd########/ivm########)
		Type        string `json:",omitempty"` // 仮想サーバ品目
	} `json:",omitempty"`
	Label          string `json:",omitempty"` // ラベル(文字列)
	OSType         string `json:",omitempty"` // OS種別(None)
	StorageGroup   string `json:",omitempty"` // ストレージグループ(Z/Y)
	StartDate      string `json:",omitempty"` // 利用開始日(YYYYMMDD)
	StorageSize    string `json:",omitempty"` // ストレージ容量(数値)
	Category       string `json:",omitempty"` // 追加ストレージ種別(BestEffort/Gurantee)
	ContractStatus string `json:",omitempty"` // 契約状態
	StopDate       string `json:",omitempty"` // 解約予定日(YYYYMMDD。未設定の場合は空文字列)
	Type           string `json:",omitempty"` // 追加ストレージ品目
}

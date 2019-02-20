package protocol

import (
	"reflect"
)

// StorageListGet 追加ストレージ情報一覧取得 (同期)
//  http://manual.iij.jp/p2/pubapi/59940121.html
type StorageListGet struct {
	GisServiceCode string `json:"-"`                // P2契約のサービスコード(gis########)
	StartIndex     string `json:"-" p2pub:",query"` // StartIndexだけスキップした位置から情報を取得する(数値)
	Count          string `json:"-" p2pub:",query"` // 取得する情報の最大数(数値)
}

// URI /{{.GisServiceCode}}/additional-storages.json
func (t StorageListGet) URI() string {
	return "/{{.GisServiceCode}}/additional-storages.json"
}

// APIName StorageListGet
func (t StorageListGet) APIName() string {
	return "StorageListGet"
}

// Method GET
func (t StorageListGet) Method() string {
	return "GET"
}

// Document http://manual.iij.jp/p2/pubapi/59940121.html
func (t StorageListGet) Document() string {
	return "http://manual.iij.jp/p2/pubapi/59940121.html"
}

// JPName 追加ストレージ情報一覧取得
func (t StorageListGet) JPName() string {
	return "追加ストレージ情報一覧取得"
}
func init() {
	APIlist = append(APIlist, StorageListGet{})
	TypeMap["StorageListGet"] = reflect.TypeOf(StorageListGet{})
}

// StorageListGetResponse 追加ストレージ情報一覧取得のレスポンス
type StorageListGetResponse struct {
	*CommonResponse
	AdditionalStorageList []struct {
		ResourceStatus        string `json:",omitempty"` // ストレージステータス
		StorageGroup          string `json:",omitempty"` // ストレージグループ(Z/Y)
		ServiceCode           string `json:",omitempty"` // 追加ストレージのサービスコード(ibg########/ibb########/icg########/icb########)
		StorageSize           string `json:",omitempty"` // ストレージ容量(数値)
		StopDate              string `json:",omitempty"` // 解約予定日(YYYYMMDD。未設定ならば空文字列)
		Type                  string `json:",omitempty"` // 追加ストレージ品目
		Label                 string `json:",omitempty"` // ラベル(文字列)
		OSType                string `json:",omitempty"` // OS種別(None)
		StartDate             string `json:",omitempty"` // 利用開始日(YYYYMMDD)
		AttachedVirtualServer struct {
			OSType      string `json:",omitempty"` // OS種別(Windows/Linux)
			URI         string `json:",omitempty"` // 仮想サーバの情報へアクセスするためのURI(URI)
			PciSlot     string `json:",omitempty"` // PCIスロット(16進数)
			ServiceCode string `json:",omitempty"` // 接続された仮想サーバのサービスコード(ivd########/ivm########)
			Type        string `json:",omitempty"` // 仮想サーバ品目
		} `json:",omitempty"`
		Category       string `json:",omitempty"` // 追加ストレージ種別(BestEffort/Gurantee)
		ContractStatus string `json:",omitempty"` // 契約状態
		Mode           string `json:",omitempty"` // ストレージのモード ("Basic", "Backup")
		Encryption     string `json:",omitempty"` // 暗号化 ("Yes", "No")
	}
}

package jmaseis

// 内容部
type Body struct {
	Earthquake   []Earthquake // 地震の諸要素 (任意項目、震源に関する情報、震源・震度に関する情報で 0-1 回。 Head/InfoType = "取消" では出現しない)
	Intensity    Intensity    // 震度 / 震度・長周期地震動階級 (任意項目、震度速報、震源・震度に関する情報、緊急地震速報（警報）で 0-1 回。 Head/InfoType = "取消" では出現しない)
	Text         string       // テキスト要素 (任意項目。例えば Head/InfoType = "取消" での取消概要等)
	Comments     []Comment    // 付加文 (任意項目。地震情報では 0-1 回。 Head/InfoType = "取消" では出現しない)
	Tsunami      Tsunami      // 津波 (任意項目。 Head/InfoType = "取消" では出現しない)
	NextAdvisory string       // 次回発表予定 (任意項目。緊急地震速報（警報警報）で 0-1 回。 "この情報をもって、緊急地震速報：最終報とします。")
	// Tsunami
	// Naming
	// Tokai
	// EarthquakeInfo
	// EarthquakeCount
	// Aftershock
}

// 地震の諸要素
type Earthquake struct {
	OriginTime  DateTime    // 地震発生時刻 (任意項目、地震情報では必須)
	ArrivalTime DateTime    // 地震発現時刻 遠地地震で発現時刻不明な場合、地震発生時刻と同値 (任意項目、地震情報では必須)
	Hypocenter  Hypocenter  // 地震の位置要素 (任意項目、地震情報では必須)
	Magnitude   []Magnitude // マグニチュード (任意項目、震源または震源・震度に関する情報では要素数 1)
	Condition   string      // 震源要素の補足情報 (任意項目、緊急地震速報（警報）で 0-1 回。 "仮定震源要素)）
}

// 地震の位置要素
type Hypocenter struct {
	Area     HypoArea // 震源位置
	Source   string   // 震源決定機関 (任意項目、遠地地震で気象庁以外の機関で決定された震源要素を採用した場合の機関略称。 "PTWC" / "WCATWC" / "USGS")
	Accuracy Accuracy // 精度情報 (任意項目、緊急地震速報（警報）では必須)
}

// 震源位置
type HypoArea struct {
	Name         string       // 震央地名
	Code         HypoAreaCode // 震央地名に対応するコード
	Coordinate   []Coordinate // 震源要素 (地震情報では要素数 1)
	DetailedName string       // 詳細震央地名 (任意項目、遠地地震において詳細な位置を発表する場合のみ)
	ReduceName   string       // 短縮用震央地名 (任意項目、緊急地震速報（警報）では必須)
	ReduceCode   string       // 短縮用震央地名コード (任意項目、緊急地震速報（警報）では必須)
	// DetailedCode
	// NameFromMark
	// MarkCode
	// Direction
	// Distance
	// LandOrSea
}

// 精度情報
type Accuracy struct {
	Epicenter            Epicenter            // 震央位置の精度値
	Depth                Depth                // 深さの精度値
	MagnitudeCalculation MagnitudeCalculation // マグニチュードの精度値
}

// 震央位置の精度値
type Epicenter struct {
	Rank  string `xml:"rank,attr"` // 震源位置の精度のランク
	Rank2 string `xml:"rank2,attr"` // 震源位置の精度のランク2
}

// 深さの精度値
type Depth struct {
	Rank string `xml:"rank,attr"` // 震源深さの精度ランク
}

// マグニチュードの精度値
type MagnitudeCalculation struct {
	Rank string `xml:"rank,attr"` // マグニチュード精度のランク
}

// 震央地名に対応するコード
type HypoAreaCode struct {
	Type string `xml:"type,attr"` // "震央地名"
	Name string `xml:",chardata"`
}

/*
	震源要素
		- 全要素が不明な場合 Description = "震源要素不明" のみ
		- 震源の深さが 0-5 km の場合 Description に "ごく浅い" 、 Value は 0 メートル
		- 震源の深さが不明の場合 Description に "深さ不明" 、 Value は深さ表現なし
*/
type Coordinate struct {
	Datum       string `xml:"datum,attr"`       // "日本測地系" (任意項目、遠地地震の震源要素は世界測地系に基づくため出現しない)
	Description string `xml:"description,attr"` // 文字列表現 (任意項目)
	Value       string `xml:",chardata"`        // ISO6709 での表現 (緯度・経度は度単位、深さはメートル単位。深さは 700 km より浅いところでは 10km 単位) (任意項目)
	// Type        string `xml:"type,attr"`        // 任意項目
	// Condition   string `xml:"condition,attr"`   // 任意項目
}

// マグニチュード
type Magnitude struct {
	Type        string         `xml:"type,attr"`        // マグニチュードの種別 ("Mj" など)
	Condition   string         `xml:"condition,attr"`   // 不明な場合や 8 を超える巨大地震と推定される場合 "不明" (任意項目)
	Description string         `xml:"description,attr"` // 文字列表現 (任意項目)
	Value       MagnitudeValue `xml:",chardata"`        // マグニチュード。不明または 8 を超える巨大地震と推定される場合 "NaN"
}

// 震度 / 震度・長周期地震動階級（緊急地震速報（警報）の場合）
type Intensity struct {
	Observation IntensityDetail // 震度の観測 (任意項目、地震情報では必須)
	Forecast    Forecast        // 震度・長周期地震動の予測 (任意項目、緊急地震速報（警報）では必須)
}

// 震度の観測
type IntensityDetail struct {
	CodeDefine CodeDefine      // コード体系の定義 (任意項目、地震情報では必須)
	MaxInt     string          // 最大震度 (任意項目、地震情報では必須。 "1" / "2" / "3" / "4" / "5-" / "5+" / "6-" / "6+" / "7")
	Pref       []IntensityPref // 都道府県 (任意項目、地震情報では必須)
	// ForecastInt
	// Appendix
}

type CodeDefine struct {
	Type []CodeDefineType
}

type CodeDefineType struct {
	XPath string `xml:"xpath,attr"`
	Value string `xml:",chardata"`
}

// 震度・長周期地震動階級の予測
type Forecast struct {
	CodeDefine    CodeDefine      // コード体系の定義
	ForecastInt   ForecastInt     // 最大予測震度
	ForecastLgInt ForecastLgInt   // 最大予測長周期地震動階級 (任意項目)
	Appendix      Appendix        // 予測震度・予測長周期地震動階級付加要素 (任意項目)
	Pref          []IntensityPref // 都道府県要素
}

type ForecastInt struct {
	From string // 最大予測震度の下限
	To   string // 最大予測震度の上限
}

type ForecastLgInt struct {
	From string // 最大予測長周期地震動階級の下限
	To   string // 最大予測長周期地震動階級の上限
}

type Appendix struct {
	MaxIntChange       string // 最大予測震度変化
	MaxLgIntChange     string // 最大予測長周期地震動階級変化 (任意項目)
	MaxIntChangeReason string // 最大予測値変化の理由
}

// 都道府県
type IntensityPref struct {
	Name   string          // 都道府県名
	Code   string          // 都道府県名に対応するコード
	MaxInt string          // 最大震度（都道府県） (任意項目。震源・震度に関する情報では、都道府県内に震度 5 弱以上と考えられるが震度の値を入手していない市町村のみしか存在しない場合は出現しない)
	Area   []IntensityArea // 地域 (任意項目。地震情報、緊急地震速報（警報）では必須)
	Revise string          // 情報の更新（都道府県） (任意項目。震源・震度に関する情報の続報について、都道府県が "追加" または震度が "上方修正" されることがある)
	// Category
	// ForecastInt
	// ArrivalTime
	// Condition
}

// 地域
type IntensityArea struct {
	Name          string
	Code          string
	MaxInt        string           // 最大震度（地域） (任意項目。震源・震度に関する情報において、地域内に震度 5 弱以上と考えられるが震度の値を入手していない市町村のみしか存在しない場合は出現しない)
	City          []IntensityCity  // 市町村 (任意項目。震度速報では存在しない。震源・震度に関する情報では必須)
	Revise        string           // 情報の更新（地域） (任意項目。震源・震度に関する情報の続報について、地域が "追加" または震度が "上方修正" されることがある)
	Category      ForecastCategory // 予報カテゴリー (任意項目、緊急地震速報（警報）のみ)
	ForecastInt   ForecastInt      // 最大予測震度 (任意項目、緊急地震速報（警報）のみ)
	ForecastLgInt ForecastLgInt    // 最大予測長周期地震動階級 (任意項目、緊急地震速報（警報）で 0-1 回)
	ArrivalTime   DateTime         // 主要動の到達予測時刻 (任意項目、緊急地震速報（警報）で 0-1 回)
	Condition     string           // 状況 (任意項目。緊急地震速報（警報）で 0-1 回。 "既に主要動到達と推測")
}

// 市町村
type IntensityCity struct {
	Name             string
	Code             string
	MaxInt           string             // 最大震度（市町村） (任意項目。震源・震度に関する情報では、市町村内に震度 5 弱以上と考えられるが震度の値を入手していない市町村のみしか存在しない場合は出現しない)
	IntensityStation []IntensityStation // 震度観測点 (任意項目、震源・震度に関する情報では必須)
	Revise           string             // 情報の更新（市町村） (任意項目。震源・震度に関する情報の続報について、市町村が "追加" または震度が "上方修正" されることがある)
	Condition        string             // "震度５弱以上未入電" (任意項目。当該市町村内において、入電した情報が最大震度 4 未満の場合のみ)
	// Category
	// ForecastInt
	// ArrivalTime
}

// 震度観測点
type IntensityStation struct {
	Name   string
	Code   string
	Int    string // "1" / "2" / "3" / "4" / "5-" / "5+" / "6-" / "6+" / "7" / "震度５弱以上未入電"
	Revise string // NULL / "追加" / "上方修正" / "下方修正"
	// K
}

// 予報カテゴリー
type ForecastCategory struct {
	Kind ForecastKind
}

// 今回予報
type ForecastKind struct {
	Name string // 警報名
	Code string // 警報コード
}

// 付加文
type Comment struct {
	ForecastComment CommentForm // 固定付加文 (任意項目。津波や緊急地震速報に関する付加的な情報)
	VarComment      CommentForm // 固定付加文（その他） (任意項目。その他の付加的な情報)
	FreeFormComment string      // 自由付加文 (任意項目)
	WarningComment  CommentForm // 固定付加文 (任意項目。緊急地震速報（警報）のみ)
	// ObservationComment
}

// 固定付加文
type CommentForm struct {
	CodeType string `xml:"codeType,attr"` // "固定付加文"
	Text     string // 複数の固定付加文を記載する場合、改行して併記
	Code     string // 複数の固定付加文を記載する場合、 xs:list (空白区切り) として併記
}

// 津波
type Tsunami struct {
	Release     string        // (任意項目)
	Observation TsunamiDetail // (任意項目)
	Estimation  TsunamiDetail // (任意項目)
	Forecast    TsunamiDetail // 津波の予測値 (任意項目。津波警報・注意報・予報 では必須)
}

// 津波警報・注意報・予報の場合、津波の予測値
type TsunamiDetail struct {
	CodeDefine CodeDefine    // (任意項目)
	Item       []TsunamiItem // 津波警報・注意報・予報の場合、津波の予測値（津波予報区毎）
}

// 津波警報・注意報・予報の場合、津波の予測値（津波予報区毎）
type TsunamiItem struct {
	Area        ForecastArea     // 津波予報区
	Category    Category         // 津波警報等の種類 (任意項目。津波警報・注意報・予報では必須)
	FirstHeight FirstHeight      // 津波の到達予想時刻 (任意項目。津波警報・注意報を解除または津波予報（若干の海面変動）を発表している予報区では出現しない)
	MaxHeight   MaxHeight        // 予想される津波の高さ (任意項目)
	Duration    Duration         // (任意項目)
	Station     []TsunamiStation // (任意項目)
}

type ForecastArea struct {
	Name string
	Code string
	City []ForecastCity // (任意項目)
}

type ForecastCity struct {
	Name string
	Code string
}

// 津波警報等の種類
type Category struct {
	Kind     Kind // 津波警報等の発表状況。大津波警報については、第 1 報を含めて新たに大津波警報となる津波予報区で "大津波警報：発表" 、継続で "大津波警報" と記載
	LastKind Kind // 一つ前の情報による発表状況 (任意項目)
}

// 発表状況
type Kind struct {
	Name string
	Code string
}

// 津波の到達予想時刻
type FirstHeight struct {
	ArrivalTime   DateTime      // 第 1 波の到達予想時刻 (任意項目。第 1 波が到達または到達と推測される場合は出現しない)
	Condition     string        // NULL / "ただちに津波来襲と予測" / "津波到達中と推測" / "第１波の到達を確認" (任意項目)
	Initial       string        // (任意項目)
	TsunamiHeight TsunamiHeight // (任意項目)
	Revise        string        // NULL / "追加"/ "更新" (任意項目)
	Period        float64       // (任意項目)
}

// 予想される津波の高さ
type MaxHeight struct {
	DateTime      DateTime      // (任意項目)
	Condition     string        // 大津波警報の予想高さが最初に発表された場合や上方修正された場合 "重要" (任意項目)
	TsunamiHeight TsunamiHeight // 予想される津波の高さ(メートル単位) (任意項目)
	Revise        string        // NULL / "追加" / "更新" (任意項目)
	Period        float64       // (任意項目)
}

type CurrentHeight struct {
	StartTime     DateTime      // (任意項目)
	EndTime       DateTime      // (任意項目)
	Condition     string        // (任意項目)
	TsunamiHeight TsunamiHeight // (任意項目)
}

type TsunamiStation struct {
	Name             string
	Code             string
	Sensor           string   // (任意項目)
	HighTideDateTime DateTime // (任意項目)
	FirstHeight      FirstHeight
	MaxHeight        MaxHeight     // (任意項目)
	CurrentHeight    CurrentHeight // (任意項目)
}

type TsunamiHeight struct {
	Type        string  `xml:"type,attr"`        // "津波の高さ"
	Unit        string  `xml:"unit,attr"`        // "m"
	Condition   string  `xml:"condition,attr"`   // 地震規模推定の不確実性が大きい場合 "不明" (任意項目)
	Description string  `xml:"description,attr"` // 文字列表現 "巨大" / "高い" / "１０ｍ超" / "１０ｍ" / "５ｍ" / "３ｍ" / "１ｍ" / "０．２ｍ未満" (任意項目)
	Value       float64 `xml:",chardata"`        // メートル単位の値。定性的表現の場合 "NaN"
}

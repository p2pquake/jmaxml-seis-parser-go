package vxse

// 内容部
type Body struct {
	Earthquake []Earthquake // 地震の諸要素 (任意項目、震源に関する情報、震源・震度に関する情報で 0-1 回。 Head/InfoType = "取消" では出現しない)
	Intensity  Intensity    // 震度 (任意項目、震度速報、震源・震度に関する情報で 0-1 回。 Head/InfoType = "取消" では出現しない)
	Text       string       // テキスト要素 (任意項目。例えば Head/InfoType = "取消" での取消概要等)
	Comments   []Comment    // 付加文 (任意項目。地震情報では 0-1 回。 Head/InfoType = "取消" では出現しない)
	// Tsunami
	// Naming
	// Tokai
	// EarthquakeInfo
	// EarthquakeCount
	// Aftershock
	// NextAdvisory
}

// 地震の諸要素
type Earthquake struct {
	OriginTime  DateTime    // 地震発生時刻 (任意項目、地震情報では必須)
	ArrivalTime DateTime    // 地震発現時刻 遠地地震で発現時刻不明な場合、地震発生時刻と同値 (任意項目、地震情報では必須)
	Hypocenter  Hypocenter  // 地震の位置要素 (任意項目、地震情報では必須)
	Magnitude   []Magnitude // マグニチュード (任意項目、震源または震源・震度に関する情報では要素数 1)
	// Condition   string     // 任意項目
}

// 地震の位置要素
type Hypocenter struct {
	Area   HypoArea // 震源位置
	Source string   // 震源決定機関 (任意項目、遠地地震で気象庁以外の機関で決定された震源要素を採用した場合の機関略称。 "PTWC" / "WCATWC" / "USGS")
	// Accuracy
}

// 震源位置
type HypoArea struct {
	Name         string       // 震央地名
	Code         HypoAreaCode // 震央地名に対応するコード
	Coordinate   []Coordinate // 震源要素 (地震情報では要素数 1)
	DetailedName string       // 詳細震央地名 (任意項目、遠地地震において詳細な位置を発表する場合のみ)
	// ReduceName
	// ReduceCode
	// DetailedCode
	// NameFromMark
	// MarkCode
	// Direction
	// Distance
	// LandOrSea
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
	Type        string  `xml:"type,attr"`        // マグニチュードの種別 ("Mj" など)
	Condition   string  `xml:"condition,attr"`   // 不明な場合や 8 を超える巨大地震と推定される場合 "不明" (任意項目)
	Description string  `xml:"description,attr"` // 文字列表現 (任意項目)
	Value       float64 `xml:",chardata"`        // マグニチュード。不明または 8 を超える巨大地震と推定される場合 "NaN"
}

// 震度
type Intensity struct {
	Observation IntensityDetail // 震度の観測 (任意項目、地震情報では必須)
	// Forecast
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

// 都道府県
type IntensityPref struct {
	Name   string          // 都道府県名
	Code   string          // 都道府県名に対応するコード
	MaxInt string          // 最大震度（都道府県） (任意項目。震源・震度に関する情報では、都道府県内に震度 5 弱以上と考えられるが震度の値を入手していない市町村のみしか存在しない場合は出現しない)
	Area   []IntensityArea // 地域 (任意項目、地震情報では必須)
	Revise string          // 情報の更新（都道府県） (任意項目。震源・震度に関する情報の続報について、都道府県が "追加" または震度が "上方修正" されることがある)
	// Category
	// ForecastInt
	// ArrivalTime
	// Condition
}

// 地域
type IntensityArea struct {
	Name   string
	Code   string
	MaxInt string          // 最大震度（地域） (任意項目。震源・震度に関する情報において、地域内に震度 5 弱以上と考えられるが震度の値を入手していない市町村のみしか存在しない場合は出現しない)
	City   []IntensityCity // 市町村 (任意項目。震度速報では存在しない。震源・震度に関する情報では必須)
	Revise string          // 情報の更新（地域） (任意項目。震源・震度に関する情報の続報について、地域が "追加" または震度が "上方修正" されることがある)
	// Category
	// ForecastInt
	// ArrivalTime
	// Condition
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

// 付加文
type Comment struct {
	ForecastComment CommentForm // 固定付加文 (任意項目。津波や緊急地震速報に関する付加的な情報)
	VarComment      CommentForm // 固定付加文（その他） (任意項目。その他の付加的な情報)
	FreeFormComment string      // 自由付加文 (任意項目)
	// WarningComment
	// ObservationComment
}

// 固定付加文
type CommentForm struct {
	CodeType string `xml:"codeType,attr"` // "固定付加文"
	Text     string // 複数の固定付加文を記載する場合、改行して併記
	Code     string // 複数の固定付加文を記載する場合、 xs:list (空白区切り) として併記
}

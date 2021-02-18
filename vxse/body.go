package vxse

// 内容部
type Body struct {
	Earthquake []Earthquake // 任意項目
	Intensity  Intensity    // 任意項目
	Comments   []Comment    // 任意項目
	// Tsunami
	// Naming
	// Tokai
	// EarthquakeInfo
	// EarthquakeCount
	// Aftershock
	// Text
	// NextAdvisory
}

// 地震の諸要素
type Earthquake struct {
	OriginTime  DateTime // 任意項目
	ArrivalTime DateTime
	Condition   string     // 任意項目
	Hypocenter  Hypocenter // 任意項目
	Magnitude   []Magnitude
}

type Hypocenter struct {
	Area   HypoArea
	Source string // 任意項目
	// Accuracy
}

type HypoArea struct {
	Name       string
	Code       HypoAreaCode
	Coordinate []Coordinate
	// ReduceName
	// ReduceCode
	// DetailedName
	// DetailedCode
	// NameFromMark
	// MarkCode
	// Direction
	// Distance
	// LandOrSea
}

type HypoAreaCode struct {
	Type string `xml:"type,attr"`
	Name string `xml:",chardata"`
}

type Coordinate struct {
	Type        string `xml:"type,attr"`        // 任意項目
	Datum       string `xml:"datum,attr"`       // 任意項目 "日本測地系"
	Condition   string `xml:"condition,attr"`   // 任意項目
	Description string `xml:"description,attr"` // 任意項目
	Value       string `xml:",chardata"`        // 任意項目
}

type Magnitude struct {
	Type        string  `xml:"type,attr"`
	Condition   string  `xml:"condition,attr"`   // 任意項目
	Description string  `xml:"description,attr"` // 任意項目
	Value       float64 `xml:",chardata"`
}

type Intensity struct {
	Observation IntensityDetail // 任意項目
	// Forecast
}

type IntensityDetail struct {
	CodeDefine CodeDefine      // 任意項目
	MaxInt     string          // 任意項目
	Pref       []IntensityPref // 任意項目
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

type IntensityPref struct {
	Name   string
	Code   string
	MaxInt string          // 任意項目
	Area   []IntensityArea // 任意項目
	Revise string          // 任意項目
	// Category
	// ForecastInt
	// ArrivalTime
	// Condition
}

type IntensityArea struct {
	Name   string
	Code   string
	MaxInt string          // 任意項目
	City   []IntensityCity // 任意項目
	Revise string          // 任意項目
	// Category
	// ForecastInt
	// ArrivalTime
	// Condition
}

type IntensityCity struct {
	Name             string
	Code             string
	MaxInt           string             // 任意項目
	IntensityStation []IntensityStation // 任意項目
	Revise           string             // 任意項目
	// Category
	// ForecastInt
	// ArrivalTime
	// Condition
}

type IntensityStation struct {
	Name   string
	Code   string
	Int    string
	Revise string // 任意項目
	// K
}

type Comment struct {
	ForecastComment CommentForm // 任意項目
	VarComment      CommentForm // 任意項目
	FreeFormComment string      // 任意項目
	// WarningComment
	// ObservationComment
}

type CommentForm struct {
	CodeType string `xml:"codeType,attr"`
	Text     string
	Code     []string
}

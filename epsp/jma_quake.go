package epsp

// 地震情報
type JMAQuake struct {
	Code       int // "551"
	Issue      Issue
	Earthquake Earthquake
	Points     []Point
}

type Issue struct {
	Source  string
	Time    string
	Type    string // "ScalePrompt" / "Destination" / "ScaleAndDestination" / "DetailScale" / "Foreign" / "Other"
	Correct string // "None" / "Unknown" / "ScaleOnly" / "DestinationOnly" / "ScaleAndDestination"
}

type Earthquake struct {
	Time            string
	Hypocenter      Hypocenter
	MaxScale        int
	DomesticTsunami string
	ForeignTsunami  string
}

type Hypocenter struct {
	Name      string
	Latitude  float64
	Longitude float64
	Depth     int
	Magnitude float64
}

type Point struct {
	Pref   string
	Addr   string
	IsArea bool
	Scale  int
}

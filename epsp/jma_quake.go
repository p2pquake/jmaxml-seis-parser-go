package epsp

// 地震情報
type JMAQuake struct {
	// Code       int        `json:"code"` // "551"
	Expire     *string    `json:"expire"`
	Issue      Issue      `json:"issue"`
	Earthquake Earthquake `json:"earthquake"`
	Points     []Point    `json:"points"`
	Comments   Comments   `json:"comments"`
}

type Issue struct {
	Source  string `json:"source"`
	Time    string `json:"time"`
	Type    string `json:"type"`    // "ScalePrompt" / "Destination" / "ScaleAndDestination" / "DetailScale" / "Foreign" / "Other"
	Correct string `json:"correct"` // "None" / "Unknown" / "ScaleOnly" / "DestinationOnly" / "ScaleAndDestination"
}

type Earthquake struct {
	Time            string     `json:"time"`
	Hypocenter      Hypocenter `json:"hypocenter"`
	MaxScale        int        `json:"maxScale"`
	DomesticTsunami string     `json:"domesticTsunami"`
	ForeignTsunami  string     `json:"foreignTsunami"`
}

type Hypocenter struct {
	Name      string  `json:"name"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	Depth     int     `json:"depth"`
	Magnitude float64 `json:"magnitude"`
}

type Point struct {
	Pref   string `json:"pref"`
	Addr   string `json:"addr"`
	Scale  int    `json:"scale"`
	IsArea bool   `json:"isArea"`
}

type Comments struct {
	FreeFormComment string `json:"freeFormComment"`
}

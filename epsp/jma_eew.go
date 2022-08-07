package epsp

// 緊急地震速報（警報）
type JMAEEW struct {
	Earthquake *EEWEarthquake `json:"earthquake"`
	Issue      *EEWIssue      `json:"issue"`
	Cancelled  bool           `json:"cancelled"`
	Areas      []EEWArea      `json:"areas"`
}

type EEWIssue struct {
	Time    string `json:"time"`
	EventID string `json:"eventId"`
	Serial  string `json:"serial"`
}

type EEWEarthquake struct {
	OriginTime  string        `json:"originTime"`
	ArrivalTime string        `json:"arrivalTime"`
	Condition   string        `json:"condition"`
	Hypocenter  EEWHypocenter `json:"hypocenter"`
}

type EEWHypocenter struct {
	Name       string  `json:"name"`
	ReduceName string  `json:"reduceName"`
	Latitude   float64 `json:"latitude"`
	Longitude  float64 `json:"longitude"`
	Depth      int     `json:"depth"`
	Magnitude  float64 `json:"magnitude"`
}

type EEWArea struct {
	Pref        string  `json:"pref"`
	Name        string  `json:"name"`
	ScaleFrom   int     `json:"scaleFrom"`
	ScaleTo     int     `json:"scaleTo"`
	KindCode    string  `json:"kindCode"`
	ArrivalTime *string `json:"arrivalTime"`
}

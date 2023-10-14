package epsp

// 津波予報
type JMATsunami struct {
	Expire    *string      `json:"expire"`
	Issue     TsunamiIssue `json:"issue"`
	Cancelled bool         `json:"cancelled"`
	Areas     []Area       `json:"areas"`
}

type TsunamiIssue struct {
	Source string `json:"source"`
	Time   string `json:"time"`
	Type   string `json:"type"`
}

type Area struct {
	Name        string      `json:"name"`
	Grade       string      `json:"grade"`
	Immediate   bool        `json:"immediate"`
	FirstHeight FirstHeight `json:"firstHeight"`
	MaxHeight   MaxHeight   `json:"maxHeight"`
}

type FirstHeight struct {
	ArrivalTime string `json:"arrivalTime,omitempty"`
	Condition   string `json:"condition,omitempty"`
}

type MaxHeight struct {
	Description string  `json:"description,omitempty"`
	Value       float64 `json:"value,omitempty"`
}

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
	Name      string `json:"name"`
	Grade     string `json:"grade"`
	Immediate bool   `json:"immediate"`
}

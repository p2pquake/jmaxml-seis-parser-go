package vxse

type Head struct {
	Title           string
	ReportDateTime  DateTime
	TargetDateTime  DateTime
	EventID         string
	InfoType        string
	Serial          string
	InfoKind        string
	InfoKindVersion string
	Headline        Headline
}

type Headline struct {
	Text string
}

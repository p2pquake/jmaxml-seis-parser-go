package vxse

import (
	"encoding/xml"
	"time"
)

// W3C XML Schema dateTime 型による日付時刻表記
// FIXME: Golang time.Time 型への変換未対応
type DateTime time.Time

func (dateTime *DateTime) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var s string
	if err := d.DecodeElement(&s, &start); err != nil {
		return err
	}

	dt, _ := time.Parse(time.RFC3339, s)
	*dateTime = DateTime(dt)
	return nil
}

func (dateTime *DateTime) Time() time.Time {
	return time.Time(*dateTime)
}

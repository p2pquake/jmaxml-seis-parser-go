package jmaseis

import (
	"encoding/json"
	"encoding/xml"
	"math"
	"strconv"
	"time"
)

// W3C XML Schema dateTime 型による日付時刻表記
// FIXME: Golang time.Time 型への変換未対応
type DateTime time.Time
type MagnitudeValue float64
type Duration string

func (dateTime *DateTime) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var s string
	if err := d.DecodeElement(&s, &start); err != nil {
		return err
	}

	dt, _ := time.Parse(time.RFC3339, s)
	*dateTime = DateTime(dt)
	return nil
}

func (dateTime *DateTime) MarshalJSON() ([]byte, error) {
	t := time.Time(*dateTime)
	return json.Marshal(t)
}

func (m *MagnitudeValue) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var s string
	if err := d.DecodeElement(&s, &start); err != nil {
		return err
	}

	if s == "NaN" {
		*m = MagnitudeValue(math.NaN())
	}
	v, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return err
	}
	*m = MagnitudeValue(v)
	return nil
}

func (m *MagnitudeValue) MarshalJSON() ([]byte, error) {
	f := float64(*m)
	if math.IsNaN(f) {
		f = -1
	}
	return json.Marshal(f)
}

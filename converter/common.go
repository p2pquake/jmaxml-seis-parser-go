package converter

import (
	"fmt"
	"time"

	"github.com/p2pquake/jmaxml-seis-parser-go/jmaseis"
)

type NotSupportedError struct {
	Key   string
	Value interface{}
}

func (e *NotSupportedError) Error() string {
	return fmt.Sprintf("Not supported: key[%s] value[%#v]", e.Key, e.Value)
}

func EPSPTime(dt jmaseis.DateTime) string {
	loc, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		panic(err)
	}

	return time.Time(dt).In(loc).Format("2006/01/02 15:04:05")
}

func EPSPTimeOrEmpty(dt jmaseis.DateTime) string {
	utcLoc, err := time.LoadLocation("UTC")
	if err != nil {
		panic(err)
	}
	if time.Time(dt).In(utcLoc).IsZero() {
		return ""
	}

	loc, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		panic(err)
	}
	return time.Time(dt).In(loc).Format("2006/01/02 15:04:05")
}

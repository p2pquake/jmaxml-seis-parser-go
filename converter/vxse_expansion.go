package converter

import (
	"time"

	"github.com/p2pquake/jmaxml-vxse-parser-go/vxse"
)

func EPSPTime(dt vxse.DateTime) string {
	loc, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		panic(err)
	}

	return dt.Time().In(loc).Format("2006/01/02 15:04:05")
}

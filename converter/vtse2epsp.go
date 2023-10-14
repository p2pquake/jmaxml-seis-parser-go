package converter

import (
	"errors"
	"strings"

	"github.com/p2pquake/jmaxml-seis-parser-go/epsp"
	"github.com/p2pquake/jmaxml-seis-parser-go/jmaseis"
)

func Vtse2Epsp(vtse jmaseis.Report) (*epsp.JMATsunami, error) {
	// "取消" は未対応
	if vtse.Head.InfoType == "取消" {
		return nil, &NotSupportedError{Key: "vtse.Head.InfoType", Value: vtse.Head.InfoType}
	}

	jmaTsunami := epsp.JMATsunami{
		Expire: nil,
		Issue: epsp.TsunamiIssue{
			Source: strings.Join(vtse.Control.PublishingOffice, "、"),
			Time:   EPSPTime(vtse.Control.DateTime),
			Type:   "Focus",
		},
		Cancelled: judgeCancelled(vtse),
		Areas:     generateAreas(vtse),
	}

	// キャンセルでもないのにエリアがないものは、若干の海面変動のみのため、 EPSP のデータ化対象外
	if jmaTsunami.Cancelled == false && len(jmaTsunami.Areas) == 0 {
		return nil, errors.New("Slight sea-level change is not supported.\n")
	}

	return &jmaTsunami, nil
}

func judgeCancelled(vtse jmaseis.Report) bool {
	// 次をすべて満たす場合、解除と判定
	// 1. LastKind に大津波警報・津波警報・津波注意報がある
	// 2. Kind に大津波警報・津波警報・津波注意報がない

	hasForecastInLastKind := false
	for _, item := range vtse.Body.Tsunami.Forecast.Item {
		if grade(item.Category.LastKind) != "" {
			hasForecastInLastKind = true
			break
		}
	}

	if !hasForecastInLastKind {
		return false
	}

	hasForecastInKind := false
	for _, item := range vtse.Body.Tsunami.Forecast.Item {
		if grade(item.Category.Kind) != "" {
			hasForecastInKind = true
			break
		}
	}

	if hasForecastInKind {
		return false
	}

	return true
}

func generateAreas(vtse jmaseis.Report) []epsp.Area {
	areas := []epsp.Area{}

	for _, item := range vtse.Body.Tsunami.Forecast.Item {
		if grade(item.Category.Kind) != "" {
			areas = append(areas, epsp.Area{
				Name:        item.Area.Name,
				Grade:       grade(item.Category.Kind),
				Immediate:   immediate(item.FirstHeight),
				FirstHeight: firstHeight(item.FirstHeight),
				MaxHeight:   maxHeight(item.MaxHeight),
			})
		}
	}

	return areas
}

func grade(kind jmaseis.Kind) string {
	switch kind.Code {
	case "51":
		return "Warning"
	case "52":
		return "MajorWarning"
	case "53":
		return "MajorWarning"
	case "62":
		return "Watch"
	}

	return ""
}

func immediate(f jmaseis.FirstHeight) bool {
	if f.Condition == "ただちに津波来襲と予測" || f.Condition == "津波到達中と推測" {
		return true
	}

	return false
}

func firstHeight(f jmaseis.FirstHeight) epsp.FirstHeight {
	return epsp.FirstHeight{
		ArrivalTime: EPSPTimeOrEmpty(f.ArrivalTime),
		Condition:   f.Condition,
	}
}

func maxHeight(m jmaseis.MaxHeight) epsp.MaxHeight {
	if m.TsunamiHeight.Condition == "不明" {
		return epsp.MaxHeight{
			Description: m.TsunamiHeight.Description,
			Value:       0,
		}
	}

	return epsp.MaxHeight{
		Description: m.TsunamiHeight.Description,
		Value:       m.TsunamiHeight.Value,
	}
}

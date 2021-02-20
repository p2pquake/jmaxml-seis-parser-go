package converter

import (
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"

	"github.com/p2pquake/jmaxml-vxse-parser-go/epsp"
	"github.com/p2pquake/jmaxml-vxse-parser-go/vxse"
)

type NotSupportedError struct {
	Key   string
	Value interface{}
}

func (e *NotSupportedError) Error() string {
	return fmt.Sprintf("Not supported: key[%s] value[%#v]", e.Key, e.Value)
}

func Vxse2Epsp(vxse vxse.Report) (*epsp.JMAQuake, error) {
	// "取消" は未対応
	if vxse.Head.InfoType == "取消" {
		return nil, &NotSupportedError{Key: "vxse.Head.InfoType", Value: vxse.Head.InfoType}
	}

	jmaQuake := epsp.JMAQuake{
		Code: 551,
		Issue: epsp.Issue{
			Source:  strings.Join(vxse.Control.PublishingOffice, "、"),
			Time:    EPSPTime(vxse.Control.DateTime),
			Type:    issueType(vxse.Head),
			Correct: "None", // FIXME: 未対応
		},
		Earthquake: epsp.Earthquake{
			Time:            EPSPTime(earthquakeTime(vxse)),
			Hypocenter:      hypocenter(vxse),
			MaxScale:        maxScale(vxse),
			DomesticTsunami: domesticTsunami(vxse),
			ForeignTsunami:  foreignTsunami(vxse),
		},
		Points: generatePoints(vxse),
	}
	return &jmaQuake, nil
}

func issueType(head vxse.Head) string {
	if head.Title == "震度速報" {
		return "ScalePrompt"
	}
	if head.Title == "震源に関する情報" {
		return "Destination"
	}
	if head.Title == "震源・震度情報" {
		return "DetailScale"
	}
	if head.Title == "遠地地震に関する情報" {
		return "Foreign"
	}

	if strings.Contains(head.InfoKind, "震度") {
		return "ScalePrompt"
	}
	if strings.Contains(head.InfoKind, "震源") {
		return "Destination"
	}
	if strings.Contains(head.InfoKind, "地震") {
		return "DetailScale"
	}

	return "Other"
}

func hasEarthquake(vxse vxse.Report) bool {
	if vxse.Head.InfoType == "取消" {
		return false
	}

	switch issueType(vxse.Head) {
	case "Destination":
		return true
	case "DetailScale":
		return true
	case "Foreign":
		return true
	}

	return false
}

func earthquakeTime(vxse vxse.Report) vxse.DateTime {
	if hasEarthquake(vxse) {
		return vxse.Body.Earthquake[0].OriginTime
	}

	return vxse.Head.TargetDateTime
}

func hypocenter(vxse vxse.Report) epsp.Hypocenter {
	if !hasEarthquake(vxse) {
		return epsp.Hypocenter{
			Name:      "",
			Latitude:  -200,
			Longitude: -200,
			Depth:     -1,
			Magnitude: -1,
		}
	}

	h := vxse.Body.Earthquake[0].Hypocenter

	name := h.Area.Name
	if h.Area.DetailedName != "" {
		name = h.Area.DetailedName
	}

	// FIXME: マグニチュード "NaN" での動作検証
	m := vxse.Body.Earthquake[0].Magnitude[0]
	magnitude := m.Value
	if m.Condition == "不明" {
		magnitude = -1
	}

	latitude := -200.0
	longitude := -200.0
	depth := -1
	var err error

	c := vxse.Body.Earthquake[0].Hypocenter.Area.Coordinate[0].Value
	exp := regexp.MustCompile("([+-][0-9.]+)([+-][0-9.]+)([+-][0-9.]+)?")
	groups := exp.FindStringSubmatch(c)

	if len(groups) >= 3 {
		latitude, err = strconv.ParseFloat(groups[1], 64)
		if err != nil {
			latitude = -200.0
			log.Panicln(err)
		}

		longitude, err = strconv.ParseFloat(groups[2], 64)
		if err != nil {
			longitude = -200.0
			log.Panicln(err)
		}
	}

	if len(groups) >= 4 {
		depth, err = strconv.Atoi(groups[3])
		if err != nil {
			depth = -1
			log.Panicln(err)
		} else {
			depth /= 1000 * -1
		}
	}

	return epsp.Hypocenter{
		Name:      name,
		Latitude:  latitude,
		Longitude: longitude,
		Depth:     depth,
		Magnitude: magnitude,
	}
}

func maxScale(vxse vxse.Report) int {
	return scale(vxse.Body.Intensity.Observation.MaxInt)
}

func scale(s string) int {
	switch s {
	case "":
		return -1
	case "1":
		return 10
	case "2":
		return 20
	case "3":
		return 30
	case "4":
		return 40
	case "5-":
		return 45
	case "5+":
		return 50
	case "6-":
		return 55
	case "6+":
		return 60
	case "7":
		return 70
	case "震度５弱以上未入電":
		return 46
	}

	return -1
}

func domesticTsunami(vxse vxse.Report) string {
	if issueType(vxse.Head) == "ScalePrompt" {
		return "Checking"
	}

	code := commentCodes(vxse)
	if strings.Contains(code, "0215") || strings.Contains(code, "0230") {
		return "None"
	}
	if strings.Contains(code, "0212") || strings.Contains(code, "0213") || strings.Contains(code, "0214") {
		return "NonEffective"
	}
	if strings.Contains(code, "0211") {
		return "Warning"
	}

	if strings.Contains(code, "0217") || strings.Contains(code, "0229") {
		return "Checking"
	}

	return "Unknown"
}

func foreignTsunami(vxse vxse.Report) string {
	if issueType(vxse.Head) != "Foreign" {
		return "Unknown"
	}

	code := commentCodes(vxse)
	if strings.Contains(code, "0215") {
		return "None"
	}

	if strings.Contains(code, "0221") {
		return "WarningPacificWide"
	}
	if strings.Contains(code, "0222") {
		return "WarningPacific"
	}
	// FIXME: 「北西太平洋」だが太平洋にマッピング
	if strings.Contains(code, "0223") {
		return "WarningPacific"
	}
	if strings.Contains(code, "0224") {
		return "WarningIndianWide"
	}
	if strings.Contains(code, "0225") {
		return "WarningIndian"
	}
	if strings.Contains(code, "0226") {
		return "WarningNearby"
	}
	if strings.Contains(code, "0227") {
		return "NonEffectiveNearby"
	}
	if strings.Contains(code, "0228") {
		return "Potential"
	}

	return "Unknown"
}

func commentCodes(vxse vxse.Report) string {
	return vxse.Body.Comments[0].ForecastComment.Code + vxse.Body.Comments[0].VarComment.Code
}

func generatePoints(vxse vxse.Report) []epsp.Point {
	points := []epsp.Point{}

	// 震度速報: 地域
	if issueType(vxse.Head) == "ScalePrompt" {
		for _, pref := range vxse.Body.Intensity.Observation.Pref {
			for _, area := range pref.Area {
				points = append(points, epsp.Point{
					Pref:   pref.Name,
					Addr:   area.Name,
					IsArea: true,
					Scale:  scale(area.MaxInt),
				})
			}
		}
	}

	// 震源・震度情報: 震度観測点
	if issueType(vxse.Head) == "DetailScale" {
		for _, pref := range vxse.Body.Intensity.Observation.Pref {
			for _, area := range pref.Area {
				for _, city := range area.City {
					for _, station := range city.IntensityStation {
						points = append(points, epsp.Point{
							Pref:   pref.Name,
							Addr:   strings.ReplaceAll(station.Name, "＊", ""),
							IsArea: false,
							Scale:  scale(station.Int),
						})
					}
				}
			}
		}

	}

	return points
}

package converter

import (
	"encoding/json"
	"encoding/xml"
	"io/ioutil"
	"regexp"
	"sort"
	"strings"
	"testing"

	"github.com/p2pquake/jmaxml-vxse-parser-go/epsp"
	"github.com/p2pquake/jmaxml-vxse-parser-go/vxse"
	"github.com/stretchr/testify/assert"
)

func TestSmoke(t *testing.T) {
	testDirectorySmoke(t, "../examples")
	testDirectorySmoke(t, "../data")
}

func TestCompareOldAnalyzer(t *testing.T) {
	dir := "../data"
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		panic(err)
	}

	for _, file := range files {
		if !strings.HasSuffix(file.Name(), ".xml") {
			continue
		}

		data, err := ioutil.ReadFile(dir + "/" + file.Name())
		if err != nil {
			panic(err)
		}

		v := &vxse.Report{}
		err = xml.Unmarshal(data, &v)
		if err != nil {
			t.Errorf("%s parse error: %#v", file.Name(), err)
			continue
		}

		actual, err := Vxse2Epsp(*v)
		if err != nil {
			_, ok := err.(*NotSupportedError)
			if ok {
				continue
			}
			t.Errorf("%s convert error: %#v", file.Name(), err)
		}

		// 比較対象のファイルがある?
		c, ok := searchFile(dir, file.Name(), *actual)
		if !ok {
			t.Errorf("%s has no comparison", file.Name())
			continue
		}

		f, err := ioutil.ReadFile(dir + "/" + c)
		if err != nil {
			panic(err)
		}

		expected := epsp.JMAQuake{}
		err = json.Unmarshal(f, &expected)
		if err != nil {
			panic(err)
		}

		// パーツごとに比較
		assert.Equal(t, expected.Expire, actual.Expire)
		assert.Equal(t, expected.Issue, actual.Issue)
		assert.Equal(t, expected.Earthquake, actual.Earthquake)

		// 震度観測点については、震度観測点名でソートする
		sort.Slice(expected.Points, func(i, j int) bool { return expected.Points[i].Addr > expected.Points[j].Addr })
		sort.Slice(actual.Points, func(i, j int) bool { return actual.Points[i].Addr > actual.Points[j].Addr })
		assert.Equal(t, expected.Points, actual.Points)
	}
}

func testDirectorySmoke(t *testing.T, dir string) {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		panic(err)
	}

	for _, file := range files {
		if !strings.HasSuffix(file.Name(), ".xml") {
			continue
		}

		data, err := ioutil.ReadFile(dir + "/" + file.Name())
		if err != nil {
			panic(err)
		}

		v := &vxse.Report{}
		t.Run("Parsable", func(t *testing.T) {
			err = xml.Unmarshal(data, &v)
			if err != nil {
				t.Errorf("%s parse error: %#v", file.Name(), err)
			}
		})

		var e *epsp.JMAQuake
		t.Run("Convertable", func(t *testing.T) {
			e, err = Vxse2Epsp(*v)
			if err != nil {
				_, ok := err.(*NotSupportedError)
				if ok {
					return
				}
				t.Errorf("%s convert error: %#v", file.Name(), err)
			}
		})

		if e == nil {
			continue
		}

		t.Run("Validate", func(t *testing.T) {
			errors := Validate(file.Name(), e)
			for _, err := range errors {
				t.Error(err)
			}
		})
	}
}

func searchFile(dir string, filename string, e epsp.JMAQuake) (string, bool) {
	// yyyyMMddHHmmss
	// ScalePrompt: 53
	// Destination: 92
	// DetailScale: 95
	// Foreign: 94

	issueType := ""
	if e.Issue.Type == "ScalePrompt" {
		issueType = "53"
	} else if e.Issue.Type == "Destination" {
		issueType = "92"
	} else if e.Issue.Type == "DetailScale" {
		issueType = "95"
	} else if e.Issue.Type == "Foreign" {
		issueType = "94"
	}

	pattern := regexp.MustCompile(filename[:13] + "[0-9]{2}" + issueType + "-\\d+.json")

	files, err := ioutil.ReadDir(dir)
	if err != nil {
		panic(err)
	}

	for _, file := range files {
		if pattern.MatchString(file.Name()) {
			return file.Name(), true
		}
	}

	return "", false
}

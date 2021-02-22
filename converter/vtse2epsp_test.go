package converter

import (
	"encoding/xml"
	"io/ioutil"
	"strings"
	"testing"

	"github.com/p2pquake/jmaxml-vxse-parser-go/epsp"
	"github.com/p2pquake/jmaxml-vxse-parser-go/jmaseis"
)

func TestVTSESmoke(t *testing.T) {
	testVTSEDirectorySmoke(t, "../examples")
	testVTSEDirectorySmoke(t, "../data")
}

func testVTSEDirectorySmoke(t *testing.T, dir string) {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		panic(err)
	}

	for _, file := range files {
		if !strings.HasSuffix(file.Name(), ".xml") || !strings.Contains(file.Name(), "VTSE") {
			continue
		}

		t.Logf("test %s...\n", file.Name())

		data, err := ioutil.ReadFile(dir + "/" + file.Name())
		if err != nil {
			panic(err)
		}

		v := &jmaseis.Report{}
		t.Run("Parsable", func(t *testing.T) {
			err = xml.Unmarshal(data, &v)
			if err != nil {
				t.Errorf("%s parse error: %#v", file.Name(), err)
			}
		})

		var e *epsp.JMATsunami
		t.Run("Convertable", func(t *testing.T) {
			e, err = Vtse2Epsp(*v)
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
			errors := ValidateJMATsunami(file.Name(), e)
			for _, err := range errors {
				t.Error(err)
			}
		})
	}
}

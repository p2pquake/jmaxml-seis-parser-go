package converter

import (
	"encoding/xml"
	"io/ioutil"
	"testing"

	"github.com/p2pquake/jmaxml-vxse-parser-go/epsp"
	"github.com/p2pquake/jmaxml-vxse-parser-go/vxse"
)

func TestSmoke(t *testing.T) {
	testDirectorySmoke(t, "../examples")
	testDirectorySmoke(t, "../data")
}

func testDirectorySmoke(t *testing.T, dir string) {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		panic(err)
	}

	for _, file := range files {
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

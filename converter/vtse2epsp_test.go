package converter

import (
	"encoding/json"
	"encoding/xml"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/p2pquake/jmaxml-seis-parser-go/epsp"
	"github.com/p2pquake/jmaxml-seis-parser-go/jmaseis"
	"github.com/stretchr/testify/assert"
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
			errors := ValidateJMATsunami(file.Name(), v, e)
			for _, err := range errors {
				// 訓練データは ValidationWarning で正しい
				if strings.Contains(err.Error(), "訓練") {
					if !strings.Contains(string(data), "<Status>訓練</Status>") {
						t.Error(err)
					}
				} else if strings.Contains(err.Error(), "試験") {
					if !strings.Contains(string(data), "<Status>試験</Status>") {
						t.Error(err)
					}
				} else {
					t.Error(err)
				}
			}
		})

		jsonFilename := getFileNameWithoutExt(file.Name()) + ".json"
		_, err = os.Stat(dir + "/" + jsonFilename)
		if err != nil {
			continue
		}

		t.Run("Compare", func(t *testing.T) {
			data, err = ioutil.ReadFile(dir + "/" + jsonFilename)
			if err != nil {
				panic(err)
			}

			expected := epsp.JMATsunami{}
			err = json.Unmarshal(data, &expected)
			if err != nil {
				panic(err)
			}

			assert.Equal(t, expected, *e, file.Name())
		})
	}
}

func getFileNameWithoutExt(path string) string {
	return filepath.Base(path[:len(path)-len(filepath.Ext(path))])
}

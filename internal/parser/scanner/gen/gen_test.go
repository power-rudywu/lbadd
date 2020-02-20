package gen

import (
	"io/ioutil"
	"testing"
	"time"

	"github.com/google/uuid"
	"gopkg.in/yaml.v2"
)

func Test_generateScannerInputAndExpectedOutput(t *testing.T) {
	// t.SkipNow() // skip this, only generate test cases manually

	for i := 0; i < 1; i++ {
		start := time.Now()
		scIn, scOut := generateScannerInputAndExpectedOutput()
		t.Logf("took %v\n", time.Since(start).Round(time.Millisecond))

		tc := generateTestCase(scIn, scOut)

		yamlOut, err := yaml.Marshal(tc)
		if err != nil {
			t.Error(err)
		}

		err = ioutil.WriteFile("../testdata/"+uuid.New().String(), yamlOut, 0600)
		if err != nil {
			t.Error(err)
		}
	}
}

package test

import (
	"flag"
	"io/ioutil"
	"path/filepath"
	"testing"
)

var update = flag.Bool("update", false, "update .golden files")

// HandleGoldenFile will read or update the golden file depending on -update flag is provided.
//
// If you run your tests with "go test ./... -v -update" the golden files will be updated.
func HandleGoldenFile(name string, t *testing.T, data []byte) []byte {
	gp := filepath.Join("testdata", name+".golden")
	if *update {
		t.Logf("update %s.golden file", t.Name())
		if err := ioutil.WriteFile(gp, data, 0644); err != nil {
			t.Fatalf("failed to update %s.golden file: %s", t.Name(), err)
		}
	}

	expected, err := ioutil.ReadFile(gp)
	if err != nil {
		t.Fatalf("failed reading %s.golden: %s", t.Name(), err)
	}

	return expected
}

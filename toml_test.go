package tomladapter

import (
	"bytes"
	"os"
	"path/filepath"
	"sort"
	"testing"
)

func TestFixtures(t *testing.T) {
	fixtures, err := filepath.Glob("testdata/fixture*")
	if err != nil {
		t.Fatal(err)
	}
	l := len(fixtures)
	if l == 0 {
		t.Fatal("no fixtures found")
	}
	if l%2 != 0 {
		t.Fatal("odd number of fixtures")
	}
	sort.Strings(fixtures)

	a := Adapter{}

	for i := 0; i < l; i += 2 {
		jn, err := os.ReadFile(fixtures[i])
		if err != nil {
			t.Error(err)
			continue
		}
		tm, err := os.ReadFile(fixtures[i+1])
		if err != nil {
			t.Error(err)
			continue
		}

		b, _, err := a.Adapt(tm, nil)
		if err != nil {
			t.Error(err)
			continue
		}

		if !bytes.Equal(bytes.TrimSpace(jn), b) {
			t.Errorf("fixture %q failed", fixtures[i])
			t.Logf("expected: %s\nactual: %s", jn, b)
		}
	}
}

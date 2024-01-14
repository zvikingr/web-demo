package fstools

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
)

func TestEnsureDir(t *testing.T) {
	dir, err := ioutil.TempDir("", "tmp")
	if err != nil {
		t.Fatal(err)
	}

	subDir := filepath.Join(dir, "dir")

	cleanup := func() {
		if err = os.RemoveAll(dir); err != nil {
			t.Error(err)
		}
	}

	if err = EnsureDir(subDir); err != nil {
		cleanup()
		t.Fatal(err)
	}

	cleanup()
}

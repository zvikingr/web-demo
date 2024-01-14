package log

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
)

func TestLog(t *testing.T) {
	dir, err := ioutil.TempDir("", "tmp")
	if err != nil {
		t.Fatal(err)
	}

	subDir := filepath.Join(dir, "log")

	cleanup := func() {
		if err = os.RemoveAll(dir); err != nil {
			t.Error(err)
		}
	}

	InitLog(&Config{
		LogFile:    filepath.Join(subDir, "test.log"),
		Level:      "fatal",
		MaxSize:    500,
		MaxBackups: 10,
		MaxAge:     10,
		Compress:   false,
	})

	Debug("debug")
	Info("info")
	Warn("warn")
	Error("error")

	Debugf("debugf")
	Infof("infof")
	Warnf("warnf")
	Errorf("errorf")

	cleanup()
}

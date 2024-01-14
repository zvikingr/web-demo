package fstools

import "os"

// EnsureDir Make sure the directory exists, create it if it doesn't
func EnsureDir(path string) error {
	if _, err := os.Stat(path); !os.IsNotExist(err) {
		return err
	}

	return os.Mkdir(path, os.ModePerm)
}

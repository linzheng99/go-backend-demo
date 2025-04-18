package tool

import (
	"os"
	"path/filepath"
)

func GetRootDir() (string, error) {
	dir, err := os.Getwd()
	if err != nil {
		return "", err
	}

	rootDir, err := filepath.Abs(dir)
	if err != nil {
		return "", err
	}

	return rootDir, nil
}

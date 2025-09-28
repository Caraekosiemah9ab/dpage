package utils

import (
	"os"
	"path/filepath"
)

func GetCurrentDir() (string, error) {
	ex, err := os.Executable()
	if err != nil {
		return "", err
	}
	return filepath.Dir(ex), nil
}

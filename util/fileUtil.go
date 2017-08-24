package util

import (
	"path/filepath"
	"os"
)

func IsFileExistent(path string) (bool, error){
	abs, abErr := filepath.Abs(path)
	if abErr != nil {
		return false, abErr
	}
	_, statErr := os.Stat(abs)
	if os.IsNotExist(statErr) {
		return false, nil
	} else {
		return true, nil
	}
}

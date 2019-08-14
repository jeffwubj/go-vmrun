package utils

import "os"

// FileExists checks whether file existed
func FileExists(path string) bool {
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		return false
	} else if err == nil {
		return true
	}
	return false
}

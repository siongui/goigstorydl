package igstorydl

import "os"

// Given folder path, create the folder if the folder does not exist.
func CreateDirIfNotExist(dir string) (err error) {
	if _, err = os.Stat(dir); os.IsNotExist(err) {
		err = os.MkdirAll(dir, 0755)
	}
	return
}

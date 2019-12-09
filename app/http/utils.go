package http

import (
	"io/ioutil"
	"path/filepath"
)

func readFileFromPathToString(subfolder string, filename string) string {
	path := filepath.Join(subfolder, filename)

	bytes, err := ioutil.ReadFile(path)

	if err != nil {
		panic("not valid options to load/read the file.")
	}

	return string(bytes)
}
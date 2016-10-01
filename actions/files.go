package actions

import (
	"io/ioutil"
)

// SaveToFile will save some data into a filePath (shortcut)
// Note: it will set the file permission to 0600 for security
func SaveToFile(filePath string, data []byte) {
	err := ioutil.WriteFile(filePath, data, 0600)
	if err != nil {
		panic(err)
	}
}

package services

import (
	"aurora-borealis/utils"
	"io/ioutil"
	"os"
	"path/filepath"
)

// SaveToLocal saves the markdown content to the local file system
func SaveToLocal(title, content string) error {
	// Ensure the local directory exists
	if _, err := os.Stat(utils.LocalDir); os.IsNotExist(err) {
		err := os.MkdirAll(utils.LocalDir, os.ModePerm)
		if err != nil {
			return err
		}
	}

	// Create the markdown file
	filename := filepath.Join(utils.LocalDir, title+".md")
	err := ioutil.WriteFile(filename, []byte(content), 0644)
	return err
}

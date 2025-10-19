package paths

import (
	"io/fs"
	"path/filepath"
	"strings"
)

func FindEncryptedFiles(root string) ([]string, error) {
	var encryptedFiles []string

	err := filepath.WalkDir(root, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if d.IsDir() {
			return nil
		}

		if strings.Contains(d.Name(), ExtensionPrefix) {
			encryptedFiles = append(encryptedFiles, path)
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return encryptedFiles, nil
}

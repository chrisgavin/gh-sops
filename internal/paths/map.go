package paths

import (
	"path/filepath"
	"strings"
)

func GetDecryptedPath(encryptedPath string) string {
	directory, encryptedFile := filepath.Split(encryptedPath)
	decryptedFile := strings.Replace(encryptedFile, ExtensionPrefix, "", 1)
	return filepath.Join(directory, decryptedFile)
}

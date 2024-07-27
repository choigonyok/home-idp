package file

import (
	"io/fs"
	"os"
	"strings"
)

type FileReader struct {
	Filepath string
}

func NewReader(filepath string) *FileReader {
	return &FileReader{
		Filepath: filepath,
	}
}

func ReadFile(filepath string) ([]byte, error) {
	return os.ReadFile(strings.TrimSpace(filepath))
}

func Exist(filepath string) bool {
	// We must explicitly check if the error is due to the file not existing (as opposed to a
	// permissions error).
	_, err := os.Stat(filepath)
	if err == fs.ErrNotExist {
		return false
	}
	return true
}

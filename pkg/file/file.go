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

func exist(filepath string) bool {
	_, err := os.Stat(filepath)
	if err == fs.ErrNotExist {
		return false
	}
	return true
}

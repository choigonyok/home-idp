package file

import (
	"io"
	"os"
	"strings"
)

func ReadFile(filename string, stdinReader io.Reader) (string, error) {
	var contentBuf []byte
	var err error
	contentBuf, err = os.ReadFile(strings.TrimSpace(filename))
	if err != nil {
		return "", err
	}
	if err != nil {
		return "", err
	}
	return string(contentBuf), nil
}

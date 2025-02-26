package templates

import (
	"io"
	"os"
	"strings"
)

// Get html template as string
func GetTemplate(filename string) (string, error) {
	file, err := os.Open("./templates/" + filename)
	if err != nil {
		return "", err
	}
	defer file.Close()

	var buf strings.Builder
	_, err = io.Copy(&buf, file)
	if err != nil {
		return "", err
	}

	return buf.String(), nil
}

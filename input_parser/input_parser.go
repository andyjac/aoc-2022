package inputparser

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
)

func Parse(problem string) (string, error) {
	cwd, err := filepath.Abs("./")
	if err != nil {
		return "", err
	}

	file := filepath.Join(cwd, fmt.Sprintf("/input_parser/inputs/%s.txt", problem))
	bytes, err := ioutil.ReadFile(file)
	return string(bytes), err
}

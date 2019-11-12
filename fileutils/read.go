package fileutils

import (
	"fmt"
	"io/ioutil"
	"strings"
)

//FileToSlice ...
func FileToSlice(path string) ([]string, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Println("File reading error", err)
		return nil, err
	}
	s := strings.Split(string(data), "\n")
	return s, nil
}

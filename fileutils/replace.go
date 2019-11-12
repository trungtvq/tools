package fileutils

import (
	"fmt"
	"os"
	"regexp"
)

//UpdateFileWithPaternForFirstLineMatch  example pattern: "p([a-z]+)ch"
//this func will help you read file, find pattern in each line, then replace that line with newValue
func UpdateFileWithPaternForFirstLineMatch(file string, pattern string, newValue string) bool {
	arr, err := FileToSlice(file)
	if err != nil {
		return false
	}
	for i := 0; i < len(arr); i++ {
		isMatch, err := regexp.MatchString(pattern, arr[i])
		if err != nil {
			return false
		}
		if isMatch == true {
			arr[i] = newValue

			var result string
			for i := 0; i < len(arr); i++ {
				result = result + arr[i] + "\n"
			}

			f, err := os.Create(file)
			defer f.Close()

			d2 := []byte(result)
			n, err := f.Write(d2)
			if err != nil {
				return false
			}
			fmt.Println("wrote", n, "bytes in line", i+1)
			return true
		}
	}
	return false
}

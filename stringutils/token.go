package stringutils

import (
	"sort"
)

//GetLike .
func GetLike(arr1 []string, arr2 []string) []string {
	var arr []string
	for _, e1 := range arr1 {
		for _, e2 := range arr2 {
			if e1 == e2 {
				arr = append(arr, e1)
				break
			}
		}
	}
	return arr
}
func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
func DeleteReduction(array []string) []string {
	var result []string
	for i, _ := range array {
		if !contains(result, array[i]) {
			result = append(result, array[i])
		}
	}
	return result

}
func CreateSet(arrays ...[]string) []string {
	var result []string
	for _, array := range arrays {
		for _, e1 := range array {
			if !contains(result, e1) {
				result = append(result, e1)
			}
		}
	}
	return result
}

func CreateSort(arrays ...[]string) []string {
	var result []string
	for _, array := range arrays {
		for _, e1 := range array {
			if !contains(result, e1) {
				result = append(result, e1)
			}
		}
	}

	sort.Slice(result, func(i, j int) bool {
		return result[i] > result[j]
	})
	return result
}

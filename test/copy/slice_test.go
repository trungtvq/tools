package test

import (
	"fmt"
	"testing"
)
func copySlice(b[]byte)[]byte{
	return b
}
func copySlicePointer(b*[]byte)*[]byte{
	return b
}
func BenchMarkCopySlice(b *testing.B){
	data := make([]byte, int(1e7), int(1e7)) // Initialize an empty byte slice
	fmt.Println(data)

}
package test

import (
	"fmt"
	"github.com/trungtvq/go-utils/define"
	"github.com/trungtvq/go-utils/test/FakeData"
	"testing"
)
func createInstanceOfStruct(t define.Film) define.Film{
	f:=FakeData.Film
	return f
}
func createInstanceOfStructPointer(t *define.Film) define.Film{
	f:=t
	return *f
}
func createInstanceOfStructReturnPointer(t define.Film) *define.Film{
	f:=t
	return &f
}

func BenchmarkCopyStruct(b *testing.B){
	f:=FakeData.Film
	for i,_:= range f.Data{
		f.Data[i]="lala"
	}

	for i:=0;i<b.N;i++{
		for j:=0;j<100;j++{
			n:=createInstanceOfStructReturnPointer(f)
			if false{
				fmt.Println(n)
			}

		}

	}


}
//out data, in pointer :37.6
//out data, in data:
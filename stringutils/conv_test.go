package stringutils_test

import (
	"github.com/trungtvq/go-utils/define"
	"github.com/trungtvq/go-utils/stringutils"
	"testing"
)

func TestJSONToObjs(t *testing.T) {

}
func TestCreateMap(t *testing.T) {

}
func TestUnderScore2UpperCamelCase(t *testing.T) {

}
func TestHmacEncrypt(t *testing.T) {

}
func TestGetLike2(t *testing.T) {

}
func BenchmarkCreateMap(b *testing.B) {
	f:=define.Film{
		Name:       "Warcraft",
		Link:       "fshare.vn",
		Avatar:     "avatar.com",
		PublicDate: "19-12-12",
		LikeNum:    10,
	}
	for i:=0;i<b.N;i++{
		stringutils.CreateMap(&f)
	}

}
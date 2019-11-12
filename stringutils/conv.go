package stringutils

import (
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"log"
	"reflect"
	"strconv"
	"strings"
)

//NewParams conv muti var to (2i: key, 2i+1: value)
func NewParams(obj ...interface{}) map[string]string {
	var m map[string]string
	for i := 1; i < len(obj); i += 2 {
		m[fmt.Sprintf("%v", obj[i-1])] = fmt.Sprintf("%v", obj[i])
	}
	return m
}

//JSONToObjs unmarshal string to mutil struct
func JSONToObjs(js string, objs ...interface{}) error {
	bs := []byte(js)
	for _, obj := range objs {
		err := json.Unmarshal(bs, obj)
		if err != nil {
			return err
		}
	}
	return nil
}

//CreateMap with key as field name && field valude as value
func CreateMap(f interface{}) map[string]string {
	m := make(map[string]string)
	r := reflect.ValueOf(f).Elem()

	for i := 0; i < r.NumField(); i++ {
		val := r.Field(i)

		switch val.Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			m[strings.ToLower(r.Type().Field(i).Name)] = strconv.FormatInt(val.Int(), 10)
		case reflect.String:
			m[strings.ToLower(r.Type().Field(i).Name)] = val.String()
		case reflect.Float64, reflect.Float32:
			m[strings.ToLower(r.Type().Field(i).Name)] = fmt.Sprintf("%v", val.Float())
			// etc...
		}
	}
	return m
}
func Dump(m map[string]string) {
	for k, v := range m {
		log.Printf("%s : %s\n", k, v)
	}
}

// HmacEncrypt hmac data to hmac field of zalopay
func HmacEncrypt(ctx context.Context, params ...string) string {
	mixString := strings.Join(params, "|")
	h := hmac.New(sha256.New, []byte("secret"))
	h.Write([]byte(mixString))
	return hex.EncodeToString(h.Sum(nil))
}

//UnderScore2UpperCamelCase .
func UnderScore2UpperCamelCase(str string) string {
	// var result []byte
	// if str[0] > 96 && str[0] < 123 {
	// 	result = result
	// 	str[0] = str[0] >> 32
	// }
	// for i, c := range str {
	// 	if c == 95 {
	// 		if str[i+1] > 96 && str[i+1] < 123 {
	// 			str[i+1] = str[i+1] - 32
	// 		}
	// 	}
	// }
	return str
}

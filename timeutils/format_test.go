package timeutils_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/trungtvq/go-utils/timeutils"
)

func TestGetYYMMDD(t *testing.T) {
	assert.Equal(t, len(timeutils.GetYYMMDD(time.Now().Unix())), 6)
	fmt.Printf("Now, time is: %s\n", timeutils.GetYYMMDD(time.Now().Unix()))

}

func BenchmarkGetYYMMDD(b *testing.B) {
	x := time.Now().Unix()
	for i := 0; i < b.N; i++ {
		timeutils.GetYYMMDD(x) //218ns/op=135+83
	}
}

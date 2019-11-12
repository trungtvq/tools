package timeutils

import (
	"time"
)

//GetYYMMDD look like 190415 as result of xxxxxxx is nsec from 1970
func GetYYMMDD(t int64) string {
	return time.Unix(t, 0).Format("060102")
}

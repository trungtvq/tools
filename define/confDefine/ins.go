package confDefine

import db "github.com/trungtvq/go-utils/service/db/models"

type Ins string
var (
	MySqlMain Ins = "MySqlMain"
	MySqlTmp Ins = "MySqlTmp"
	Redis Ins = "Redis"
)
var instance = map[Ins] interface{}{
	MySqlMain: &db.DB{},
	MySqlTmp: &db.DB{},
	Redis: &db.DB{},
}

func (i Ins) Get() interface{} {
	p, ok := instance[i]
	if ok {
		return p
	}
	return ""
}
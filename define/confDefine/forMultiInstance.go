package confDefine

/*/////////////////////////////////////////////
////	Copy to use, DO NOT import it	//////
///////////////////////////////////////////*/

type ConfPrefix uint8
var (
	pre1 ConfPrefix =  1
	pre2 ConfPrefix = 2
)

var prefix = map[ConfPrefix]string{
	pre1    : "SQLSRC",
	pre2  : "SQLDES",
}

func (c ConfPrefix) GetPrefix() string {
	p, ok := prefix[c]
	if ok {
		return p
	}
	return ""
}
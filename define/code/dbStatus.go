package code

type Status string

var (
	INVALID  Status = "INVALID"
	ERROR    Status = "ERROR"
	CREATED  Status = "CREATED"
	WAITING  Status = "WAITING"
	ACCEPTED Status = "ACCEPTED"
	REJECTED Status = "REJECTED"
	EXPIRED  Status = "EXPIRED"
)

//ErrorMessage error message from code
var ErrorMessage = map[Status]int32{
	INVALID:  0,
	ERROR:    1,
	CREATED:  2,
	WAITING:  3,
	ACCEPTED: 4,
	REJECTED: 5,
	EXPIRED:  6,
}

func GetCode(message Status) int32 {
	code, ok := ErrorMessage[message]
	if ok {
		return code
	}
	return 0
}

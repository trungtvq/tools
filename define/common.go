package define

type Film struct {
	Name       string
	Link       string
	Avatar     string
	PublicDate string
	LikeNum    int64
	Data       [99999]string
}
type IRead interface {
	Read()
}

package ctype

import "encoding/json"

type Location int

const (
	Local Location = 1 // 本地
	QiNiu Location = 2
)

func (l Location) MarshalJSON() ([]byte, error) {
	return json.Marshal(l.Parse())
}

func (l Location) Parse() string {
	var str string
	switch l {
	case Local:
		str = "Local"
	case QiNiu:
		str = "QiNiu"
	default:
		str = "其他"
	}
	return str
}

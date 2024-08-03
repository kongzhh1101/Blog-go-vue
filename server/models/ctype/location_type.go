package ctype

import "encoding/json"

type Location int

const (
	Local Location = 1 // 本地
	QiNiu Location = 2
)

func (s Location) MarshalJSON() ([]byte, error) {
	return json.Marshal(s)
}

func (l Location) ParseLocation() string {
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

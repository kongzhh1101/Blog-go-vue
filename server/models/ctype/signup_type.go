package ctype

import "encoding/json"

type SignupSource int

const (
	QQ     SignupSource = 1
	Github SignupSource = 2
	Email  SignupSource = 3
)

func (s SignupSource) MarshalJSON() ([]byte, error) {
	return json.Marshal(s)
}

func (s SignupSource) ParseRole() string {
	var str string
	switch s {
	case QQ:
		str = "QQ"
	case Github:
		str = "Github"
	case Email:
		str = "Email"
	default:
		str = "其他"
	}
	return str
}

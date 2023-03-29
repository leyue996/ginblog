package ctype

import "encoding/json"

type SignStatus int

const (
	SignQQ     SignStatus = 1 //qq
	SignEmail  SignStatus = 2 //邮箱
	SignGithub SignStatus = 3 //github

)

func (s SignStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.String())
}
func (s SignStatus) String() string {
	var str string
	switch s {
	case SignQQ:
		str = "管理员"
	case SignEmail:
		str = "用户"
	case SignGithub:
		str = "游客"
	default:
		str = "其他"
	}
	return str
}

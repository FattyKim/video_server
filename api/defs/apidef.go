package defs

//requests
type UserCredential struct {
	//序列化时会自动转换成JSON格式
	Username string `json:"user_name"`
	Pwd      string `json:"pwd"`
}

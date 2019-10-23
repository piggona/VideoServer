package defs

// UserCredential requests 声明用户请求数据结构
type UserCredential struct {
	Username string `json:"user_name"`
	Pwd      string `json:"pwd`
}

// Data model
type VideoInfo struct {
	Id           string
	AuthorId     int
	Name         string
	DisplayCtime string
}

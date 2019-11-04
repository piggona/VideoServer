package defs

// UserCredential requests 声明用户请求数据结构
type UserCredential struct {
	Username string `json:"user_name"`
	Pwd      string `json:"pwd`
}

// response
type SignedUp struct {
	Success   bool   `json:"success"`
	SessionId string `json:"session_id"`
}

// Video info model
type VideoInfo struct {
	Id           string
	AuthorId     int
	Name         string
	DisplayCtime string
}

// Comment info model
type Comment struct {
	Id      string
	VideoId string
	Author  string
	Content string
}

type SimpleSession struct {
	Username string
	TTL      int64
}

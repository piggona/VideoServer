package main

import (
	"net/http"
	"video_server/api/defs"
	"video_server/api/session"
)

var HEADER_FILED_SESSION = "X-Session-Id"
var HEADER_FILED_UNAME = "X-User-Name"

// validateUserSession:
// Session是否存在（用户是否存在有效的Session）
// 1.从Request Header中获取Session Id(如果没有sessionId则为第一次登录的新用户)
// 2.验证获取到的session Id是否已经失效
// 3.未失效则将从数据库获得的用户信息（此处是userId）加入到request头中，说明已经登录
func validateUserSession(r *http.Request) bool {
	sid := r.Header.Get(HEADER_FILED_SESSION)
	if len(sid) == 0 {
		return false
	}

	uname, ok := session.IsSessionExpired(sid)
	if ok {
		return false
	}
	r.Header.Add(HEADER_FILED_UNAME, uname)
	return true
}

// 判断是否是已经登录的状态
func ValidateUser(w http.ResponseWriter, r *http.Request) bool {
	uname := r.Header.Get(HEADER_FILED_UNAME)
	if len(uname) == 0 {
		sendErrorResponse(w, defs.ErrorRequestBodyParseFailed)
		return false
	}
	return true

}

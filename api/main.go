package main

import (
	"net/http"
	"video_server/api/session"

	"github.com/julienschmidt/httprouter"
)

type middlewareHandler struct {
	r *httprouter.Router
}

func NewMiddleWareHandler(r *httprouter.Router) http.Handler {
	m := &middlewareHandler{}
	m.r = r
	return m
}

func (m *middlewareHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// check session
	validateUserSession(r)
	m.r.ServeHTTP(w, r)
}

// 业务功能：
// upload,logout,ListVideos,ShowComments,PostComment,DeleteVideo

func RegisterHandlers() *httprouter.Router {
	router := httprouter.New()

	router.POST("/user", CreateUser)

	router.POST("/user/:user_name", Login)

	// router.POST("/user/:username/videos", AddNewVideo)

	// router.GET("/user/:username/videos", ListAllVideos)

	// router.DELETE("/user/:username/videos/:vid-id", DeleteVideo)

	// router.POST("/videos/:vid-id/comments", PostComment)

	// router.GET("/videos/:vid-id/comments", ShowComments)

	return router
}

func Prepare() {
	session.LoadSessionsFromDB()
}

func main() {
	Prepare()
	r := RegisterHandlers()
	m := NewMiddleWareHandler(r)
	http.ListenAndServe("localhost:8000", m)
}

package main

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"text/template"

	"github.com/julienschmidt/httprouter"
)

type HomePage struct {
	Name string
}

type UserPage struct {
	Name string
}

func HomeHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	cname, err1 := r.Cookie("username")
	sid, err2 := r.Cookie("session")

	if err1 != nil || err2 != nil {
		// 缺少某个Cookie则进行登录页面的渲染（说明用户从来没有登录过）
		p := &HomePage{Name: "haohao"}
		t, e := template.ParseFiles("./templates/home.html")

		if e != nil {
			log.Printf("Parsing template home.html error: %s", e)
			return
		}

		t.Execute(w, p)
		return
	}
	if len(cname.Value) != 0 && len(sid.Value) != 0 {
		// 若不缺少Cookie则跳转到已经登录的页面（说明用户登录过，或是cheater:更改了Cookie）
		http.Redirect(w, r, "/userhome", http.StatusFound)
		return
	}
}

func userHomeHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	cname, err1 := r.Cookie("username")
	_, err2 := r.Cookie("session")

	if err1 != nil || err2 != nil {
		// 缺少Cookie字段则跳转到登录页面
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}

	// 渲染已登录界面，然后前端通过Ajax向api后端发起数据请求。
	// 如果api请求的session校验没有通过，则在前端设置Cookie为空，并跳转到登录页面。
	fname := r.FormValue("username")

	var p *UserPage
	if len(cname.Value) != 0 {
		p = &UserPage{Name: cname.Value}
	} else if len(fname) != 0 {
		p = &UserPage{Name: fname}
	}

	t, e := template.ParseFiles("./templates/userhome.html")
	if e != nil {
		log.Printf("Parsing userhome.html error: %s", e)
		return
	}

	t.Execute(w, p)
	return
}

func apiHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// 传来的request是否有问题
	if r.Method != http.MethodPost {
		re, _ := json.Marshal(ErrorRequestNotRecognized)
		io.WriteString(w, string(re))
		return
	}

	res, _ := ioutil.ReadAll(r.Body)
	apiBody := &ApiBody{}
	if err := json.Unmarshal(res, apiBody); err != nil {
		re, _ := json.Marshal(ErrorRequestParseFailed)
		io.WriteString(w, string(re))
		return
	}

	request(apiBody, w, r)
	defer r.Body.Close()
}

package main

import (
	"log"
	"net/http"
	"text/template"

	"github.com/julienschmidt/httprouter"
)

type HomePage struct {
	Name string
}

func HomeHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	p := &HomePage{Name: "haohao"}
	t, e := template.ParseFiles("./template/home.html")

	if e != nil {
		log.Printf("Parsing template home.html error: %s", e)
		return
	}

	t.Execute(w, p)
	return
}

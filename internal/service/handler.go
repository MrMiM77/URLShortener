package service

import (
	"github.com/MrMiM77/URLShortener.git/internal/db"
	"github.com/MrMiM77/URLShortener.git/internal/models"
	"html/template"
	"net/http"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[1:]
	if title == "encode" {
		encodeHandler(w,r)
	}else if title == "decode"{
		decodeHandler(w,r)
	}else {
		redirectHandler(w, r,title)
	}
}

func encodeHandler(w http.ResponseWriter, r *http.Request){
	tmp1 := template.Must(template.ParseFiles("./static/encode.html"))
	tmp1.Execute(w,"")
}
func decodeHandler(w http.ResponseWriter, r *http.Request){
	tmp1 := template.Must(template.ParseFiles("./static/decode.html"))
	tmp1.Execute(w,"")
}
func redirectHandler(w http.ResponseWriter, r *http.Request,title string){
	database := db.GetInstance()
	var urlMap models.URLMap
	database.Where("shorturl = ?",title).First(urlMap)
}

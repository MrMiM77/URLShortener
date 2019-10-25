package service

import (
	"crypto/sha256"
	"fmt"
	"github.com/MrMiM77/URLShortener.git/config"
	"github.com/MrMiM77/URLShortener.git/internal/db"
	"github.com/MrMiM77/URLShortener.git/internal/models"
	"html/template"
	"net/http"
	"strconv"
	"strings"
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

type ShortUrl struct {
	Url string
}
func encodeHandler(w http.ResponseWriter, r *http.Request){
	tmp1 := template.Must(template.ParseFiles("./template/encode.html"))
	shortUrl := ShortUrl{}
	if r.Method == http.MethodPost{
		if err := r.ParseForm(); err!= nil{
			fmt.Fprintf(w, "ParseForm () err : %v", err)
			return
		}
		url := r.FormValue("url")
		if !strings.Contains(url, "http"){
			url = "http://" + url
		}
		fmt.Println(url)
		shortUrl = calculateUrl(url)
	}
	fmt.Println(shortUrl)
	tmp1.Execute(w,shortUrl)
}
func calculateUrl(url string) ShortUrl{
	database := db.GetInstance()
	hash := fmt.Sprintf("%x", sha256.Sum256([]byte(url)))[:6]
	urlMap := models.URLMap{}
	database.FirstOrCreate(&urlMap, models.URLMap{AbsoluteURL:url,ShortURL:hash})
	return ShortUrl{Url:config.GetInstance().SERVER.HOST+":"+strconv.Itoa(config.GetInstance().SERVER.PORT)+"/"+hash}
}
func decodeHandler(w http.ResponseWriter, r *http.Request){
	tmp1 := template.Must(template.ParseFiles("./template/decode.html"))
	tmp1.Execute(w,"")
}
func redirectHandler(w http.ResponseWriter, r *http.Request,title string){
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	database := db.GetInstance()
	var urlMap models.URLMap
	var count int
	if database.Model(&urlMap).Where("short_url = ?",title).Count(&count); count!=1{
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("url not found"))
		return
	}
	database.Where("short_url = ?",title).First(&urlMap)
	http.Redirect(w, r, urlMap.AbsoluteURL, http.StatusMovedPermanently)
}

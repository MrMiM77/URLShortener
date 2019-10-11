package service

import (
	"fmt"
	"io"
	"net/http"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[1:]
	if title == "encode" {

	}else if title == "decode"{

	}else {
		redirectHandler(w, r)
	}
}

func redirectHandler(w http.ResponseWriter, r *http.Request){
	fmt.Println(title)
	io.WriteString(w, "hello world")
}

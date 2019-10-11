package service

import "net/http"

func RunServer()  {
	http.HandleFunc("/", Handler)
	http.ListenAndServe(":8080", nil)

	
}

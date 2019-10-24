package main

import (
	"github.com/MrMiM77/URLShortener.git/config"
	"github.com/MrMiM77/URLShortener.git/internal/service"
)

func main() {
	_, err := config.GetInstance()
	if err != nil{
		panic(err)
	}

	service.RunServer()
}

package main

import (
	"log"
	"net/http"

	"github.com/miguelfoliveira/go-web-service/product"
)

const (
	basePath   = "/api"
	listenPort = ":3000"
)

func main() {
	product.SetupRoutes(basePath)
	err := http.ListenAndServe(listenPort, nil)
	if err != nil {
		log.Fatal(err)
	}
}

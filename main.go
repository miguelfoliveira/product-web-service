package main

import (
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/miguelfoliveira/go-web-service/database"
	"github.com/miguelfoliveira/go-web-service/product"
	"github.com/miguelfoliveira/go-web-service/receipt"
)

const (
	basePath   = "/api"
	listenPort = ":3000"
)

func main() {
	database.SetupDatabse()
	receipt.SetupRoutes(basePath)
	product.SetupRoutes(basePath)
	err := http.ListenAndServe(listenPort, nil)
	if err != nil {
		log.Fatal(err)
	}
}

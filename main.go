package main

import (
	"log"
	"net/http"

	"github.com/jameselliothart/inventoryservice/database"
	"github.com/jameselliothart/inventoryservice/receipt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jameselliothart/inventoryservice/product"
)

const apiBasePath = "/api"

func main() {
	database.SetupDatabase()
	receipt.SetupRoutes(apiBasePath)
	product.SetupRoutes(apiBasePath)
	log.Fatal(http.ListenAndServe(":5000", nil))
}

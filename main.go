package main

import (
	"github.com/jameselliothart/inventoryservice/database"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jameselliothart/inventoryservice/product"
)

const apiBasePath = "/api"

func main() {
	database.SetupDatabase()
	product.SetupRoutes(apiBasePath)
	http.ListenAndServe(":5000", nil)
}

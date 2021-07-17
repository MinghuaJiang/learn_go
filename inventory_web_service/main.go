package main

import (
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/minghuajiang/learn_go/inventory_web_service/database"
	"github.com/minghuajiang/learn_go/inventory_web_service/product_v2"
	"github.com/minghuajiang/learn_go/inventory_web_service/receipt"
)

const apiBasePath = "/api"

func main() {
	database.SetupDatabases()
	receipt.SetupRoutes(apiBasePath)
	product_v2.SetupRoutes(apiBasePath)
	log.Fatal(http.ListenAndServe(":5000", nil))
}

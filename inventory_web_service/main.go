package main

import (
	"net/http"

	"github.com/minghuajiang/learn_go/inventory_web_service/product"
)

const apiBasePath = "/api"

func main() {
	product.SetupRoutes(apiBasePath)
	http.ListenAndServe(":5000", nil)
}

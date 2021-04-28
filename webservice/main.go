package main

import (
	"net/http"

	"github.com/minghuajiang/webservice/controllers"
)

func main() {
	controllers.RegisterControllers()
	http.ListenAndServe(":3000", nil)
}

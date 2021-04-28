package main

import (
	"fmt"

	"github.com/minghuajiang/webservice/models"
)

func main() {
	u := models.User{
		ID:        1,
		FirstName: "Pocky",
		LastName:  "Jiang",
	}

	fmt.Println(u)
}

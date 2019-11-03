package main

import (
	"net/http"

	"github.com/sk000f/go-web-service/controllers"
)

func main() {
	controllers.RegisterControllers()
	http.ListenAndServe(":3000", nil)
}

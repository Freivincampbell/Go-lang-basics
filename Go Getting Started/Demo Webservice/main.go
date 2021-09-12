package main

import (
	"Demo_Webservice/controllers"
	"net/http"
)

func main() {
	controllers.RegisterControllers()
	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		return 
	}
}

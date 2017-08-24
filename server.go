package main

import (
	"net/http"
	"jwt_golang/routes"
)

func main() {
	http.ListenAndServe("localhost:8080", routes.BASE_ROUTER)
}
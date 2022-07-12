package main

import (
	"net/http"

	asset "github.com/diegogenta1980/idealAPI/asset"
)

const apiBasePath = "/api"

func main() {
	println("Initial structure")
	asset.SetupRoutes(apiBasePath)
	http.ListenAndServe("localhost:3000", nil)
}

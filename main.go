package main

import (
	"net/http"

	asset "github.com/diegogenta1980/idealAPI/asset"
	"github.com/diegogenta1980/idealAPI/database"
)

const apiBasePath = "/api"

func main() {
	database.SetupDatabase()
	asset.SetupRoutes(apiBasePath)
	//http.ListenAndServe("127.0.0.1:3000", nil)
	http.ListenAndServe("0.0.0.0:3000", nil)
}

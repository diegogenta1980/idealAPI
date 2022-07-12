package asset

import (
	"fmt"
	"net/http"
)

const assetsBasePath = "assets"

func SetupRoutes(basePath string) {
	handleAssets := http.HandlerFunc(assetsHandler)
	handleAsset := http.HandlerFunc(assetHandler)
	http.Handle(fmt.Sprintf("%s/%s", basePath, assetsBasePath), handleAssets)
	http.Handle(fmt.Sprintf("%s/%s/", basePath, assetsBasePath), handleAsset)
}

//Handler for the /api/assets/*, aimed to collective operations
func assetHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello API world - individual"))
}

//Handler for the /api/assets, aimed to collective operations
func assetsHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello API world - collection"))
}

package asset

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
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
	urlPathSegments := strings.Split(r.URL.Path, "assets/")
	assetID, err := strconv.Atoi(urlPathSegments[len(urlPathSegments)-1])
	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusNotFound)
		return
	}
	asset, err := getAsset(assetID)
	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if asset == nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	switch r.Method {
	case http.MethodGet:
		//return a single active
		assetJSON, err := json.Marshal(asset)
		if err != nil {
			log.Print(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(assetJSON)
	case http.MethodPut:
		//update an active in the list
		var updatedAsset Asset
		bodyBytes, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Print(err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		err = json.Unmarshal(bodyBytes, &updatedAsset)
		if err != nil {
			log.Print(err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		if updatedAsset.AssetID != assetID {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		err = updateAsset(updatedAsset)
		if err != nil {
			log.Print(err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		w.WriteHeader(http.StatusOK)
		return
		//todo delete, update
	case http.MethodDelete:
		removeAsset(assetID)
	case http.MethodOptions:
		return
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

//Handler for the /api/assets, aimed to collective operations
func assetsHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		activeList, err := getAssetsList()
		if err != nil {
			log.Print(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		activesJson, err := json.Marshal(activeList)
		if err != nil {
			log.Print(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(activesJson)
	case http.MethodPost:
		var newActive Asset
		bodyBytes, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Print(err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		err = json.Unmarshal(bodyBytes, &newActive)
		if err != nil {
			log.Print(err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		if newActive.AssetID != 0 {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		_, err = insertAsset(newActive)
		if err != nil {
			log.Print(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusCreated)
		return
	case http.MethodOptions:
		return
	}
}

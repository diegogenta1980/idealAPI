package asset

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"sort"
	"sync"
)

//mutex necessary, map not thread safe
var assetMap = struct {
	sync.RWMutex
	m map[int]Asset
}{m: make(map[int]Asset)}

func init() {
	fmt.Println("loading assets...")

	assMap, err := loadAssetMap()
	assetMap.m = assMap
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d actives loaded", len(assetMap.m))
}

func loadAssetMap() (map[int]Asset, error) {
	fileName := "assets.json"
	_, err := os.Stat(fileName)
	if os.IsNotExist(err) {
		return nil, fmt.Errorf("file [%s] does not exist", fileName)
	}

	file, _ := ioutil.ReadFile(fileName)
	activeList := make([]Asset, 0)
	err = json.Unmarshal([]byte(file), &activeList)
	if err != nil {
		log.Fatal(err)
	}
	actMap := make(map[int]Asset)
	for i := 0; i < len(activeList); i++ {
		actMap[activeList[i].AssetID] = activeList[i]
	}
	return actMap, nil
}

func getAsset(assetID int) *Asset {
	assetMap.RLock()
	defer assetMap.RUnlock() //wait until other threads release
	if active, ok := assetMap.m[assetID]; ok {
		return &active
	}
	return nil
}

func removeAsset(assetID int) {
	assetMap.Lock()
	defer assetMap.Unlock()
	delete(assetMap.m, assetID)
}

func getAssetList() []Asset {
	assetMap.RLock()
	actives := make([]Asset, 0, len(assetMap.m))
	for _, value := range assetMap.m {
		actives = append(actives, value)
	}
	defer assetMap.RUnlock()
	return actives
}

func getAssetIds() []int {
	assetMap.RLock()
	activeIds := []int{}
	for key := range assetMap.m {
		activeIds = append(activeIds, key)
	}
	assetMap.RUnlock()
	sort.Ints(activeIds)
	return activeIds
}

func getNextAssetID() int {
	activeIDs := getAssetIds()
	return activeIDs[len(activeIDs)-1] + 1
}

func addOrUpdateAsset(asset Asset) (int, error) {
	addOrUpdateID := -1
	if asset.AssetID > 0 {
		oldActive := getAsset(asset.AssetID)
		if oldActive == nil {
			return 0, fmt.Errorf("ative id [%d] does not exist", asset.AssetID)
		}
		addOrUpdateID = asset.AssetID
	} else {
		addOrUpdateID = getNextAssetID()
		asset.AssetID = addOrUpdateID
	}
	assetMap.Lock()
	assetMap.m[addOrUpdateID] = asset
	assetMap.Unlock()
	return addOrUpdateID, nil
}

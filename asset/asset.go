package asset

type Asset struct {
	AssetID int     `json:"assetId"`
	Name    string  `json:"name"`
	Price   float32 `json:"price"`
	Order   uint    `json:"order"`
}

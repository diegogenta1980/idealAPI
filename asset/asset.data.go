package asset

import (
	"database/sql"
	"fmt"

	"github.com/diegogenta1980/idealAPI/database"
)

func getAsset(assetID int) (*Asset, error) {
	userid := `7a42525d-dd03-447a-b83b-ba6b8fc16bb4` //TODO passar para parâmetro da função
	query := fmt.Sprintf(`SELECT assetid, "order", name FROM public.assets WHERE userid = '%s' AND assetid = '%d';`, userid, assetID)
	row := database.DbConnection.QueryRow(query)
	asset := &Asset{}
	err := row.Scan(&asset.AssetID, &asset.Order, &asset.Name)
	if err == sql.ErrNoRows {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return asset, nil
}

func removeAsset(assetID int) error {
	userid := `7a42525d-dd03-447a-b83b-ba6b8fc16bb4` //TODO passar para parâmetro da função
	query := fmt.Sprintf(`DELETE FROM public.assets where assetid = %d and userid = '%s'`, assetID, userid)
	_, err := database.DbConnection.Query(query)
	if err != nil {
		return err
	}
	return nil
}

func getAssetsList() ([]Asset, error) {
	userid := `7a42525d-dd03-447a-b83b-ba6b8fc16bb4` //TODO substituir pelo token, passado como parametro
	query := fmt.Sprintf(`SELECT assetid, "order", name FROM public.assets WHERE userid = '%s';`, userid)
	println(query)
	results, err := database.DbConnection.Query(query)
	if err != nil {
		return nil, err
	}
	defer results.Close()
	assets := make([]Asset, 0)
	for results.Next() {
		var asset Asset
		results.Scan(&asset.AssetID, &asset.Order, &asset.Name)
		assets = append(assets, asset)
	}
	return assets, nil
}

func updateAsset(asset Asset) error {
	userid := `7a42525d-dd03-447a-b83b-ba6b8fc16bb4` //TODO passar para parâmetro da função
	query := fmt.Sprintf(`UPDATE public.assets SET "order"=%d WHERE assetid = %d AND userid = '%s'`, asset.Order, asset.AssetID, userid)
	_, err := database.DbConnection.Exec(query)
	if err != nil {
		return err
	}
	return nil
}

func insertAsset(asset Asset) (int, error) {
	userid := `7a42525d-dd03-447a-b83b-ba6b8fc16bb4` //TODO passar para parâmetro da função
	query := fmt.Sprintf(`INSERT INTO public.assets (userid, "order", name) VALUES ('%s', %d, '%s')`, userid, asset.Order, asset.Name)
	result, err := database.DbConnection.Exec(query)
	if err != nil {
		return 0, nil
	}
	insertid, err := result.LastInsertId()
	if err != nil {
		return 0, nil
	}
	return int(insertid), nil
}

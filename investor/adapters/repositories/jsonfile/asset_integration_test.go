package jsonfile

import (
	assetEntity "investor/entities/asset"
	"investor/helpers/file"
	"testing"
)

func TestAssetRepository_Integration_ListAll(t *testing.T) {
	jsonFile := file.NewJSONFile(file.NewPlainFile(file.GetFilePath("test_list_all_assets.json")))
	repo := NewAssetRepository(NewStorage(jsonFile))

	_, err := repo.CreateBulk([]assetEntity.Asset{
		assetEntity.NewPlainAsset("1", assetEntity.PreciousMetal, "gold"),
		assetEntity.NewPlainAsset("2", assetEntity.PreciousMetal, "silver"),
		assetEntity.NewPlainAsset("3", assetEntity.PreciousMetal, "diamond"),
	})
	checkErr(t, err, "asset bulk creation")

	// test list in the same session works
	assets, err := repo.ListAll()
	checkErr(t, err, "assets list")
	if len(assets) != 3 {
		t.Errorf("3 asset expected")
	}

	// test restore from existent storage
	repo2 := NewAssetRepository(NewStorage(jsonFile))
	assets2, err := repo2.ListAll()
	checkErr(t, err, "assets list")
	if len(assets2) != 3 {
		t.Errorf("3 assets expected")
	}

	// test create works after restore (restored with first ListAll() call)
	err = repo2.Create(assetEntity.NewPlainAsset("4", assetEntity.PreciousMetal, "rubin"))
	checkErr(t, err, "asset creation")
	assets2, err = repo2.ListAll()
	checkErr(t, err, "assets list")
	if len(assets2) != 4 {
		t.Errorf("4 assets expected")
	}

	// test create work before first restoring
	repo3 := NewAssetRepository(NewStorage(jsonFile))
	err = repo3.Create(assetEntity.NewPlainAsset("5", assetEntity.PreciousMetal, "rubin"))
	checkErr(t, err, "asset creation")
	assets3, err := repo3.ListAll()
	checkErr(t, err, "assets list")
	if len(assets3) != 5 {
		t.Errorf("5 assets expected")
	}
}

package jsonfile

import (
	assetEntity "investor/entities/asset"
	"investor/helpers/file"
	"testing"
)

func TestPaymentRepository_ListAll(t *testing.T) {
	jsonFile := file.NewJsonFile(file.NewPlainFile(file.GetFilePath("test_list_all.json")))
	repo := NewAssetRepository(*NewStorage(jsonFile))

	_, err := repo.CreateBulk([]assetEntity.Asset{
		assetEntity.NewAsset("1", assetEntity.PreciousMetal, "gold"),
		assetEntity.NewAsset("2", assetEntity.PreciousMetal, "silver"),
		assetEntity.NewAsset("3", assetEntity.PreciousMetal, "diamond"),
	})
	if err != nil {
		t.Errorf("Unexpected assets creation err %+v", err)
	}

	// test list in the same session works
	assets := repo.ListAll()
	if len(assets) != 3 {
		t.Errorf("3 assets expected")
	}

	// test restore from existent storage
	repo2 := NewAssetRepository(*NewStorage(jsonFile))
	assets2 := repo2.ListAll()
	if len(assets2) != 3 {
		t.Errorf("3 assets expected")
	}

	// test create works after restore (restored with first ListAll() call)
	err = repo2.Create(assetEntity.NewAsset("4", assetEntity.PreciousMetal, "rubin"))
	if err != nil {
		t.Errorf("Unexpected assets creation err %+v", err)
	}
	assets2 = repo2.ListAll()
	if len(assets2) != 4 {
		t.Errorf("4 assets expected")
	}

	// test create work before first restoring
	repo3 := NewAssetRepository(*NewStorage(jsonFile))
	err = repo3.Create(assetEntity.NewAsset("5", assetEntity.PreciousMetal, "rubin"))
	if err != nil {
		t.Errorf("Unexpected asset creation err %+v", err)
	}
	assets3 := repo3.ListAll()
	if len(assets3) != 5 {
		t.Errorf("5 assets expected")
	}
}

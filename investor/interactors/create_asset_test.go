package interactors

import (
	"fmt"
	"investor/entities/asset"
	"testing"
)

func TestCreateAsset_Create(t *testing.T) {
	idGeneratorMock := IdGeneratorMock{GenerateFunc: func() string { return "1" }}
	assetCreatorMock := AssetCreatorMock{
		CreateFunc: func(asset asset.Asset) error { return nil },
	}
	createRequest := CreateAssetRequest{
		Name:     "test",
		Category: asset.PreciousMetal,
	}

	// test create without errors
	interactor := NewCreateAsset(assetCreatorMock, idGeneratorMock)
	response := interactor.Create(createRequest)

	expectedResponse := CreateAssetResponse{
		Created:     true,
		GeneratedId: "1",
		Err:         nil,
	}
	if response != expectedResponse {
		t.Errorf("Unepxected create result, should be created without errors")
	}

	// test create error
	mockedErr := fmt.Errorf("mocked error")
	assetCreatorMock.CreateFunc = func(asset asset.Asset) error { return mockedErr }
	interactor = NewCreateAsset(assetCreatorMock, idGeneratorMock)
	response = interactor.Create(createRequest)
	expectedResponse = CreateAssetResponse{
		Created:     false,
		GeneratedId: "1",
		Err:         mockedErr,
	}
	if response != expectedResponse {
		t.Errorf("Unepxected create result, should failed creation due to error in response")
	}
}

package interactors

import (
	assetEntity "investor/entities/asset"
)

type CreateAsset struct {
	repository  AssetCreator
	idGenerator IdGenerator
}

type CreateAssetRequest struct {
	Name     string
	Category assetEntity.Category
}

type CreateAssetResponse struct {
	Created     bool
	GeneratedId string
	Err         error
}

func (ca CreateAsset) Create(assetModel CreateAssetRequest) CreateAssetResponse {
	id := ca.idGenerator.Generate()
	p := assetEntity.NewPlainAsset(
		id, assetModel.Category, assetModel.Name,
	)

	err := ca.repository.Create(p)
	if err != nil {
		return CreateAssetResponse{
			Created:     false,
			GeneratedId: id,
			Err:         err,
		}
	} else {
		return CreateAssetResponse{
			Created:     true,
			GeneratedId: id,
			Err:         nil,
		}
	}
}

func NewCreateAsset(repository AssetCreator, idGenerator IdGenerator) CreateAsset {
	return CreateAsset{repository: repository, idGenerator: idGenerator}
}

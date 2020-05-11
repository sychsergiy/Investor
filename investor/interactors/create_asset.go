package interactors

import (
	assetEntity "investor/entities/asset"
	"investor/interactors/ports"
)

type CreateAsset struct {
	repository  ports.AssetCreator
	idGenerator ports.IDGenerator
}

type CreateAssetRequest struct {
	Name     string
	Category assetEntity.Category
}

type CreateAssetResponse struct {
	Created     bool
	GeneratedID string
	Err         error
}

func (ca CreateAsset) Create(assetModel CreateAssetRequest) CreateAssetResponse {
	id := ca.idGenerator.Generate()
	p := assetEntity.NewPlain(
		id, assetModel.Category, assetModel.Name,
	)

	err := ca.repository.Create(p)
	if err != nil {
		return CreateAssetResponse{
			Created:     false,
			GeneratedID: id,
			Err:         err,
		}
	}
	return CreateAssetResponse{
		Created:     true,
		GeneratedID: id,
		Err:         nil,
	}
}

func NewCreateAsset(repository ports.AssetCreator, idGenerator ports.IDGenerator) CreateAsset {
	return CreateAsset{repository: repository, idGenerator: idGenerator}
}

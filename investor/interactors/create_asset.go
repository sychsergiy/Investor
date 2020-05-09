package interactors

import (
	assetEntity "investor/entities/asset"
	"investor/interactors/ports"
)

type CreateAsset struct {
	repository  ports.AssetCreator
	idGenerator ports.IdGenerator
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

func NewCreateAsset(repository ports.AssetCreator, idGenerator ports.IdGenerator) CreateAsset {
	return CreateAsset{repository: repository, idGenerator: idGenerator}
}

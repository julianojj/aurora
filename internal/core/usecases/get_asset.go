package usecases

import (
	"github.com/julianojj/aurora/internal/infra/adapters"
	"go.uber.org/zap"
)

type GetAsset struct {
	bucket adapters.Bucket
	logger *zap.Logger
}

func NewGetAsset(
	bucket adapters.Bucket,
	logger *zap.Logger,
) *GetAsset {
	return &GetAsset{
		bucket,
		logger,
	}
}

func (ga *GetAsset) Execute(fileID string) ([]byte, error) {
	asset, err := ga.bucket.GetObject(fileID)
	if err != nil {
		ga.logger.Info(err.Error())
		return nil, err
	}
	return asset, nil
}

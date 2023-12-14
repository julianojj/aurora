package usecases

import (
	"github.com/julianojj/aurora/internal/core/domain"
	"github.com/julianojj/aurora/internal/core/exceptions"
)

type GetArtboard struct {
	artboardRepository domain.ArtboardRepository
}

type GetArtboardOutput struct {
	Id    string                   `json:"id"`
	Name  string                   `json:"name"`
	Layer *GetArtboardLayersOutput `json:"layer"`
}

type GetArtboardLayersOutput struct {
	LayerID string `json:"layer_id"`
	Name    string `json:"name"`
	Type    string `json:"type"`
}

func NewGetArtboard(
	artboardRepository domain.ArtboardRepository,
) *GetArtboard {
	return &GetArtboard{
		artboardRepository,
	}
}

func (ga *GetArtboard) Execute(artboardId string) (*GetArtboardOutput, error) {
	artboard, err := ga.artboardRepository.Find(artboardId)
	if err != nil {
		return nil, err
	}
	if artboard == nil {
		return nil, exceptions.NewNotFoundException("artboard not found")
	}
	return &GetArtboardOutput{
		Id:   artboard.ArtboardID,
		Name: artboard.Name,
		Layer: &GetArtboardLayersOutput{
			LayerID: artboard.Layer.LayerID,
			Name:    artboard.Layer.Name,
			Type:    artboard.Layer.LayerType,
		},
	}, nil
}

package usecases

import (
	"github.com/julianojj/aurora/internal/core/domain"
)

type GetArtboards struct {
	artboardRepository domain.ArtboardRepository
}

type GetArtboardsOutput struct {
	ProjectID  string          `json:"project_id"`
	ArtboardID string          `json:"artboard_id"`
	Name       string          `json:"name"`
	Layers     []*LayersOutput `json:"layers"`
}

type LayersOutput struct {
	LayerID string `json:"layer_id"`
	Name    string `json:"name"`
	Type    string `json:"type"`
}

func NewGetArtboards(
	artboardRepository domain.ArtboardRepository,
) *GetArtboards {
	return &GetArtboards{
		artboardRepository,
	}
}

func (ga *GetArtboards) Execute(projectID string) ([]*GetArtboardsOutput, error) {
	artboards, err := ga.artboardRepository.FindByProjectID(projectID)
	if err != nil {
		return nil, err
	}
	var output []*GetArtboardsOutput
	for _, artboard := range artboards {
		var layersOutput []*LayersOutput
		for _, layer := range artboard.Layers {
			layersOutput = append(layersOutput, &LayersOutput{
				LayerID: layer.LayerID,
				Name:    layer.Name,
				Type:    string(layer.LayerType),
			})
		}
		output = append(output, &GetArtboardsOutput{
			ProjectID:  projectID,
			ArtboardID: artboard.ArtboardID,
			Name:       artboard.Name,
			Layers:     layersOutput,
		})
	}
	return output, nil
}

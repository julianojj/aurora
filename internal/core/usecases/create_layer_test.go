package usecases

import (
	"testing"

	"github.com/julianojj/aurora/internal/core/domain"
	"github.com/julianojj/aurora/internal/infra/repository"
	"github.com/stretchr/testify/assert"
)

func TestShouldCreateLayer(t *testing.T) {
	projectRepository := repository.NewProjectRepositoryMemory()
	artboardRepository := repository.NewArtboardRepositoryMemory()
	layerRepository := repository.NewLayerRepositoryMemory()
	createProject := NewCreateProject(projectRepository)
	createArtboard := NewCreateArtboard(projectRepository, artboardRepository)
	inputCreateProject := CreateProjectInput{
		Name: "Untitled Project",
	}
	output, _ := createProject.Execute(inputCreateProject)
	inputCreateArtboard := CreateArtboardInput{
		ProjectID: output.ProjectID,
		Name:      "My Artboard",
	}
	err := createArtboard.Execute(inputCreateArtboard)
	assert.NoError(t, err)
	size := &Size{
		Width:  100,
		Height: 100,
	}
	position := &Position{
		X: 0,
		Y: 0,
	}
	rotation := &Rotation{
		Angle: 90,
	}
	properties := &Properties{
		FillColor:   "#FFF",
		StrokeColor: "#000",
		StrokeWidth: 1,
		Opacity:     100,
	}
	input := CreateLayerInput{
		ArtboardID: artboardRepository.Artboards[0].ArtboardID,
		LayerName:  "test",
		LayerType:  domain.Reactangle,
		Size:       size,
		Position:   position,
		Rotation:   rotation,
		Properties: properties,
	}
	createLayer := NewCreateLayer(artboardRepository, layerRepository)
	err = createLayer.Execute(input)
	assert.NoError(t, err)
	assert.Len(t, artboardRepository.Artboards[0].Layers, 1)
	assert.Equal(t, artboardRepository.Artboards[0].Layers[0].Name, "test")
	assert.Equal(t, artboardRepository.Artboards[0].ArtboardID, input.ArtboardID)
}

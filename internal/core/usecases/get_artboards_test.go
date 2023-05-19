package usecases

import (
	"testing"

	"github.com/julianojj/aurora/internal/infra/repository"
	"github.com/stretchr/testify/assert"
)

func TestGetArtboards(t *testing.T) {
	projectRepository := repository.NewProjectRepositoryMemory()
	artboardRepository := repository.NewArtboardRepositoryMemory()
	createProject := NewCreateProject(projectRepository)
	createArtboard := NewCreateArtboard(projectRepository, artboardRepository)
	getArtboards := NewGetArtboards(artboardRepository)
	inputCreateProject := CreateProjectInput{
		Name: "Untitled Project",
	}
	outputCreateProject, _ := createProject.Execute(inputCreateProject)
	inputCreateArtboard := CreateArtboardInput{
		ProjectID: outputCreateProject.ProjectID,
		Name:      "My Artboard",
	}
	err := createArtboard.Execute(inputCreateArtboard)
	artboards, _ := getArtboards.Execute(outputCreateProject.ProjectID)
	assert.NoError(t, err)
	assert.Equal(t, 1, len(artboards))
	assert.Equal(t, "My Artboard", artboards[0].Name)
}

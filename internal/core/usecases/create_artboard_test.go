package usecases

import (
	"testing"

	"github.com/julianojj/aurora/internal/infra/repository"
	"github.com/stretchr/testify/assert"
)

func TestShouldCreateDesign(t *testing.T) {
	projectRepository := repository.NewProjectRepositoryMemory()
	artboardRepository := repository.NewArtboardRepositoryMemory()
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
	artboards := artboardRepository.Artboards
	assert.Len(t, artboards, 1)
	assert.Equal(t, inputCreateArtboard.Name, artboards[0].Name)
}

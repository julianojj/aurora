package usecases

import (
	"testing"

	"github.com/julianojj/aurora/internal/infra/repository"
	"github.com/stretchr/testify/assert"
)

func TestReturnErrorIfProjectNotFound(t *testing.T) {
	projectRepository := repository.NewProjectRepositoryMemory()
	artboardRepository := repository.NewArtboardRepositoryMemory()
	createArtboard := NewCreateArtboard(projectRepository, artboardRepository)
	inputCreateArtboard := CreateArtboardInput{
		ProjectID: "1",
		Name:      "My Artboard",
	}
	err := createArtboard.Execute(inputCreateArtboard)
	assert.EqualError(t, err, "project not found")
}

func TestReturnErrorIfEmptyArtboardName(t *testing.T) {
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
		Name:      "",
	}
	err := createArtboard.Execute(inputCreateArtboard)
	assert.EqualError(t, err, "Name cannot be empty")
}

func TestShouldCreateArtboard(t *testing.T) {
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

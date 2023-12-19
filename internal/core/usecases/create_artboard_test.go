package usecases

import (
	"testing"

	"github.com/julianojj/aurora/internal/infra/repository"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
)

func TestReturnErrorIfProjectNotFound(t *testing.T) {
	projectRepository := repository.NewProjectRepositoryMemory()
	artboardRepository := repository.NewArtboardRepositoryMemory()
	logger, _ := zap.NewProduction()
	createArtboard := NewCreateArtboard(projectRepository, artboardRepository, logger)
	inputCreateArtboard := CreateArtboardInput{
		ProjectID: "1",
		Name:      "My Artboard",
	}
	err := createArtboard.Execute(inputCreateArtboard)
	assert.EqualError(t, err, "project not found")
}

func TestShouldCreateArtboard(t *testing.T) {
	projectRepository := repository.NewProjectRepositoryMemory()
	artboardRepository := repository.NewArtboardRepositoryMemory()
	logger, _ := zap.NewProduction()
	createProject := NewCreateProject(projectRepository, logger)
	createArtboard := NewCreateArtboard(projectRepository, artboardRepository, logger)
	inputCreateProject := CreateProjectInput{
		Name: "Untitled Project",
	}
	output, _ := createProject.Execute(inputCreateProject)
	inputCreateArtboard := CreateArtboardInput{
		ProjectID: output.ProjectID,
		Name:      "My Artboard",
	}
	err := createArtboard.Execute(inputCreateArtboard)
	artboards := artboardRepository.Artboards
	assert.NoError(t, err)
	assert.Len(t, artboards, 1)
	assert.Equal(t, inputCreateArtboard.Name, artboards[0].Name)
}

func TestReturnErrorIfErrorToSaveArtboard(t *testing.T) {
	projectRepository := repository.NewProjectRepositoryMemory()
	artboardRepository := repository.NewArtboardRepositoryMemory()
	logger, _ := zap.NewProduction()
	createProject := NewCreateProject(projectRepository, logger)
	projectRepository.MockFind()
	createArtboard := NewCreateArtboard(projectRepository, artboardRepository, logger)
	inputCreateProject := CreateProjectInput{
		Name: "Untitled Project",
	}
	output, _ := createProject.Execute(inputCreateProject)
	inputCreateArtboard := CreateArtboardInput{
		ProjectID: output.ProjectID,
		Name:      "My Artboard",
	}
	err := createArtboard.Execute(inputCreateArtboard)
	assert.EqualError(t, err, "project not found")
}

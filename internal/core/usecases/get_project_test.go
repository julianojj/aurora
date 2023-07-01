package usecases

import (
	"testing"

	"github.com/julianojj/aurora/internal/infra/repository"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
)

func TestReturnExceptionIfProjectNotFound(t *testing.T) {
	projectRepository := repository.NewProjectRepositoryMemory()
	getProject := NewGetProject(projectRepository)
	_, err := getProject.Execute("1")
	assert.EqualError(t, err, "Project not found")
}

func TestGetProject(t *testing.T) {
	logger, _ := zap.NewProduction()
	projectRepository := repository.NewProjectRepositoryMemory()
	createProject := NewCreateProject(projectRepository, logger)
	getProject := NewGetProject(projectRepository)
	input := CreateProjectInput{
		Name: "Test",
	}
	output, _ := createProject.Execute(input)
	project, _ := getProject.Execute(output.ProjectID)
	assert.Equal(t, "Test", project.Name)
}

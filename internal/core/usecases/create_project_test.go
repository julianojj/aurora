package usecases

import (
	"testing"

	"github.com/julianojj/aurora/internal/infra/repository"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
)

func TestShouldCreateProject(t *testing.T) {
	projectRepository := repository.NewProjectRepositoryMemory()
	logger, _ := zap.NewProduction()
	createProject := NewCreateProject(projectRepository, logger)
	input := CreateProjectInput{
		Name: "Untitled Project",
	}
	output, _ := createProject.Execute(input)
	project, _ := projectRepository.Find(output.ProjectID)
	assert.Equal(t, "Untitled Project", project.Name)
}

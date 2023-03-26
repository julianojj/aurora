package usecases

import (
	"testing"

	"github.com/julianojj/aurora/internal/infra/repository"
	"github.com/stretchr/testify/assert"
)

func TestShouldCreateProject(t *testing.T) {
	projectRepository := repository.NewProjectRepositoryMemory()
	createProject := NewCreateProject(projectRepository)
	input := CreateProjectInput{
		Name: "Untitled Project",
	}
	output, _ := createProject.Execute(input)
	project, _ := projectRepository.Find(output.ProjectID)
	assert.Equal(t, "Untitled Project", project.Name)
}

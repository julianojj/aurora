package usecases

import (
	"testing"

	"github.com/julianojj/aurora/internal/infra/repository"
	"github.com/stretchr/testify/assert"
)

func TestShouldCreateProject(t *testing.T) {
	projectRepository := repository.NewProjectRepositoryMemory()
	createProject := NewCreateProject(projectRepository)
	input := &CreateProjectInput{
		Name: "Untitled Project",
	}
	err := createProject.Execute(*input)
	assert.NoError(t, err)
	projects, err := projectRepository.FindAll()
	assert.NoError(t, err)
	assert.Len(t, projects, 1)
	assert.Equal(t, "Untitled Project", projects[0].Name)
}

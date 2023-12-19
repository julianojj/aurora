package usecases

import (
	"testing"

	"github.com/julianojj/aurora/internal/core/exceptions"
	"github.com/julianojj/aurora/internal/infra/repository"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
)

func TestCreateProject(t *testing.T) {
	tests := []struct {
		name string
		fn   func(t *testing.T)
	}{
		{
			name: "should return no error when create project with valid data",
			fn: func(t *testing.T) {
				projectRepository := repository.NewProjectRepositoryMemory()
				logger, _ := zap.NewProduction()
				createProject := NewCreateProject(projectRepository, logger)
				input := CreateProjectInput{
					Name: "Untitled Project",
				}
				output, _ := createProject.Execute(input)
				project, _ := projectRepository.Find(output.ProjectID)
				assert.Equal(t, "Untitled Project", project.Name)
			},
		},
		{
			name: "should return error when create project with error to save",
			fn: func(t *testing.T) {
				projectRepository := repository.NewProjectRepositoryMemory().MockSave()
				logger, _ := zap.NewProduction()
				createProject := NewCreateProject(projectRepository, logger)
				input := CreateProjectInput{
					Name: "Untitled Project",
				}
				_, err := createProject.Execute(input)
				assert.EqualError(t, err, "error to save")
			},
		},
		{
			name: "should return error when create project with error to save",
			fn: func(t *testing.T) {
				projectRepository := repository.NewProjectRepositoryMemory()
				logger, _ := zap.NewProduction()
				createProject := NewCreateProject(projectRepository, logger)
				input := CreateProjectInput{
					Name: "",
				}
				_, err := createProject.Execute(input)
				assert.EqualError(t, err, exceptions.EMPTY_PROJECT_NAME)
			},
		},
	}
	for _, test := range tests {
		t.Run(test.name, test.fn)
	}
}

package usecases

import (
	"time"

	"github.com/google/uuid"
	"github.com/julianojj/aurora/internal/core/domain"
)

type CreateProject struct {
	ProjectRepository domain.ProjectRepository
}

type CreateProjectInput struct {
	Name string
}

type CreateProjectOutput struct {
	ProjectID string `json:"project_id"`
}

func NewCreateProject(projectRepository domain.ProjectRepository) *CreateProject {
	return &CreateProject{
		ProjectRepository: projectRepository,
	}
}

func (c *CreateProject) Execute(input CreateProjectInput) (*CreateProjectOutput, error) {
	today := time.Now()
	project, err := domain.NewProject(
		uuid.NewString(),
		input.Name,
		today,
		today,
	)
	if err != nil {
		return nil, err
	}
	err = c.ProjectRepository.Save(project)
	if err != nil {
		return nil, err
	}
	output := &CreateProjectOutput{
		ProjectID: project.ID,
	}
	return output, nil
}

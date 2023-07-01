package usecases

import (
	"time"

	"github.com/google/uuid"
	"github.com/julianojj/aurora/internal/core/domain"
	"go.uber.org/zap"
)

type CreateProject struct {
	ProjectRepository domain.ProjectRepository
	Logger            *zap.Logger
}

type CreateProjectInput struct {
	Name string
}

type CreateProjectOutput struct {
	ProjectID string `json:"project_id"`
}

func NewCreateProject(
	projectRepository domain.ProjectRepository,
	logger *zap.Logger,
) *CreateProject {
	return &CreateProject{
		ProjectRepository: projectRepository,
		Logger:            logger,
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
		c.Logger.Info(err.Error())
		return nil, err
	}
	err = c.ProjectRepository.Save(project)
	if err != nil {
		c.Logger.Info(err.Error())
		return nil, err
	}
	output := &CreateProjectOutput{
		ProjectID: project.ID,
	}
	c.Logger.Info("created.project")
	return output, nil
}

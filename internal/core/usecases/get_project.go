package usecases

import (
	"github.com/julianojj/aurora/internal/core/domain"
	"github.com/julianojj/aurora/internal/core/exceptions"
)

type GetProject struct {
	projectRepository domain.ProjectRepository
}

type GetProjectOutput struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

func NewGetProject(projectRepository domain.ProjectRepository) *GetProject {
	return &GetProject{
		projectRepository,
	}
}

func (gp *GetProject) Execute(projectID string) (*GetProjectOutput, error) {
	existingProject, err := gp.projectRepository.Find(projectID)
	if err != nil {
		return nil, err
	}
	if existingProject == nil {
		return nil, exceptions.NewNotFoundException("Project not found")
	}
	return &GetProjectOutput{
		Id:   existingProject.ID,
		Name: existingProject.Name,
	}, nil
}

package repository

import "github.com/julianojj/aurora/internal/core/domain"

type ProjectRepositoryMemory struct {
	Projects []*domain.Project
}

func NewProjectRepositoryMemory() *ProjectRepositoryMemory {
	return &ProjectRepositoryMemory{
		Projects: make([]*domain.Project, 0),
	}
}

func (r *ProjectRepositoryMemory) Save(project *domain.Project) error {
	r.Projects = append(r.Projects, project)
	return nil
}

func (r *ProjectRepositoryMemory) FindAll() ([]*domain.Project, error) {
	return r.Projects, nil
}

func (r *ProjectRepositoryMemory) Find(projectID string) (*domain.Project, error) {
	for _, project := range r.Projects {
        if project.ID == projectID {
            return project, nil
        }
    }
    return nil, nil
}

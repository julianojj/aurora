package repository

import (
	"errors"

	"github.com/julianojj/aurora/internal/core/domain"
)

type ProjectRepositoryMemory struct {
	isMockSave    bool
	isMockFindAll bool
	isMockFind    bool
	Projects      []*domain.Project
}

func NewProjectRepositoryMemory() *ProjectRepositoryMemory {
	return &ProjectRepositoryMemory{
		Projects: make([]*domain.Project, 0),
	}
}

func (r *ProjectRepositoryMemory) Save(project *domain.Project) error {
	if r.isMockSave {
		return errors.New("error to save")
	}
	r.Projects = append(r.Projects, project)
	return nil
}

func (r *ProjectRepositoryMemory) FindAll() ([]*domain.Project, error) {
	return r.Projects, nil
}

func (r *ProjectRepositoryMemory) Find(projectID string) (*domain.Project, error) {
	if r.isMockFind {
		return nil, errors.New("project not found")
	}
	for _, project := range r.Projects {
		if project.ID == projectID {
			return project, nil
		}
	}
	return nil, nil
}

func (r *ProjectRepositoryMemory) MockSave() *ProjectRepositoryMemory {
	r.isMockSave = true
	return r
}

func (r *ProjectRepositoryMemory) MockFind() *ProjectRepositoryMemory {
	r.isMockFind = true
	return r
}

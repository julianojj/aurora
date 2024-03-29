package domain

import (
	"time"

	"github.com/julianojj/aurora/internal/core/exceptions"
)

type Project struct {
	ID        string
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewProject(
	id, name string,
	createdAt, updatedAt time.Time,
) (*Project, error) {
	project := &Project{
		ID:        id,
		Name:      name,
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
	}
	err := project.Validate()
	if err != nil {
		return nil, err
	}
	return project, nil
}

func (p *Project) Validate() error {
	if p.ID == "" {
		return exceptions.NewValidationException(exceptions.EMPTY_PROJECT_ID)
	}
	if p.Name == "" {
		return exceptions.NewValidationException(exceptions.EMPTY_PROJECT_NAME)
	}
	return nil
}

package usecases

import (
	"github.com/google/uuid"
	"github.com/julianojj/aurora/internal/core/domain"
	"github.com/julianojj/aurora/internal/core/exceptions"
)

type CreateArtboard struct {
	projectRepository  domain.ProjectRepository
	artboardRepository domain.ArtboardRepository
}

type CreateArtboardInput struct {
	ProjectID string `json:"project_id"`
	Name      string `json:"name"`
}

func NewCreateArtboard(
	projectRepository domain.ProjectRepository,
	artboardRepository domain.ArtboardRepository,
) *CreateArtboard {
	return &CreateArtboard{
		projectRepository,
		artboardRepository,
	}
}

func (c *CreateArtboard) Execute(input CreateArtboardInput) error {
	existingProject, err := c.projectRepository.Find(input.ProjectID)
	if err != nil {
		return err
	}
	if existingProject == nil {
		return exceptions.NewNotFoundException("project not found")
	}
	artboard, err := domain.NewArtboard(
		uuid.NewString(),
		existingProject.ID,
		input.Name,
		&domain.Layer{},
	)
	if err != nil {
		return err
	}
	err = c.artboardRepository.Save(artboard)
	if err != nil {
		return err
	}
	return nil
}

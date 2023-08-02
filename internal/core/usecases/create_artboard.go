package usecases

import (
	"github.com/google/uuid"
	"github.com/julianojj/aurora/internal/core/domain"
	"github.com/julianojj/aurora/internal/core/exceptions"
	"go.uber.org/zap"
)

type CreateArtboard struct {
	projectRepository  domain.ProjectRepository
	artboardRepository domain.ArtboardRepository
	logger             *zap.Logger
}

type CreateArtboardInput struct {
	ProjectID string `json:"project_id"`
	Name      string `json:"name"`
}

func NewCreateArtboard(
	projectRepository domain.ProjectRepository,
	artboardRepository domain.ArtboardRepository,
	logger *zap.Logger,
) *CreateArtboard {
	return &CreateArtboard{
		projectRepository,
		artboardRepository,
		logger,
	}
}

func (c *CreateArtboard) Execute(input CreateArtboardInput) error {
	existingProject, err := c.projectRepository.Find(input.ProjectID)
	if err != nil {
		c.logger.Info(err.Error())
		return err
	}
	if existingProject == nil {
		c.logger.Info("project not found")
		return exceptions.NewNotFoundException("project not found")
	}
	artboard, err := domain.NewArtboard(
		uuid.NewString(),
		existingProject.ID,
		input.Name,
		&domain.Layer{},
	)
	if err != nil {
		c.logger.Info(err.Error())
		return err
	}
	err = c.artboardRepository.Save(artboard)
	if err != nil {
		c.logger.Info(err.Error())
		return err
	}
	c.logger.Info("created artboard")
	return nil
}

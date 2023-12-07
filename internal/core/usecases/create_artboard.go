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
	size := &domain.Size{
		Width:  100,
		Height: 100,
	}
	position := &domain.Position{
		X: 0,
		Y: 0,
	}
	rotation := &domain.Rotation{
		Angle: 0,
	}
	properties := &domain.Properties{
		Size:        size,
		Position:    position,
		Rotation:    rotation,
		FillColor:   "#FFF",
		StrokeColor: "#000",
		StrokeWidth: 1,
		Opacity:     100,
	}
	artboardId := uuid.NewString()
	layer, err := domain.NewLayer(
		uuid.NewString(),
		artboardId,
		"Main",
		"layer",
		properties,
	)
	if err != nil {
		return err
	}
	artboard, err := domain.NewArtboard(
		artboardId,
		existingProject.ID,
		input.Name,
		layer,
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
	c.logger.Info("created.artboard", zap.Any("log", artboard))
	return nil
}

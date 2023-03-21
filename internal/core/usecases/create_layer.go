package usecases

import (
	"github.com/google/uuid"
	"github.com/julianojj/aurora/internal/core/domain"
	"github.com/julianojj/aurora/internal/core/exceptions"
)

type CreateLayer struct {
	artboardRepository domain.ArtboardRepository
	layerRepository    domain.LayerRepository
}

type CreateLayerInput struct {
	ArtboardID string           `json:"artboard_id"`
	LayerName  string           `json:"layer_name"`
	LayerType  domain.LayerType `json:"layer_type"`
	Size       *Size
	Position   *Position
	Rotation   *Rotation
	Properties *Properties
}

type Size struct {
	Width  float64 `json:"width"`
	Height float64 `json:"height"`
}

type Position struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
}

type Rotation struct {
	Angle float64 `json:"angle"`
}

type Properties struct {
	FillColor   string
	StrokeColor string
	StrokeWidth float64
	Opacity     float64
}

func NewCreateLayer(
	artboardRepository domain.ArtboardRepository,
	layerRepository domain.LayerRepository,
) *CreateLayer {
	return &CreateLayer{
		artboardRepository,
		layerRepository,
	}
}

func (cl *CreateLayer) Execute(input CreateLayerInput) error {
	existingArtboard, err := cl.artboardRepository.Find(input.ArtboardID)
	if err != nil {
		return err
	}
	if existingArtboard == nil {
		return exceptions.NewNotFoundException("Artboard not found")
	}
	size := &domain.Size{
		Width:  input.Size.Width,
		Height: input.Size.Height,
	}
	position := &domain.Position{
		X: input.Position.X,
		Y: input.Position.Y,
	}
	rotation := &domain.Rotation{
		Angle: input.Rotation.Angle,
	}
	properties := &domain.Properties{
		FillColor:   input.Properties.FillColor,
		StrokeColor: input.Properties.StrokeColor,
		StrokeWidth: input.Properties.StrokeWidth,
		Opacity:     input.Properties.Opacity,
	}
	layer, err := domain.NewLayer(
		uuid.NewString(),
		input.ArtboardID,
		input.LayerName,
		input.LayerType,
		size,
		position,
		rotation,
		properties,
	)
	if err != nil {
		return err
	}
	existingArtboard.AddLayer(layer)
	return cl.layerRepository.Save(layer)
}

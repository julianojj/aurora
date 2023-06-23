package domain

import (
	"github.com/julianojj/aurora/internal/core/exceptions"
)

type Layer struct {
	LayerID    string
	ArtboardID string
	Name       string
	LayerType  string
	Size       *Size
	Position   *Position
	Rotation   *Rotation
	Properties *Properties
	Children   []any
}

type Size struct {
	Width  float64
	Height float64
}

type Position struct {
	X float64
	Y float64
}

type Rotation struct {
	Angle float64
}

type Properties struct {
	Size        *Size
	Position    *Position
	Rotation    *Rotation
	FillColor   string
	StrokeColor string
	StrokeWidth float64
	Opacity     float64
}

func NewLayer(
	layerID string,
	artboardID string,
	name string,
	layerType string,
	properties *Properties,
) (*Layer, error) {
	layer := &Layer{
		LayerID:    layerID,
		ArtboardID: artboardID,
		Name:       name,
		LayerType:  layerType,
		Properties: properties,
	}
	err := layer.Validate()
	if err != nil {
		return nil, err
	}
	return layer, nil
}

func (l *Layer) Validate() error {
	if l.LayerID == "" {
		return exceptions.NewValidationException("Layer ID cannot be empty")
	}
	if l.ArtboardID == "" {
		return exceptions.NewValidationException("Artboard ID cannot be empty")
	}
	if l.Name == "" {
		return exceptions.NewValidationException("Name cannot be empty")
	}
	if l.LayerType == "" {
		return exceptions.NewValidationException("Layer Type cannot be empty")
	}
	if l.Properties == nil {
		return exceptions.NewValidationException("Properties cannot be empty")
	}
	if l.Properties.Size == nil {
		return exceptions.NewValidationException("Size cannot be empty")
	}
	if l.Properties.Position == nil {
		return exceptions.NewValidationException("Position cannot be empty")
	}
	if l.Properties.Rotation == nil {
		return exceptions.NewValidationException("Rotation cannot be empty")
	}
	return nil
}

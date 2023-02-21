package domain

import (
	"github.com/julianojj/aurora/internal/core/exceptions"
)

type Artboard struct {
	ArtboardID string
	ProjectID  string
	Name       string
	Layers     []*Layer
}

func NewArtboard(id string, projectID string, name string) (*Artboard, error) {
	artboard := &Artboard{
		ArtboardID: id,
		ProjectID:  projectID,
		Name:       name,
		Layers:     make([]*Layer, 0),
	}
	err := artboard.Validate()
	if err != nil {
		return nil, err
	}
	return artboard, nil
}

func (a *Artboard) Validate() error {
	if a.ArtboardID == "" {
		return exceptions.NewValidationException("Artboard ID cannot be empty")
	}
	if a.ProjectID == "" {
		return exceptions.NewValidationException("Project ID cannot be empty")
	}
	if a.Name == "" {
		return exceptions.NewValidationException("Name cannot be empty")
	}
	return nil
}

func (a *Artboard) AddLayer(layer *Layer) {
	a.Layers = append(a.Layers, layer)
}

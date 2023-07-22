package domain

import (
	"github.com/julianojj/aurora/internal/core/exceptions"
)

type Artboard struct {
	ArtboardID string
	ProjectID  string
	Name       string
	Layer      *Layer
}

func NewArtboard(id string, projectID string, name string, layer *Layer) (*Artboard, error) {
	artboard := &Artboard{
		ArtboardID: id,
		ProjectID:  projectID,
		Name:       name,
		Layer:      layer,
	}
	err := artboard.Validate()
	if err != nil {
		return nil, err
	}
	return artboard, nil
}

func (a *Artboard) Validate() error {
	if a.ArtboardID == "" {
		return exceptions.NewValidationException(exceptions.EMPTY_ARTBOARD_ID)
	}
	if a.ProjectID == "" {
		return exceptions.NewValidationException(exceptions.EMPTY_PROJECT_ID)
	}
	if a.Name == "" {
		return exceptions.NewValidationException(exceptions.EMPTY_ARTBOARD_NAME)
	}
	return nil
}

func (a *Artboard) AddChildren(children any) {
	a.Layer.Children = append(a.Layer.Children, children)
}

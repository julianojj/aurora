package repository

import "github.com/julianojj/aurora/internal/core/domain"

type ArtboardRepositoryMemory struct {
	Artboards []*domain.Artboard
}

func NewArtboardRepositoryMemory() *ArtboardRepositoryMemory {
	return &ArtboardRepositoryMemory{
		Artboards: make([]*domain.Artboard, 0),
	}
}

func (a *ArtboardRepositoryMemory) Save(artboard *domain.Artboard) error {
	a.Artboards = append(a.Artboards, artboard)
	return nil
}

func (a *ArtboardRepositoryMemory) Find(artboardID string) (*domain.Artboard, error) {
	for _, artboard := range a.Artboards {
		if artboard.ArtboardID == artboardID {
			return artboard, nil
		}
	}
	return nil, nil
}

func (a *ArtboardRepositoryMemory) FindByProjectID(projectID string) ([]*domain.Artboard, error) {
	var artboards []*domain.Artboard
	for _, artboard := range a.Artboards {
		if artboard.ProjectID == projectID {
			artboards = append(artboards, artboard)
		}
	}
	return artboards, nil
}

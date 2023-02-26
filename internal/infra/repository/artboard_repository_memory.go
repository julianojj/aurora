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

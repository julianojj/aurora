package repository

import "github.com/julianojj/aurora/internal/core/domain"

type LayerRepositoryMemory struct {
	Layers []*domain.Layer
}

func NewLayerRepositoryMemory() *LayerRepositoryMemory {
	return &LayerRepositoryMemory{
		Layers: []*domain.Layer{},
	}
}

func (r *LayerRepositoryMemory) Save(layer *domain.Layer) error {
	r.Layers = append(r.Layers, layer)
	return nil
}

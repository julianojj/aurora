package domain

type LayerRepository interface {
	Save(layer *Layer) error
}

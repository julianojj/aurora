package domain

type ArtboardRepository interface {
	Save(artboard *Artboard) error
}

package domain

type ArtboardRepository interface {
	Save(artboard *Artboard) error
	Find(artboardID string) (*Artboard, error)
}

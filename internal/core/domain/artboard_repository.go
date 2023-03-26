package domain

type ArtboardRepository interface {
	Save(artboard *Artboard) error
	Find(artboardID string) (*Artboard, error)
	FindByProjectID(projectID string) ([]*Artboard, error)
}

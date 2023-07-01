package domain

type FileRepository interface {
	Save(file *File) error
	FindAll() ([]*File, error)
	Find(fileId string) (*File, error)
	Delete(fileID string) error
}

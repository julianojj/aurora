package domain

type FileRepository interface {
	Save(file *File) error
	FindAll() ([]*File, error)
}

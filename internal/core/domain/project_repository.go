package domain

type ProjectRepository interface {
	Save(project *Project) error
	FindAll() ([]*Project, error)
	Find(projectID string) (*Project, error)
}

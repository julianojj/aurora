package domain

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

const ProjectName = "My Project"

func TestNotShouldCreateProjectIfEmptyProjectID(t *testing.T) {
	today := time.Now()
	_, err := NewProject("", ProjectName, today, today)
	assert.EqualError(t, err, "Project ID cannot be empty")
}

func TestNotShouldCreateProjectIfEmptyName(t *testing.T) {
	today := time.Now()
	_, err := NewProject("1", "", today, today)
	assert.EqualError(t, err, "Name cannot be empty")
}

func TestShouldCreateProject(t *testing.T) {
	today := time.Now()
	project, _ := NewProject("1", ProjectName, today, today)
	assert.Equal(t, "1", project.ID)
	assert.Equal(t, ProjectName, project.Name)
	assert.Equal(t, today, project.CreatedAt)
	assert.Equal(t, today, project.UpdatedAt)
}

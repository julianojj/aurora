package domain

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNotShouldCreateProjectIfEmptyProjectID(t *testing.T) {
	today := time.Now()
	project, err := NewProject("", "My Project", today, today)
	assert.EqualError(t, err, "Project ID cannot be empty")
	assert.Nil(t, project)
}

func TestNotShouldCreateProjectIfEmptyName(t *testing.T) {
	today := time.Now()
	project, err := NewProject("1", "", today, today)
	assert.EqualError(t, err, "Name cannot be empty")
	assert.Nil(t, project)
}

func TestShouldCreateProject(t *testing.T) {
	today := time.Now()
	project, err := NewProject("1", "My Project", today, today)
	assert.NoError(t, err)
	assert.Equal(t, "1", project.ID)
	assert.Equal(t, "My Project", project.Name)
	assert.Equal(t, today, project.CreatedAt)
	assert.Equal(t, today, project.UpdatedAt)
}

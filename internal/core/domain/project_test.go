package domain

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNotShouldCreateProjectIfEmptyProjectID(t *testing.T) {
	today := time.Now()
	_, err := NewProject("", "My Project", today, today)
	assert.EqualError(t, err, "Project ID cannot be empty")
}

func TestNotShouldCreateProjectIfEmptyName(t *testing.T) {
	today := time.Now()
	_, err := NewProject("1", "", today, today)
	assert.EqualError(t, err, "Name cannot be empty")
}

func TestShouldCreateProject(t *testing.T) {
	today := time.Now()
	project, _ := NewProject("1", "My Project", today, today)
	assert.Equal(t, "1", project.ID)
	assert.Equal(t, "My Project", project.Name)
	assert.Equal(t, today, project.CreatedAt)
	assert.Equal(t, today, project.UpdatedAt)
}

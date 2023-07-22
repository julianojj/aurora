package domain

import (
	"testing"
	"time"

	"github.com/julianojj/aurora/internal/core/exceptions"
	"github.com/stretchr/testify/assert"
)

const ProjectName = "My Project"

func TestProject(t *testing.T) {
	today := time.Now()
	project, _ := NewProject("1", "My Project", today, today)

	assert.Equal(t, "1", project.ID)
	assert.Equal(t, "My Project", project.Name)
	assert.Equal(t, today, project.CreatedAt)
	assert.Equal(t, today, project.UpdatedAt)

	t.Run("Return exception if empty project id", func(t *testing.T) {
		_, err := NewProject("", "My Project", today, today)
		assert.EqualError(t, err, exceptions.EMPTY_PROJECT_ID)
	})

	t.Run("Return exception if empty project name", func(t *testing.T) {
		_, err := NewProject("1", "", today, today)
		assert.EqualError(t, err, exceptions.EMPTY_PROJECT_NAME)
	})
}

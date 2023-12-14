package usecases

import (
	"testing"

	"github.com/julianojj/aurora/internal/core/domain"
	"github.com/julianojj/aurora/internal/infra/repository"
	"github.com/stretchr/testify/assert"
)

func TestGetArtboard(t *testing.T) {
	tests := []struct {
		name string
		fn   func(t *testing.T)
	}{
		{
			name: "should return no error when valid get artboard",
			fn: func(t *testing.T) {
				artboardRepository := repository.NewArtboardRepositoryMemory()
				artboard, _ := domain.NewArtboard("1", "1", "Test", &domain.Layer{})
				artboardRepository.Save(artboard)
				getArtboard := NewGetArtboard(artboardRepository)
				output, err := getArtboard.Execute("1")
				assert.NoError(t, err)
				assert.Equal(t, "1", output.Id)
				assert.Equal(t, "Test", output.Name)
				assert.Equal(t, &GetArtboardLayersOutput{}, output.Layer)
			},
		},
		{
			name: "should return error when artboard not found",
			fn: func(t *testing.T) {
				artboardRepository := repository.NewArtboardRepositoryMemory()
				getArtboard := NewGetArtboard(artboardRepository)
				_, err := getArtboard.Execute("1")
				assert.EqualError(t, err, "artboard not found")
			},
		},
	}
	for _, test := range tests {
		t.Run(test.name, test.fn)
	}
}

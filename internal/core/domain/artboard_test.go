package domain

import (
	"testing"

	"github.com/julianojj/aurora/internal/core/exceptions"
	"github.com/stretchr/testify/assert"
)

func TestArtboard(t *testing.T) {
	size := &Size{
		Width:  100,
		Height: 100,
	}
	position := &Position{
		X: 0,
		Y: 0,
	}
	rotation := &Rotation{
		Angle: 90,
	}
	properties := &Properties{
		Size:        size,
		Position:    position,
		Rotation:    rotation,
		FillColor:   "#FFF",
		StrokeColor: "#000",
		StrokeWidth: 1,
		Opacity:     100,
	}
	layer, _ := NewLayer(
		"1",
		"1",
		"My Layer",
		"layer",
		properties,
	)
	artboard, _ := NewArtboard(
		"1",
		"1",
		"My Artboard",
		layer,
	)

	assert.Equal(t, "1", artboard.ArtboardID)
	assert.Equal(t, "1", artboard.ProjectID)
	assert.Equal(t, "My Artboard", artboard.Name)
	assert.Equal(t, layer, artboard.Layer)

	t.Run("Return exception if empty artboard id", func(t *testing.T) {
		_, err := NewArtboard("", "1", "My Artboard", &Layer{})
		assert.EqualError(t, err, exceptions.EMPTY_ARTBOARD_ID)
	})

	t.Run("Return exception if empty project id", func(t *testing.T) {
		_, err := NewArtboard("1", "", "My Artboard", &Layer{})
		assert.EqualError(t, err, exceptions.EMPTY_PROJECT_ID)
	})

	t.Run("Return exception if empty artboard name", func(t *testing.T) {
		_, err := NewArtboard("1", "1", "", &Layer{})
		assert.EqualError(t, err, exceptions.EMPTY_ARTBOARD_NAME)
	})

	t.Run("Add children", func(t *testing.T) {
		layer, _ := NewLayer("1", artboard.ArtboardID, "test", "Text", &Properties{
			Size: &Size{
				Width:  100,
				Height: 100,
			},
			Position: &Position{
				X: 0,
				Y: 0,
			},
			Rotation: &Rotation{
				Angle: 0,
			},
			FillColor:   "#000",
			StrokeColor: "#000",
			StrokeWidth: 1,
			Opacity:     100,
		})
		artboard.AddChildren(layer)
		assert.Len(t, artboard.Layer.Children, 1)
	})
}

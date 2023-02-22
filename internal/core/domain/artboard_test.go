package domain

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNotShouldCreateArtboardIfEmptyArtboardID(t *testing.T) {
	artboard, err := NewArtboard(
		"",
		"1",
		"My Artboard",
	)
	assert.EqualError(t, err, "Artboard ID cannot be empty")
	assert.Nil(t, artboard)
}

func TestNotShouldCreateArtboardIfEmptyProjectID(t *testing.T) {
	artboard, err := NewArtboard(
		"1",
		"",
		"My Artboard",
	)
	assert.EqualError(t, err, "Project ID cannot be empty")
	assert.Nil(t, artboard)
}

func TestNotShouldCreateArtboardIfEmptyName(t *testing.T) {
	artboard, err := NewArtboard(
		"1",
		"1",
		"",
	)
	assert.EqualError(t, err, "Name cannot be empty")
	assert.Nil(t, artboard)
}

func TestShouldCreateArtboard(t *testing.T) {
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
		FillColor:   "#FFF",
		StrokeColor: "#000",
		StrokeWidth: 1,
		Opacity:     100,
	}
	layer, err := NewLayer(
		"1",
		"1",
		"My Layer",
		Reactangle,
		size,
		position,
		rotation,
		properties,
	)
	assert.NoError(t, err)
	artboard, err := NewArtboard(
		"1",
		"1",
		"My Artboard",
	)
	artboard.AddLayer(layer)
	assert.NoError(t, err)
	assert.NotNil(t, artboard)
	assert.Equal(t, "My Artboard", artboard.Name)
	assert.Len(t, artboard.Layers, 1)
}

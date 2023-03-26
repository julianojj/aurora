package domain

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReturnErrorIfEmptyArtboardID(t *testing.T) {
	_, err := NewArtboard(
		"",
		"1",
		"My Artboard",
	)
	assert.EqualError(t, err, "Artboard ID cannot be empty")
}

func TestReturnErrorIfEmptyProjectID(t *testing.T) {
	_, err := NewArtboard(
		"1",
		"",
		"My Artboard",
	)
	assert.EqualError(t, err, "Project ID cannot be empty")
}

func TestReturnErrorIfEmptyName(t *testing.T) {
	_, err := NewArtboard(
		"1",
		"1",
		"",
	)
	assert.EqualError(t, err, "Name cannot be empty")
}

func TestCreateArtboard(t *testing.T) {
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
	layer, _ := NewLayer(
		"1",
		"1",
		"My Layer",
		Reactangle,
		size,
		position,
		rotation,
		properties,
	)
	artboard, _ := NewArtboard(
		"1",
		"1",
		"My Artboard",
	)
	artboard.AddLayer(layer)
	assert.NotNil(t, artboard)
	assert.Equal(t, "My Artboard", artboard.Name)
	assert.Len(t, artboard.Layers, 1)
}

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
		&Layer{},
	)
	assert.EqualError(t, err, "Artboard ID cannot be empty")
}

func TestReturnErrorIfEmptyProjectID(t *testing.T) {
	_, err := NewArtboard(
		"1",
		"",
		"My Artboard",
		&Layer{},
	)
	assert.EqualError(t, err, "Project ID cannot be empty")
}

func TestReturnErrorIfEmptyName(t *testing.T) {
	_, err := NewArtboard(
		"1",
		"1",
		"",
		&Layer{},
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
		"layer",
		size,
		position,
		rotation,
		properties,
	)
	artboard, _ := NewArtboard(
		"1",
		"1",
		"My Artboard",
		layer,
	)
	artboard.AddChildren(map[string]any{})
	assert.NotNil(t, artboard)
	assert.Equal(t, "My Artboard", artboard.Name)
	assert.Equal(t, layer, artboard.Layer)
}

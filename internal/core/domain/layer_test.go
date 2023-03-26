package domain

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNotShouldCreateLayerIfEmptyLayerID(t *testing.T) {
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
	_, err := NewLayer(
		"",
		"1",
		"My Layer",
		Reactangle,
		size,
		position,
		rotation,
		properties,
	)
	assert.EqualError(t, err, "Layer ID cannot be empty")
}

func TestNotShouldCreateLayerIfEmptyProjectID(t *testing.T) {
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
	_, err := NewLayer(
		"1",
		"",
		"My Layer",
		Reactangle,
		size,
		position,
		rotation,
		properties,
	)
	assert.EqualError(t, err, "Artboard ID cannot be empty")
}

func TestNotShouldCreateLayerIfEmptyName(t *testing.T) {
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
	_, err := NewLayer(
		"1",
		"1",
		"",
		Reactangle,
		size,
		position,
		rotation,
		properties,
	)
	assert.EqualError(t, err, "Name cannot be empty")
}

func TestNotShouldCreateLayerIfEmptyLayerType(t *testing.T) {
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
	_, err := NewLayer(
		"1",
		"1",
		"My Layer",
		"",
		size,
		position,
		rotation,
		properties,
	)
	assert.EqualError(t, err, "Layer Type cannot be empty")
}

func TestNotShouldCreateLayerIfEmptySize(t *testing.T) {
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
	_, err := NewLayer(
		"1",
		"1",
		"My Layer",
		Reactangle,
		nil,
		position,
		rotation,
		properties,
	)
	assert.EqualError(t, err, "Size cannot be empty")
}

func TestNotShouldCreateLayerIfEmptyPosition(t *testing.T) {
	size := &Size{
		Width:  100,
		Height: 100,
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
	_, err := NewLayer(
		"1",
		"1",
		"My Layer",
		Reactangle,
		size,
		nil,
		rotation,
		properties,
	)
	assert.EqualError(t, err, "Position cannot be empty")
}

func TestNotShouldCreateLayerIfEmptyRotation(t *testing.T) {
	size := &Size{
		Width:  100,
		Height: 100,
	}
	position := &Position{
		X: 0,
		Y: 0,
	}
	properties := &Properties{
		FillColor:   "#FFF",
		StrokeColor: "#000",
		StrokeWidth: 1,
		Opacity:     100,
	}
	_, err := NewLayer(
		"1",
		"1",
		"My Layer",
		Reactangle,
		size,
		position,
		nil,
		properties,
	)
	assert.EqualError(t, err, "Rotation cannot be empty")
}

func TestNotShouldCreateLayerIfEmptyProperties(t *testing.T) {
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

	_, err := NewLayer(
		"1",
		"1",
		"My Layer",
		Reactangle,
		size,
		position,
		rotation,
		nil,
	)
	assert.EqualError(t, err, "Properties cannot be empty")
}

func TestShouldCreateLayer(t *testing.T) {
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
	assert.Equal(t, "My Layer", layer.Name)
	assert.Equal(t, Reactangle, layer.LayerType)
	assert.Equal(t, size, layer.Size)
	assert.Equal(t, position, layer.Position)
	assert.Equal(t, rotation, layer.Rotation)
	assert.Equal(t, properties, layer.Properties)
}

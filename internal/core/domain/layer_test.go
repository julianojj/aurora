package domain

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const LayerName = "My Layer"

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
		Size:        size,
		Position:    position,
		Rotation:    rotation,
		FillColor:   "#FFF",
		StrokeColor: "#000",
		StrokeWidth: 1,
		Opacity:     100,
	}
	_, err := NewLayer(
		"",
		"1",
		"My Layer",
		"layer",
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
		Size: size,
		Position: position,
		Rotation: rotation,
		FillColor:   "#FFF",
		StrokeColor: "#000",
		StrokeWidth: 1,
		Opacity:     100,
	}
	_, err := NewLayer(
		"1",
		"",
		LayerName,
		"layer",
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
		Size: size,
		Position: position,
		Rotation: rotation,
		FillColor:   "#FFF",
		StrokeColor: "#000",
		StrokeWidth: 1,
		Opacity:     100,
	}
	_, err := NewLayer(
		"1",
		"1",
		"",
		"layer",
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
		Size: size,
		Position: position,
		Rotation: rotation,
		FillColor:   "#FFF",
		StrokeColor: "#000",
		StrokeWidth: 1,
		Opacity:     100,
	}
	_, err := NewLayer(
		"1",
		"1",
		LayerName,
		"",
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
		Size: nil,
		Position: position,
		Rotation: rotation,
		FillColor:   "#FFF",
		StrokeColor: "#000",
		StrokeWidth: 1,
		Opacity:     100,
	}
	_, err := NewLayer(
		"1",
		"1",
		LayerName,
		"layer",
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
		Size: size,
		Position: nil,
		Rotation: rotation,
		FillColor:   "#FFF",
		StrokeColor: "#000",
		StrokeWidth: 1,
		Opacity:     100,
	}
	_, err := NewLayer(
		"1",
		"1",
		LayerName,
		"layer",
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
		Size: size,
		Position: position,
		Rotation: nil,
		FillColor:   "#FFF",
		StrokeColor: "#000",
		StrokeWidth: 1,
		Opacity:     100,
	}
	_, err := NewLayer(
		"1",
		"1",
		LayerName,
		"layer",
		properties,
	)
	assert.EqualError(t, err, "Rotation cannot be empty")
}

func TestNotShouldCreateLayerIfEmptyProperties(t *testing.T) {
	_, err := NewLayer(
		"1",
		"1",
		LayerName,
		"layer",
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
		Size: size,
		Position: position,
		Rotation: rotation,
		FillColor:   "#FFF",
		StrokeColor: "#000",
		StrokeWidth: 1,
		Opacity:     100,
	}
	layer, _ := NewLayer(
		"1",
		"1",
		LayerName,
		"layer",
		properties,
	)
	assert.Equal(t, LayerName, layer.Name)
	assert.Equal(t, "layer", layer.LayerType)
	assert.Equal(t, properties, layer.Properties)
	assert.Equal(t, size, layer.Properties.Size)
	assert.Equal(t, position, layer.Properties.Position)
	assert.Equal(t, rotation, layer.Properties.Rotation)
}

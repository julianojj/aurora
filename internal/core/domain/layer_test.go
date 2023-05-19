package domain

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLayer(t *testing.T) {
	LayerName := "My Layer"
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
	t.Run("TestNotShouldCreateLayerIfEmptyLayerID", func(t *testing.T) {
		_, err := NewLayer(
			"",
			"1",
			"My Layer",
			"layer",
			properties,
		)
		assert.EqualError(t, err, "Layer ID cannot be empty")
	})
	t.Run("TestNotShouldCreateLayerIfEmptyProjectID", func(t *testing.T) {
		_, err := NewLayer(
			"1",
			"",
			LayerName,
			"layer",
			properties,
		)
		assert.EqualError(t, err, "Artboard ID cannot be empty")
	})
	t.Run("TestNotShouldCreateLayerIfEmptyName", func(t *testing.T) {
		_, err := NewLayer(
			"1",
			"1",
			"",
			"layer",
			properties,
		)
		assert.EqualError(t, err, "Name cannot be empty")
	})
	t.Run("TestNotShouldCreateLayerIfEmptyLayerType", func(t *testing.T) {
		_, err := NewLayer(
			"1",
			"1",
			LayerName,
			"",
			properties,
		)
		assert.EqualError(t, err, "Layer Type cannot be empty")
	})
	t.Run("TestNotShouldCreateLayerIfEmptySize", func(t *testing.T) {
		properties.Size = nil
		_, err := NewLayer(
			"1",
			"1",
			LayerName,
			"layer",
			properties,
		)
		assert.EqualError(t, err, "Size cannot be empty")
	})
	t.Run("TestNotShouldCreateLayerIfEmptyPosition", func(t *testing.T) {
		properties.Size = size
		properties.Position = nil
		_, err := NewLayer(
			"1",
			"1",
			LayerName,
			"layer",
			properties,
		)
		assert.EqualError(t, err, "Position cannot be empty")
	})
	t.Run("TestNotShouldCreateLayerIfEmptyRotation", func(t *testing.T) {
		properties.Position = position
		properties.Rotation = nil
		_, err := NewLayer(
			"1",
			"1",
			LayerName,
			"layer",
			properties,
		)
		assert.EqualError(t, err, "Rotation cannot be empty")
	})
	t.Run("TestNotShouldCreateLayerIfEmptyProperties", func(t *testing.T) {
		_, err := NewLayer(
			"1",
			"1",
			LayerName,
			"layer",
			nil,
		)
		assert.EqualError(t, err, "Properties cannot be empty")
	})
	t.Run("TestShouldCreateLayer", func(t *testing.T) {
		properties.Rotation = rotation
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
	})
}

package main

import (
	"image/color"

	"github.com/hajimehoshi/ebiten"
)

// Platform - a platform object
type Platform struct {
	image  *ebiten.Image
	ops    *ebiten.DrawImageOptions
	x      int
	y      int
	width  int
	height int
}

func createPlatform(width, height int) (*Platform, error) {
	var w, h int
	if height > 0 {
		h = height
	} else {
		h = 20
	}
	if width > 0 {
		w = width
	} else {
		w = 200
	}
	image, err := ebiten.NewImage(w, h, ebiten.FilterDefault)
	if err != nil {
		return nil, err
	}
	image.Fill(color.RGBA{255, 255, 255, 255})
	ops := &ebiten.DrawImageOptions{}
	return &Platform{image: image, ops: ops, width: width, height: 20}, nil
}

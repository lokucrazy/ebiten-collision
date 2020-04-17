package main

import (
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten"
)

// Player - the player character
type Player struct {
	image  *ebiten.Image
	ops    *ebiten.DrawImageOptions
	x      int
	y      int
	width  int
	height int
}

func createPlayer() (*Player, error) {
	image, err := ebiten.NewImage(10, 10, ebiten.FilterDefault)
	if err != nil {
		return nil, err
	}
	image.Fill(color.RGBA{255, 255, 100, 255})
	ops := &ebiten.DrawImageOptions{}

	return &Player{image: image, ops: ops, x: 0, y: 0, width: 10, height: 10}, nil
}

func (p *Player) handleHorizontal() {
	if ebiten.IsKeyPressed(ebiten.KeyRight) {
		p.MovePlayer(true, 1)
	} else if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		p.MovePlayer(true, -1)
	}
}

func (p *Player) handleVertical() {
	if ebiten.IsKeyPressed(ebiten.KeySpace) {
		p.MovePlayer(false, -5)
	} else {
		p.MovePlayer(false, 1)
	}
}

func (p *Player) isCollidingWith(platform *Platform) (bool, int) {
	if inBetween(p.x, platform.x, platform.x+platform.width) || inBetween(p.x+p.width, platform.x, platform.x+platform.width) {
		py0 := platform.y + platform.height
		if inBetween(p.y, platform.y, platform.y+platform.height) {
			return false, py0 - p.y
		}
		y1 := p.y + p.height
		if inBetween(y1, platform.y, platform.y+platform.height) {
			return false, platform.y - y1
		}
	}
	if inBetween(p.y, platform.y, platform.y+platform.height) || inBetween(p.y+p.height, platform.y, platform.y+platform.height) {
		px0 := platform.x + platform.width
		log.Println(p.x, platform.x, platform.x+platform.width)
		if inBetween(p.x, platform.x, platform.x+platform.width) {
			log.Println(px0 - p.x)
			return true, px0 - p.x
		}
		x1 := p.x + p.width
		if inBetween(x1, platform.x, platform.x+platform.width) {
			log.Println(platform.x - x1)
			return true, platform.x - x1
		}
	}
	return false, 0
}

// MovePlayer - moves the player a ceratin distance in a direction (x=true or y=false)
func (p *Player) MovePlayer(direction bool, distance int) {
	if direction {
		p.ops.GeoM.Translate(float64(distance), 0)
		p.x += distance
	} else {
		p.ops.GeoM.Translate(0, float64(distance))
		p.y += distance
	}
}

package main

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/inpututil"

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
	speed  int
	jump   bool
}

func createPlayer() (*Player, error) {
	image, err := ebiten.NewImage(10, 10, ebiten.FilterDefault)
	if err != nil {
		return nil, err
	}
	image.Fill(color.RGBA{255, 255, 100, 255})
	ops := &ebiten.DrawImageOptions{}

	return &Player{image: image, ops: ops, x: 0, y: 0, width: 10, height: 10, speed: 2, jump: false}, nil
}

func (p *Player) handleHorizontal() {
	if ebiten.IsKeyPressed(ebiten.KeyRight) {
		p.MovePlayer(true, p.speed)
	} else if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		p.MovePlayer(true, -p.speed)
	}
}

func (p *Player) handleVertical() {
	if inpututil.IsKeyJustReleased(ebiten.KeySpace) {
		p.jump = !p.jump
	} else if ebiten.IsKeyPressed(ebiten.KeySpace) && !p.jump {
		p.MovePlayer(false, -20)
		p.jump = !p.jump
	}
	p.MovePlayer(false, 1)
}

func (p *Player) isCollidingWith(platform *Platform) {
	if inBetween(p.x+1, platform.x, (platform.x+platform.width)) || inBetween((p.x+p.width-1), platform.x, (platform.x+platform.width)) {
		if ((p.y + p.height) > platform.y) && (p.y < platform.y) {
			p.MovePlayer(false, platform.y-(p.y+p.height))
		} else if ((platform.y + platform.height) > p.y) && ((p.y + p.height) > (platform.y + platform.height)) {
			p.MovePlayer(false, (platform.y+platform.height)-p.y)
		}
	}
	if inBetween(p.y+1, platform.y, (platform.y+platform.height)) || inBetween((p.y+p.height-1), platform.y, (platform.y+platform.height)) {
		if ((p.x + p.width) > platform.x) && (p.x < platform.x) {
			p.MovePlayer(true, platform.x-(p.x+p.width))
		} else if ((platform.x + platform.width) > p.x) && ((p.x + p.width) > (platform.x + platform.width)) {
			p.MovePlayer(true, (platform.x+platform.width)-p.x)
		}
	}
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

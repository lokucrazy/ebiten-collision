package main

import (
	"log"

	"github.com/hajimehoshi/ebiten"
)

const (
	width  int = 320
	height int = 240
)

// Game - game object
type Game struct {
	player  *Player
	floor   *Platform
	middle  *Platform
	tower   *Platform
	ceiling *Platform
}

// Update - updates the game
func (g *Game) Update(screen *ebiten.Image) error {
	g.player.handleHorizontal()
	g.player.handleVertical()
	g.player.isCollidingWith(g.floor)
	g.player.isCollidingWith(g.ceiling)
	g.player.isCollidingWith(g.middle)
	g.player.isCollidingWith(g.tower)
	return nil
}

// Draw - draws the screen
func (g *Game) Draw(screen *ebiten.Image) {
	screen.DrawImage(g.floor.image, g.floor.ops)
	screen.DrawImage(g.middle.image, g.middle.ops)
	screen.DrawImage(g.tower.image, g.tower.ops)
	screen.DrawImage(g.ceiling.image, g.ceiling.ops)
	screen.DrawImage(g.player.image, g.player.ops)
}

// Layout - layout of screen
func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return width, height
}

func createGame() *Game {
	player, err := createPlayer()
	floor, err := createPlatform(width, 0)
	ceiling, err := createPlatform(width, 0)
	tower, err := createPlatform(10, 100)
	middle, err := createPlatform(40, 0)
	if err != nil {
		log.Fatalf("could not build game due to: %s", err.Error())
	}
	floor.ops.GeoM.Translate(0, float64(height-20))
	floor.x = 0
	floor.y = height - 20
	ceiling.ops.GeoM.Translate(0, -float64(10))
	ceiling.x = 0
	ceiling.y = -10
	middle.ops.GeoM.Translate(100, 150)
	middle.x = 100
	middle.y = 150
	tower.ops.GeoM.Translate(200, 100)
	tower.x = 200
	tower.y = 100
	player.x = 20
	player.y = 20
	player.ops.GeoM.Translate(float64(player.x), float64(player.y))
	return &Game{player: player, floor: floor, middle: middle, ceiling: ceiling, tower: tower}
}

func main() {
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Collision")
	game := createGame()
	if err := ebiten.RunGame(game); err != nil {
		panic(err)
	}
}

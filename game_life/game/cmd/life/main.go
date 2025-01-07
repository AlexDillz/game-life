package main

import (
	"game/pkg/life"
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

const (
	screenWidth  = 800
	screenHeight = 800
	cellSize     = 10
	gridWidth    = screenWidth / cellSize
	gridHeight   = screenHeight / cellSize
)

type Game struct {
	currentWorld *life.World
	nextWorld    *life.World
}

func NewGame() *Game {
	rand.Seed(time.Now().UnixNano())
	currentWorld, _ := life.NewWorld(gridHeight, gridWidth)
	currentWorld.RandInit(40)

	nextWorld, _ := life.NewWorld(gridHeight, gridWidth)

	return &Game{
		currentWorld: currentWorld,
		nextWorld:    nextWorld,
	}
}

func (g *Game) Update() error {
	// Обновляем состояние игры
	life.NextState(g.currentWorld, g.nextWorld)
	g.currentWorld, g.nextWorld = g.nextWorld, g.currentWorld
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	// Рисуем текущий мир
	for y := 0; y < gridHeight; y++ {
		for x := 0; x < gridWidth; x++ {
			if g.currentWorld.Cells[y][x] {
				// Рисуем живую клетку
				ebitenutil.DrawRect(screen, float64(x*cellSize), float64(y*cellSize), cellSize, cellSize, ebiten.ColorWhite)
			}
		}
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func main() {
	game := NewGame()

	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Game of Life")

	// Запускаем игру
	if err := ebiten.RunGame(game); err != nil {
		panic(err)
	}
}

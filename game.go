package main

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	current_scene	Scene
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.current_scene.Draw(screen)
}

func (g *Game) Update() error {
	return g.current_scene.Update(g)	
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return g.current_scene.Layout(1280, 720)
}

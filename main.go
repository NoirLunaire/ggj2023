package main

import (
	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	ebiten.SetWindowSize(1280, 720) 
	applySettings(nil)

	ImguiStyle()
	g := &Game{
		current_scene: &Menu{},
	}

	if err := ebiten.RunGame(g); err != nil {
		if err != quit_game {
			panic(err)
		}
	}
}


package main

import (
	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	ebiten.SetWindowSize(1280, 720) 
	//ebiten.SetFullscreen(true);

	g := &Game{
		current_scene: &Menu{},
	}

	if err := ebiten.RunGame(g); err != nil {
		if err != quit_game {
			panic(err)
		}
	}
}


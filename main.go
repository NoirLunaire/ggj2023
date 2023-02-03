package main

import (
	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	ebiten.SetFullscreen(true);

	g := &Game{
		current_scene: NewMenu(),
	}

	if err := ebiten.RunGame(g); err != nil {
		if err != quit_game {
			panic(err)
		}
	}
}


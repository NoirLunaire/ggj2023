package main

import (
	"fmt"
	"image/color"
	
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type GameScene struct {
}

func (m *GameScene) Draw (screen *ebiten.Image) {
	screen.Fill(color.RGBA{ 0, 0, 0, 0xff })
	ebitenutil.DebugPrint(screen, fmt.Sprintf("Game Scene"))
	mgr.Draw(screen)
}

func (m *GameScene) Update(g *Game) error {
	mgr.Update(1.0/60.0)
	mgr.BeginFrame()
	mgr.EndFrame()
	return nil
}

func (m *GameScene) Layout (outsideWidth, outsideHeight int) (int, int) {	
	mgr.SetDisplaySize(float32(1280), float32(720))
	return 1280, 720
}

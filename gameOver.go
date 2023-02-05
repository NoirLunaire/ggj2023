package main

import (
	. "ggj2023/game"
	"image/color"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/inkyblackness/imgui-go/v4"
)

type GameOver struct {
	game_state	*State
}

func (m *GameOver) Draw (screen *ebiten.Image) {
	screen.Fill(color.RGBA{ 0, 0, 0, 0xff })
	DrawTPS(screen)
	mgr.Draw(screen)
}

func (m *GameOver) Update(g *Game) error {
	mgr.Update(1.0/60.0)
	bole := true
	mgr.BeginFrame()
	{
		imgui.SetNextWindowPos(imgui.Vec2{ 1280 / 2 - 100, 720 / 2 })
		imgui.BeginV("Game over", &bole, gui_flags)
		imgui.Text("Vous avez perdu...")
		if imgui.ButtonV("Quitter", imgui.Vec2{ 200, 200 }) {
			g.current_scene = &Menu{}
		}
		imgui.End()
	}
	mgr.EndFrame()
	return nil
}

func (m *GameOver) Layout (outsideWidth, outsideHeight int) (int, int) {	
	mgr.SetDisplaySize(float32(1280), float32(720))
	return 1280, 720
}

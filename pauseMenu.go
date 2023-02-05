package main

import (
	. "ggj2023/game"
	"image/color"
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/inkyblackness/imgui-go/v4"
)

type PauseMenu struct {
	gameData	*GameScene
	saved		bool
}

func (m *PauseMenu) Draw (screen *ebiten.Image) {
	screen.Fill(color.RGBA{ 0, 0, 0, 0xff })
	DrawTPS(screen)
	mgr.Draw(screen)
}

func (m *PauseMenu) Update (g *Game) error {
	mgr.Update(1.0/60.0)
	bole := true
	mgr.BeginFrame()
	{
		imgui.SetNextWindowPos(imgui.Vec2{ 1280 / 2 - 300, 720 / 2 - 150})
		imgui.BeginV("Menu", &bole, gui_flags)
		if m.saved {
			imgui.Text("saved")
		}

		setColor(0);
		if imgui.ButtonV("Reprendre", imgui.Vec2{ 200, 300 }) {
			g.current_scene = m.gameData
		}
		imgui.PopStyleColor()
		imgui.PopStyleColor()
		setColor(1)
		imgui.SameLine();
		if imgui.ButtonV("Sauvegarder", imgui.Vec2{ 200, 300 }) {
			fmt.Println("saving...")
			SaveGame("testing", m.gameData.game_state)
		}
		imgui.PopStyleColor()
		imgui.PopStyleColor()
		setColor(2)
		imgui.SameLine();
		if imgui.ButtonV("Quitter", imgui.Vec2{ 200, 300 }) {
			return quit_game
		}
		imgui.PopStyleColor()
		imgui.PopStyleColor()
		imgui.End()
	}
	mgr.EndFrame()
	return nil

}

func (m *PauseMenu) Layout (outsideWidth, outsideHeight int) (int, int) {	
	mgr.SetDisplaySize(float32(1280), float32(720))
	return 1280, 720
}


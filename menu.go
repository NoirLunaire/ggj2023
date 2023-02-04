package main

import (
	"fmt"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/inkyblackness/imgui-go/v4"
)

type Menu struct {}

func (m *Menu) Draw (screen *ebiten.Image) {
	screen.Fill(color.RGBA{ 0, 0, 0, 0xff })
	ebitenutil.DebugPrint(screen, fmt.Sprintf("TPS: %.2f", ebiten.CurrentTPS()))
	mgr.Draw(screen)
}

func (m *Menu) Update(g *Game) error {
	mgr.Update(1.0/60.0)
	bole := true
	mgr.BeginFrame()
	{
		imgui.SetNextWindowPos(imgui.Vec2{ 1280 / 2 - 300, 720 / 2 - 150})
		imgui.BeginV("Menu", &bole, gui_flags)
		setColor(0);
		if imgui.ButtonV("Nouvelle partie", imgui.Vec2{ 200, 300 }) {
			fmt.Println("nouvelle partie :)")
			g.current_scene = NewGame()
		}
		imgui.PopStyleColor()
		imgui.PopStyleColor()
		setColor(1)
		imgui.SameLine();
		if imgui.ButtonV("Paramètre", imgui.Vec2{ 200, 300 }) {
			fmt.Println("paramètre scene")
			g.current_scene = NewSettings()
		}
		imgui.PopStyleColor()
		imgui.PopStyleColor()
		setColor(2)
		imgui.SameLine();
		if imgui.ButtonV("Quit Game", imgui.Vec2{ 200, 300 }) {
			return quit_game
		}
		imgui.PopStyleColor()
		imgui.PopStyleColor()
		imgui.End()
	}
	mgr.EndFrame()
	return nil
}

func (m *Menu) Layout (outsideWidth, outsideHeight int) (int, int) {	
	mgr.SetDisplaySize(float32(1280), float32(720))
	return 1280, 720
}


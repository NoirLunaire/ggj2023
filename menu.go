package main

import (
	"fmt"
	"image/color"

	"github.com/gabstv/ebiten-imgui/renderer"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/inkyblackness/imgui-go/v4"
)

type Menu struct {
	mgr *renderer.Manager
}

func NewMenu () *Menu {
	return &Menu {
		mgr: renderer.New(nil),
	}
}

func (m *Menu) Draw (screen *ebiten.Image) {
	screen.Fill(color.RGBA{ 0, 0, 0, 0xff })
	ebitenutil.DebugPrint(screen, fmt.Sprintf("TPS: %.2f", ebiten.CurrentTPS()))
	m.mgr.Draw(screen)
}

func (m *Menu) Update() error {
	m.mgr.Update(1.0/60.0)
	bole := true
	flags := imgui.WindowFlagsNoTitleBar + imgui.WindowFlagsNoResize + imgui.WindowFlagsNoMove + imgui.WindowFlagsNoCollapse + imgui.WindowFlagsAlwaysAutoResize
	m.mgr.BeginFrame()
	{
		imgui.SetNextWindowPos(imgui.Vec2{ 1280 / 2 - 100, 720 / 2 - 200})
		imgui.BeginV("Menu", &bole, flags)

		if imgui.ButtonV("New Game", imgui.Vec2{ 200, 200 }) {
		}
		if imgui.ButtonV("Quit Game", imgui.Vec2{ 200, 200 }) {
			return quit_game
		}
		imgui.End()
	}
	m.mgr.EndFrame()
	return nil
}

func (m *Menu) Layout (outsideWidth, outsideHeight int) (int, int) {	
	m.mgr.SetDisplaySize(float32(1280), float32(720))
	return 1280, 720
}

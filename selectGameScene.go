package main

import (
	"fmt"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/inkyblackness/imgui-go/v4"
)

type SelectGame struct {}

func (m *SelectGame) Draw (screen *ebiten.Image) {
	screen.Fill(color.RGBA{ 0, 0, 0, 0xff })
	ebitenutil.DebugPrint(screen, fmt.Sprintf("TPS: %.2f", ebiten.CurrentTPS()))
	mgr.Draw(screen)
}

func (m *SelectGame) Update(g *Game) error {
	mgr.Update(1.0/60.0)
	bole := true
	var items = []string{"Item 1", "Item 2", "Item 3", "Item 2", "Item 3", "Item 2", "Item 3", "Item 2", "Item 3", "Item 2", "Item 3"}
	choice := int32(0)
	mgr.BeginFrame()
	{
		imgui.SetNextWindowPos(imgui.Vec2{ 1280 / 2 - 300, 720 / 2 - 150})
		imgui.BeginV("SelectGame", &bole, gui_flags)

		imgui.ListBoxV("Sauvegardes", &choice, items, 5)
		if imgui.ButtonV("Retour", imgui.Vec2{ 100, 50 }) {
			fmt.Println("Retour menu")
			g.current_scene = &Menu{}
		}
		imgui.SameLine()
		if imgui.ButtonV("Nouvelle Partie", imgui.Vec2{ 150, 50 }) {
			fmt.Println("Création d'une nouvelle partie")
			g.current_scene = NewGame()
		}

		imgui.End()
	}
	mgr.EndFrame()
	return nil
}

func (m *SelectGame) Layout (outsideWidth, outsideHeight int) (int, int) {	
	mgr.SetDisplaySize(float32(1280), float32(720))
	return 1280, 720
}


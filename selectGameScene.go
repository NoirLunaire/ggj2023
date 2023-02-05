package main

import (
	"fmt"
	"image/color"

	. "ggj2023/game"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/inkyblackness/imgui-go/v4"
)

type SelectGame struct {
	choice	int32
	loaded	error
}

func (m *SelectGame) Draw (screen *ebiten.Image) {
	screen.Fill(color.RGBA{ 0, 0, 0, 0xff })
	DrawTPS(screen)
	mgr.Draw(screen)
}

func (m *SelectGame) Update(g *Game) error {
	mgr.Update(1.0/60.0)
	bole := true
	items := GetSaves()
	mgr.BeginFrame()
	{
		if m.loaded != nil {
			imgui.SetNextWindowPos(imgui.Vec2{ 1280 / 2 - 100, 100 })
			imgui.BeginV("Error", &bole, gui_flags)
			imgui.Text("Erreur lors du chargement de la sauvegarde")
			imgui.End()
		}

		imgui.SetNextWindowPos(imgui.Vec2{ 1280 / 2 - 300, 720 / 2 - 150})
		imgui.BeginV("SelectGame", &bole, gui_flags)

		if len(items) > 0 {
			imgui.ListBoxV("Sauvegardes", &m.choice, items, 5)
		}

		if imgui.ButtonV("Retour", imgui.Vec2{ 100, 50 }) {
			fmt.Println("Retour menu")
			g.current_scene = &Menu{}
		}
		imgui.SameLine()
		if imgui.ButtonV("Nouvelle Partie", imgui.Vec2{ 150, 50 }) {
			fmt.Println("Création d'une nouvelle partie")
			g.current_scene = NewGame()
		}

		if len(items) > 0 {
			if imgui.ButtonV("Charger", imgui.Vec2{ 150, 50 }) {
				fmt.Println("Chargement d'une partie")
				s, err := LoadSave(items[m.choice])
				if err == nil {
					g.current_scene = LoadGame(s)
				} else { m.loaded = err }
			}
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


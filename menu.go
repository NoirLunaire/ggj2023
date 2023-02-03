package main

import (
	"fmt"
	"image/color"

	. "ggj2023/game"	
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
		imgui.SetNextWindowPos(imgui.Vec2{ 1280 / 2 - 100, 720 / 2 - 200})
		imgui.BeginV("Menu", &bole, gui_flags)

		if imgui.ButtonV("New Game", imgui.Vec2{ 200, 200 }) {
			fmt.Println("nouvelle partie :)")
			g.current_scene = &GameScene{
				nil,
				true,
				&Event{
					0,
					"Test",
					"Ceci est un test messir ! Vous avez les choix suivants : ",
					[]int{ 0, 1 },
				},
			}
		}
		if imgui.ButtonV("Quit Game", imgui.Vec2{ 200, 200 }) {
			return quit_game
		}
		imgui.End()
	}
	mgr.EndFrame()
	return nil
}

/*
	id		int
	title		string
	description	string
	choices		[]int

	game_state	*State
	has_event	bool
	current_event	*Event

*/

func (m *Menu) Layout (outsideWidth, outsideHeight int) (int, int) {	
	mgr.SetDisplaySize(float32(1280), float32(720))
	return 1280, 720
}

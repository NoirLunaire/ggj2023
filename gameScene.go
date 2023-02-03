package main

import (
	"fmt"
	"image/color"
	"strconv"

	. "ggj2023/game"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/inkyblackness/imgui-go/v4"
)

type GameScene struct {
	game_state	*State
	has_event	bool
	current_event	*Event
}

func (m *GameScene) Draw (screen *ebiten.Image) {
	screen.Fill(color.RGBA{ 0, 0, 0, 0xff })
	ebitenutil.DebugPrint(screen, fmt.Sprintf("TPS: %.2f", ebiten.CurrentTPS()))
	mgr.Draw(screen)
}

func (m *GameScene) Update(g *Game) error {
	mgr.Update(1.0/60.0)
	bole := true
	mgr.BeginFrame()
	{
		if m.has_event {
			imgui.SetNextWindowPos(imgui.Vec2{ 10, 10 })
			imgui.BeginV(m.current_event.Title, &bole, imgui.WindowFlagsNoResize + imgui.WindowFlagsNoMove + imgui.WindowFlagsNoCollapse + imgui.WindowFlagsAlwaysAutoResize)
			imgui.Text(m.current_event.Description)
			for i := 0; i < len(m.current_event.Choices); i++ {
				if imgui.Button( strconv.Itoa(m.current_event.Choices[i]) ) {
					fmt.Println(" :) tu as appuyé sur le choix ", m.current_event.Choices[i])
					m.has_event = false
				}
			}
			imgui.End()
		}
	}
	mgr.EndFrame()

	/*
	mgr.BeginFrame()
	{
		imgui.SetNextWindowPos(imgui.Vec2{ 1200, 30 })
		imgui.BeginV("Map", &bole, gui_flags)
	}
	mgr.EndFrame()
	*/
	return nil
}

func (m *GameScene) Layout (outsideWidth, outsideHeight int) (int, int) {	
	mgr.SetDisplaySize(float32(1280), float32(720))
	return 1280, 720
}

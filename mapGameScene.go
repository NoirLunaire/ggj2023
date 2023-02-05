package main

import (
	//"fmt"
	"log"
	"image/color"
	"strconv"
	_ "image/png"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"github.com/inkyblackness/imgui-go/v4"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type MapGameScene struct {
	gameData	*GameScene
	imgIconPopulation	*ebiten.Image
	imgIconGold	*ebiten.Image
	imgIconHapiness	*ebiten.Image	
	imgBackground	*ebiten.Image
}
	
	

func NewGameMap (m *GameScene) *MapGameScene {

	imgIconPopulation, _, err := ebitenutil.NewImageFromFile("data/image/populations.png")
	imgIconGold, _, erre := ebitenutil.NewImageFromFile("data/image/coin.png")
	imgIconHapiness, _, errer := ebitenutil.NewImageFromFile("data/image/prestige.png")
	imgBackground, _, errers := ebitenutil.NewImageFromFile("data/image/menuCard0.png")
	if err != nil  || erre != nil || errer != nil || errers != nil{
		log.Fatalf("Failed to load image: %v", err)
	}

	return &MapGameScene{
		m,
		imgIconPopulation,
		imgIconGold,
		imgIconHapiness,
		imgBackground,
	}
}

func (m *MapGameScene) Draw (screen *ebiten.Image) {

	ot := &ebiten.DrawImageOptions{}
	ot.GeoM.Translate(0, 0)
	ot.GeoM.Scale(0.665, 0.665)
	screen.DrawImage(m.imgBackground, ot)
	// draw resources
	text.Draw(screen, "Prestige : " + strconv.Itoa(m.gameData.game_state.Happiness), m.gameData.font, 100, 535, color.White)
	text.Draw(screen, "Trésorie : " + strconv.Itoa(m.gameData.game_state.Money), m.gameData.font, 100, 610, color.White)
	text.Draw(screen, "Populations : " + strconv.Itoa(m.gameData.game_state.Population), m.gameData.font, 100, 685, color.White)
	
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(200, 2500)
	op.GeoM.Scale(0.2, 0.2)
	screen.DrawImage(m.imgIconHapiness, op)
	op.GeoM.Translate(0, 75)
	screen.DrawImage(m.imgIconGold, op)
	op.GeoM.Translate(0, 75)
	screen.DrawImage(m.imgIconPopulation, op)

	DrawDate(screen,m.gameData)
	DrawTPS(screen)
	mgr.Draw(screen)
}

func (m *MapGameScene) Update(g *Game) error {
	mgr.Update(1.0/60.0)
	bole := true
	// draw event window
	mgr.BeginFrame()
	{
		imgui.SetNextWindowPos(imgui.Vec2{ 25, 24 })
		imgui.BeginV("next", &bole, gui_flags)
		if imgui.ButtonV("Retour", imgui.Vec2{50, 50}) {
			g.current_scene = m.gameData
		}
		imgui.End()
	}
	mgr.EndFrame()

	return nil
}

func (m *MapGameScene) Layout (outsideWidth, outsideHeight int) (int, int) {	
	mgr.SetDisplaySize(float32(1280), float32(720))
	return 1280, 720
}

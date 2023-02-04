package main

import (
	"fmt"
	"log"
	"bytes"
	"image/color"
	"math/rand"
	"strconv"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
	"github.com/hajimehoshi/ebiten/v2/examples/resources/fonts"
	"github.com/hajimehoshi/ebiten/v2/audio/mp3"
	music "ggj2023/data/music"

	. "ggj2023/game"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/inkyblackness/imgui-go/v4"
	"github.com/hajimehoshi/ebiten/v2/text"
	"github.com/hajimehoshi/ebiten/v2/audio"
)

type GameScene struct {
	game_state	*State
	font		font.Face
	audioContext	*audio.Context
	audioPlayer	*audio.Player
	has_event	bool
	is_pause	bool
	timer		float64
	current_event	*Event
}

func NewGame () *GameScene {
	s, err := mp3.DecodeWithoutResampling(bytes.NewReader(music.Ost_mp3))
	settings := NewSettings()

	if err != nil {
		log.Fatal(err)
	}
	context := audio.NewContext(44100)
	loop := audio.NewInfiniteLoop(s, s.Length() - 1)
	player, err := context.NewPlayer(loop)
	if err != nil {
		log.Fatal(err)
	}
	player.SetVolume(float64(settings.musicVolume))

	tt, err := opentype.Parse(fonts.MPlus1pRegular_ttf)
	if err != nil {
		log.Fatal(err)
	}

	const dpi = 72
	ourFont, err := opentype.NewFace(tt, &opentype.FaceOptions{
		Size:    36,
		DPI:     dpi,
		Hinting: font.HintingFull,
	})
	
	player.Play()
	return &GameScene{
		NewState(),
		ourFont,
		context,
		player,
		false,
		false,
		5 * ebiten.ActualTPS(),
		nil,
	}
}

func (m *GameScene) Draw (screen *ebiten.Image) {
	if !m.is_pause {
		if int(m.timer) % int(ebiten.CurrentTPS()) == 0 {
			m.game_state.Date = m.game_state.Date.AddDate(0, 0, 1)
		}
		m.timer -= 1
	}
	screen.Fill(color.RGBA{ 0, 0, 0, 0xff })
	mgr.Update(1.0/60.0)
	bole := true
	// draw event window
	mgr.BeginFrame()
	{
		if m.has_event {
			imgui.SetNextWindowPos(imgui.Vec2{ 640, 360 })
			imgui.BeginV(m.current_event.Title, &bole, imgui.WindowFlagsNoResize + imgui.WindowFlagsNoMove + imgui.WindowFlagsNoCollapse + imgui.WindowFlagsAlwaysAutoResize)
			imgui.Text(m.current_event.Description)
			for i := 0; i < len(m.current_event.Choices); i++ {
				if imgui.Button( m.game_state.ChoiceList[m.current_event.Choices[i]].Title ) {
					m.game_state.Effects[m.current_event.Choices[i]](m.game_state)
					m.has_event = false
					m.timer = ebiten.ActualTPS() * 6
					m.current_event = nil
					m.is_pause = false
					break
				}
				if imgui.IsItemHovered() {
					imgui.SetTooltip(m.game_state.ChoiceList[m.current_event.Choices[i]].Description)
				}
			}
			imgui.End()
		}

		if m.timer <= 0 {
			if len(m.game_state.EventPool) > 0 {
				r := rand.Intn(len(m.game_state.EventPool))
				m.current_event = m.game_state.EventList[m.game_state.EventPool[r]]
				m.has_event = true
				m.is_pause = true
			}
		}

		imgui.SetNextWindowPos(imgui.Vec2{ 1000, 650 })
		imgui.BeginV("next", &bole, gui_flags)
		if imgui.ButtonV("Pause", imgui.Vec2{50, 50}) {
			m.is_pause = !m.is_pause
		}
		imgui.End()
	}
	mgr.EndFrame()

	// draw resources
	text.Draw(screen, "Happiness : " + strconv.Itoa(m.game_state.Happiness), m.font, 30, 100, color.White)
	text.Draw(screen, "Money : " + strconv.Itoa(m.game_state.Money), m.font, 30, 150, color.White)
	text.Draw(screen, "Population : " + strconv.Itoa(m.game_state.Population), m.font, 30, 200, color.White)
	if m.is_pause {
		text.Draw(screen, m.game_state.Date.Format("02-01-2006"), m.font, 30, 250, color.White)
	} else {
		text.Draw(screen, m.game_state.Date.Format("02-01-2006"), m.font, 30, 250, color.White)
	}
	// draw tps
	ebitenutil.DebugPrint(screen, fmt.Sprintf("TPS: %.2f", ebiten.CurrentTPS()))
	mgr.Draw(screen)
}

func (m *GameScene) Update(g *Game) error {
	if !m.audioPlayer.IsPlaying() {
		m.audioPlayer.Play()
	}
	return nil
}

func (m *GameScene) Layout (outsideWidth, outsideHeight int) (int, int) {	
	mgr.SetDisplaySize(float32(1280), float32(720))
	return 1280, 720
}

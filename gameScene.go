package main

import (
	"log"
	"bytes"
	"image/color"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
	"github.com/hajimehoshi/ebiten/v2/examples/resources/fonts"
	"github.com/hajimehoshi/ebiten/v2/audio/mp3"
	music "ggj2023/data/music"
	_ "image/png"

	. "ggj2023/game"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/inkyblackness/imgui-go/v4"
	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	
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
	name	string
	imgHall	*ebiten.Image

	imgBgLeft	*ebiten.Image
	imgBgRight	*ebiten.Image
	imgBorderDate	*ebiten.Image
	
	imgTower	*ebiten.Image
}

func LoadGame (name string ,s *State) *GameScene {
	scene := NewGame()
	scene.name = name
	if (s != nil){
		scene.game_state = s
	}
	return scene
}

func NewGame () *GameScene {
	s, err := mp3.DecodeWithoutResampling(bytes.NewReader(music.Ost_mp3))
	settings := NewSettings(nil)

	if err != nil {
		log.Fatal(err)
	}
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

	imgHall, _, err := ebitenutil.NewImageFromFile("data/image/hall.png")
	imgBgLeft, _, erre := ebitenutil.NewImageFromFile("data/image/backgroundLeft.png")
	imgBgRight, _, errer := ebitenutil.NewImageFromFile("data/image/backgroundRight.png")
	imgBorderDate, _, errero := ebitenutil.NewImageFromFile("data/image/dateBorder.png")
	imgTower, _, erreror := ebitenutil.NewImageFromFile("data/image/tower1.png")
	if err != nil  || erre != nil || errer != nil || errero != nil || erreror != nil{
		log.Fatalf("Failed to load image: %v", err)
	}
	

	return &GameScene{
		NewState(),
		ourFont,
		context,
		player,
		false,
		false,
		5 * ebiten.ActualTPS(),
		nil,
		"default",
		imgHall,
		imgBgLeft,
		imgBgRight,
		imgBorderDate,
		imgTower,
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
	
	x, _ := ebiten.CursorPosition()
	diff := (x - 1280/2)/16


	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(-100 - float64(diff), 0)
	op.GeoM.Scale(0.80, 0.562)
	screen.DrawImage(m.imgBgLeft, op)
	op.GeoM.Translate(740 - float64(diff), 0)
	screen.DrawImage(m.imgBgRight, op)

	
	
	
	for i := 0 ; i < len( m.game_state.Village.TabBuild);i++ {
		op = &ebiten.DrawImageOptions{}
		op.GeoM.Scale(m.game_state.Village.TabPositionBuild[i][2], m.game_state.Village.TabPositionBuild[i][3])
		op.GeoM.Translate(m.game_state.Village.TabPositionBuild[i][0] - float64(diff), m.game_state.Village.TabPositionBuild[i][1])
		screen.DrawImage(m.imgTower,op)
	}

	op = &ebiten.DrawImageOptions{}
	op.GeoM.Translate(0, 0)
	op.GeoM.Scale(0.67, 0.562)
	screen.DrawImage(m.imgHall, op)

	DrawDate(screen,m)
	DrawTPS(screen)
	mgr.Draw(screen)
}

func (m *GameScene) Update(g *Game) error {
	if EndGame(m.game_state) {
		m.audioPlayer.Close()
		g.current_scene = &GameOver{ m.game_state }
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyEscape) {
		g.current_scene = &PauseMenu{
			m,
			false,
		}
	}

	mgr.Update(1.0/60.0)
	bole := true
	// draw event window
	mgr.BeginFrame()
	{
		
		if m.has_event {
			imgui.SetNextWindowPos(imgui.Vec2{ 1280 / 2 - 300, 720 / 2 })
			imgui.BeginV(m.current_event.Title, &bole, imgui.WindowFlagsNoResize + imgui.WindowFlagsNoMove + imgui.WindowFlagsNoCollapse + imgui.WindowFlagsAlwaysAutoResize)
			imgui.Text(m.current_event.Description)
			for i := 0; i < len(m.current_event.Choices); i++ {
				if imgui.Button( m.game_state.ChoiceList[m.current_event.Choices[i]].Title ) {
					m.game_state.Effects[m.current_event.Choices[i]](m.game_state)
					if m.current_event.Choices[i] == 10 {
						m.current_event = m.game_state.EventList[5]
					} else {
						m.has_event = false
						m.timer = ebiten.ActualTPS() * (random.Intn(5) + 5)
						m.current_event = nil
						m.is_pause = false
					}
					break
				}
				if imgui.IsItemHovered() {
					imgui.SetTooltip(m.game_state.ChoiceList[m.current_event.Choices[i]].Description)
				}
			}
			imgui.End()
		}

		if m.timer <= 0 {
			if len(m.game_state.EventPool) > 0 && m.current_event == nil {
				r := random.Intn(len(m.game_state.EventPool))
				m.current_event = m.game_state.EventList[m.game_state.EventPool[r]]
				m.has_event = true
				m.is_pause = true
			}
		}

		imgui.SetNextWindowPos(imgui.Vec2{ 1000, 650 })
		imgui.BeginV("next", &bole, gui_flags)
		if imgui.ButtonV("Pause", imgui.Vec2{50, 50}) {
			if !m.has_event {
				m.is_pause = !m.is_pause
			}
		}
		imgui.SameLine();
		if imgui.ButtonV("MAP", imgui.Vec2{50, 50}) {
			m.is_pause = true
			g.current_scene = NewGameMap(m)
		}
		imgui.End()
	}
	mgr.EndFrame()


	return nil
}

func (m *GameScene) Layout (outsideWidth, outsideHeight int) (int, int) {	
	mgr.SetDisplaySize(float32(1280), float32(720))
	return 1280, 720
}

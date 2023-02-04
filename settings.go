package main

import (
	"fmt"
	"image/color"
	"bufio"
	"os"
	"strconv"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/inkyblackness/imgui-go/v4"
)

type SettingsScene struct {
	fullscreen		bool
	musicVolume		float32
	effectsVolume	float32
}

func NewSettings () *SettingsScene {
	file, err := os.Open("./config/settings.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return nil
	}
	defer file.Close()

	// Create a new scanner to read the file line by line
	scanner := bufio.NewScanner(file)

	// Create a slice to store the contents of the file
	lines := make([]string, 0)

	// Read each line of the file
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}

	// Check for any errors that may have occurred during scanning
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return nil
	}

	fullscreenn, errer := strconv.ParseBool(lines[0])
	musicVolume, erre := strconv.ParseFloat(lines[1], 32)
	musicEffects, err := strconv.ParseFloat(lines[2], 32)
	if erre != nil || err != nil || errer != nil{
		fmt.Println("Error while trying to parse settings.Txt file")
		return nil
	}
	return &SettingsScene{
		fullscreenn,
		float32(musicVolume),
		float32(musicEffects),
	}
}

func (m *SettingsScene) Draw (screen *ebiten.Image) {
	screen.Fill(color.RGBA{ 0, 0, 0, 0xff })
	ebitenutil.DebugPrint(screen, fmt.Sprintf("TPS: %.2f", ebiten.CurrentTPS()))
	mgr.Draw(screen)
}

func (m *SettingsScene) Update(g *Game) error {
	mgr.Update(1.0/60.0)
	bole := true
	mgr.BeginFrame()
	{
		imgui.SetNextWindowPos(imgui.Vec2{ 1280 / 2 - 150, 720 / 2 - 100})
		imgui.BeginV("Menu", &bole, gui_flags)
		
		
		imgui.Checkbox("Fullscreen", &m.fullscreen)

		imgui.SliderFloat("Music Volume", &m.musicVolume, 0, 1)
		imgui.SliderFloat("Effects Volume", &m.effectsVolume, 0, 1)
		if imgui.ButtonV("Retour", imgui.Vec2{ 100, 50 }) {
			fmt.Println("Retour menu")
			g.current_scene = &Menu{}
		}
		imgui.SameLine();
		if imgui.ButtonV("Appliquer changements", imgui.Vec2{ 200, 50 }) {
			applySettings()

			fmt.Println("Appliquer changements")
			
			file, err := os.OpenFile("./config/settings.txt", os.O_CREATE|os.O_WRONLY, 0644)
			if err != nil {
				fmt.Println("Error opening config file settings.txt:", err)
				return nil
			}
			defer file.Close()

			boolstr := "0\n"
			if (m.fullscreen){
				boolstr = "1\n"
			}
			_, err = file.WriteString(boolstr)
			_, erre := file.WriteString((strconv.FormatFloat(float64(m.musicVolume), 'f', 2, 32)+"\n"))
			_, errer := file.WriteString(strconv.FormatFloat(float64(m.effectsVolume), 'f', 2, 32))
			if err != nil || erre != nil || errer != nil{
				fmt.Println("Error writing to file settings.txt")
				return nil
			}

			fmt.Println("Write successful")
			
		}
		//fmt.Println("Effects Volume :",&m.effectsVolume)
		imgui.End()

}
	mgr.EndFrame()
	return nil
}

func (m *SettingsScene) Layout (outsideWidth, outsideHeight int) (int, int) {	
	mgr.SetDisplaySize(float32(1280), float32(720))
	return 1280, 720
}


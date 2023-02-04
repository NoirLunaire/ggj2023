package main

import (
	"github.com/inkyblackness/imgui-go/v4"
	"github.com/hajimehoshi/ebiten/v2"
)

type Scene interface {
	Draw(screen *ebiten.Image)
	Layout(outsideWidth, outsideHeight int) (int, int)	
	Update(g *Game) error
}

func ImguiStyle () {
	imgui.PushStyleColor(21, imgui.Vec4{ 0.7, 0, 0, 1.0 })
	imgui.PushStyleColor(22, imgui.Vec4{ 1.0, 0, 0, 1.0 })
	imgui.PushStyleColor(23, imgui.Vec4{ 1.0, 0, 0, 1.0 })
	imgui.PushStyleColor(11, imgui.Vec4{ 0.7, 0, 0, 1.0 })
}



func setColor(a int) {
	r := float32(0.0)
	g := float32(0.0)
	b := float32(0.0)
	switch a {
		case 0:
			r = 0.0
			g = 0.33
			b = 0.64
		case 1:
			r = 1.0
			g = 1.0
			b = 1.0
		case 2:
			r = 0.93
			g = 0.25
			b = 0.20
		case 3:
	}
	imgui.PushStyleColor(21, imgui.Vec4{ r,g,b, 1.0 })
	imgui.PushStyleColor(0, imgui.Vec4{ 0, 0, 0, 1.0 })
	
}
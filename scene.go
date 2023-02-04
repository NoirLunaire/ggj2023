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

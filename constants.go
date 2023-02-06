package main

import (
	"math/rand"
	"time"
	"errors"
	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/gabstv/ebiten-imgui/renderer"
	"github.com/inkyblackness/imgui-go/v4"
)

var (
	random = rand.New(rand.NewSource(time.Now().UnixNano()))
	context = audio.NewContext(44100)
	quit_game = errors.New("regular termination")
	mgr = renderer.New(nil)
	gui_flags = imgui.WindowFlagsNoTitleBar + imgui.WindowFlagsNoResize + imgui.WindowFlagsNoMove + imgui.WindowFlagsNoCollapse + imgui.WindowFlagsAlwaysAutoResize
)

const (
	Blue int = iota
	White 
	Red
)

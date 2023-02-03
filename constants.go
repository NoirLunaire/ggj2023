package main

import (
	"errors"
	"github.com/gabstv/ebiten-imgui/renderer"
)

var (
	quit_game = errors.New("regular termination")
	mgr = renderer.New(nil)
)


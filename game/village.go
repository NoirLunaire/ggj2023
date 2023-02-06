package game

import (
)

const (
	None = iota
	Tower
	Tavern
	Church
)

type Village struct {
	TabBuild	[][2]int 
	TabPositionBuild	[][2]float64
}

func NewVillage () *Village {
	var TabPositionBuild  [][2]float64
	TabPositionBuild = append(TabPositionBuild,[2]float64{1250, 400})
	
	return &Village{
		[][2]int{},
		TabPositionBuild,
	}
}


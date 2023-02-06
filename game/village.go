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
	TabPositionBuild	[][4]float64
}

func NewVillage () *Village {
	var TabPositionBuild  [][4]float64

	TabPositionBuild = append(TabPositionBuild,[4]float64{1100, 350, 0.50,0.50})
	TabPositionBuild = append(TabPositionBuild,[4]float64{1250,400, 0.50,0.50})
	TabPositionBuild = append(TabPositionBuild,[4]float64{900,200, 0.25,0.25})
	return &Village{
		[][2]int{},
		TabPositionBuild,
	}
}


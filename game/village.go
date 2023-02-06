package game

import (
	"math/rand"
	"time"
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

func shuffle(data [][4]float64) [][4]float64 {
	rand.Seed(time.Now().UnixNano())
	n := len(data)
	for i := n - 1; i > 0; i-- {
		j := rand.Intn(i + 1)
		data[i], data[j] = data[j], data[i]
	}
	return data
}

func NewVillage () *Village {
	var TabPositionBuild  [][4]float64

	TabPositionBuild = append(TabPositionBuild,[4]float64{1100, 350, 0.50,0.50})
	TabPositionBuild = append(TabPositionBuild,[4]float64{1250,400, 0.50,0.50})
	TabPositionBuild = append(TabPositionBuild,[4]float64{900,200, 0.25,0.25})
	TabPositionBuild = append(TabPositionBuild,[4]float64{0,400, 0.50,0.50})
	TabPositionBuild = append(TabPositionBuild,[4]float64{50,300, 0.50,0.50})
	TabPositionBuild = append(TabPositionBuild,[4]float64{-50,400, 0.50,0.50})
	TabPositionBuild = append(TabPositionBuild,[4]float64{300,200, 0.25,0.25})
	shuffle(TabPositionBuild) 
	
	return &Village{
		[][2]int{},
		TabPositionBuild,
	}
}


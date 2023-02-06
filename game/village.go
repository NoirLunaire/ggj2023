package game

import (
	"math/rand"
	"time"
	"log"

	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2"
)

const (
	Tower = iota
	House
	Tavern
	Church
)

type Village struct {
	TabBuild	[][2]int 
	TabPositionBuild	[][4]float64
	TabImg	[][]*ebiten.Image
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
	
	imgTower0, _, err := ebitenutil.NewImageFromFile("data/image/tower0.png")
	imgTower1, _, erre := ebitenutil.NewImageFromFile("data/image/tower1.png")
	imgTower2, _, errer := ebitenutil.NewImageFromFile("data/image/tower2.png")
	imgHouse1, _, errere := ebitenutil.NewImageFromFile("data/image/House1.png")
	if err != nil || erre != nil || errer != nil || errere != nil{
		log.Fatalf("Failed to village build load image: %v", err)
	}

	TabImg := make([][]*ebiten.Image, 3)
	TabImg[0] = []*ebiten.Image{imgTower0, imgTower1, imgTower2}
	TabImg[1] = []*ebiten.Image{nil, imgHouse1, nil}
	
	return &Village{
		[][2]int{},
		TabPositionBuild,
		TabImg,
	}
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

func ChooseBuildImg(villa *Village, build [2]int) *ebiten.Image {
	return villa.TabImg[build[0]][build[1]]
}

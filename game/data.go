package game

import (
	"time"
)

type State struct {
	King_age	int
	Date		time.Time

	Happiness	int
	Money		int
	Population	int
	EventPool	[]int

	EventList	[]*Event
	ChoiceList	[]*Choice
	Effects		map[int]func(s *State)
}

func NewState () *State {
	return &State{
		18,
		time.Date(1000, time.January, 1, 12, 0, 0, 0, time.UTC),
		10,
		10,
		10,
		[]int{ 1 },
		LoadEvents(),
		LoadChoices(),
		LoadEffects(),
	}
}

func LoadEffects () map[int]func(s *State) {
	m := make(map[int]func(s *State))
	m[0] = LosePop
	m[1] = LoseMoney
	m[2] = WinHap
	m[3] = HapForMoney
	return m
}

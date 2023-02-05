package game

func Nothing (s *State) {}

// 10 should add event
// H M P

func addEvents (s *State) {
	s.EventPool = []int{ 1, 2, 3, 4 }
}

func P1Hap (s *State) {
	s.Happiness += 1
}

func M2HapP1Mon (s *State) {
	s.Happiness -= 2
	s.Money += 1
}

func M1HapP3Pop (s *State) {
	s.Happiness -= 1
	s.Population += 3
}

func P1HapP1Pop (s *State) {
	s.Population += 1
	s.Happiness += 1
}

func M1HapP1Mon (s *State) {
	s.Money += 1
	s.Happiness -= 1
}

func M1Hap (s *State) {
	s.Happiness -= 1
}

func M1Mon (s *State) {
	s.Money -= 1
}

func P1HapM2Mon (s *State) {
	s.Money -= 2
	s.Happiness += 1
}

func P1HapM1Mon (s *State) {
	s.Money -= 1
	s.Happiness += 1
}

func M5Mon (s *State) {
	s.Money -= 5
}

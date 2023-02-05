package game

func LosePop (s *State) {
	s.Population -= 3
}

func LoseMoney (s *State) {
	s.Money -= 3
}

func LoseHap (s *State) {
	s.Happiness -= 3
}

func WinHap (s *State) {
	s.Happiness += 1
}

func HapForMoney (s *State) {
	s.Happiness -= 2
	s.Money += 1
}

func Bienvenue (s *State) {
	s.Happiness -= 1
} 

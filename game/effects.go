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
 

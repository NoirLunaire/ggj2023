package game

type State struct {
	King_age	int
	Year		int	

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
		1000,
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

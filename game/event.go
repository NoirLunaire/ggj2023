package game

type Event struct {
	Id		int
	Title		string
	Description	string
	Choices		[]int
}

func LoadEvents () []*Event {
	return []*Event{}	
}

package game

import (
	"log"
	"errors"
	"strings"
	"os"
	"fmt"
	"strconv"
	"time"
)

func CheckError (e error) {
	if e != nil {
		log.Fatal(e)
	}
}

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
		[]int{ 0, 1, 2, 3 },
		LoadEvents(),
		LoadChoices(),
		LoadEffects(),
	}
}

func LoadEffects () map[int]func(s *State) {
	m := make(map[int]func(s *State))
	m[0] = Nothing
	m[1] = Nothing
	m[2] = P1Hap
	m[3] = M2HapP1Mon
	m[4] = M1HapP3Pop
	m[5] = P1HapP1Pop
	m[6] = M1HapP1Mon
	m[7] = M1Hap
	m[8] = M1Mon
	m[9] = Nothing
	return m
}

func SaveGame (name string, s *State) {
	path := "save/" + name + ".sav"
	f, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE, 0644)
	CheckError(err)
	err = f.Truncate(0)
	CheckError(err)
	save := ""
	save += strconv.Itoa(s.King_age) + ";"
	save += s.Date.Format("02-01-2006") + ";"
	save += strconv.Itoa(s.Happiness) + ";"
	save += strconv.Itoa(s.Money) + ";"
	save += strconv.Itoa(s.Population) + ";"
	for i := 0; i < len(s.EventPool); i++ {
		fmt.Println("saving ", s.EventPool[i])
		save += strconv.Itoa(s.EventPool[i]) + ";"
	}

	_, err = f.WriteString(save)
	if err != nil {
		log.Fatal(err)
	}
}

func LoadSave (name string) (*State, error)  {
	path := "save/" + name
	f, err := os.Open(path)
	if err != nil { return nil, err }
	info, err := f.Stat()
	if err != nil { return nil, err }
	size := info.Size()
	list := make([]byte, size)
	_, err = f.Read(list)
	if err != nil {
		return nil, err
	}

	data := string(list)
	tab := strings.Split(data, ";")
	if len(tab) <= 5 {
		log.Fatal("Corrupted file")
	}
	fmt.Println(tab)
	var state State
	state.King_age, err = strconv.Atoi(tab[0])
	if err != nil {
		return nil, err
	}

	state.Date, err = time.Parse("02-01-2006", tab[1])
	if err != nil {
		return nil, err
	}

	state.Happiness, err = strconv.Atoi(tab[2])
	if err != nil {
		return nil, err
	}

	state.Money, err = strconv.Atoi(tab[3])

	state.Population, err = strconv.Atoi(tab[4])
	if err != nil {
		return nil, err
	}
	for i := 5; i < len(tab) - 1; i++ {
		j, err := strconv.Atoi(tab[i])
		if err != nil {
			return nil, errors.New("Invalid save")
		}
		state.EventPool = append(state.EventPool, j)
	}

	state.EventList = LoadEvents()
	state.ChoiceList = LoadChoices()
	state.Effects = LoadEffects()
	return &state, nil
}

func GetSaves () []string {
	path := "save/"
	saves := []string{}
	f, err := os.Open(path)
	CheckError(err)
	c, err := f.ReadDir(0)
	CheckError(err)
	for i := 0; i < len(c); i++ {
		ext := strings.Split(c[i].Name(), ".")
		if len(ext) < 2 {
			continue
		}
		if ext[1] != "sav" {
			continue
		}
		saves = append(saves, c[i].Name())
	}
	return saves
}


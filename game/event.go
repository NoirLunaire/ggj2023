package game

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)


type Event struct {
	Id		int
	Title		string
	Description	string
	Choices		[]int
}

func LoadEvents () []*Event {
	file, err := os.Open("../data/dialog/events.csv")
	if err != nil {
		fmt.Println("Error while trying opening csv file (events.csv):", err)
		return nil
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		fmt.Println("Error reading CSV (events.csv):", err)
		return nil
	}	
	
	tab := []*Event{}

	for _, value := range records {
		tabChoice := []int{}
	
		for i := 3 ; i < len(value); i++{
			casted, err := strconv.Atoi(value[i])
			if err != nil {
				fmt.Println("Error converting string to int:", err)
				return nil
			}
			tabChoice = append(tabChoice,casted)
		}

		id, err := strconv.Atoi(value[0])
			if err != nil {
				fmt.Println("Error converting string to int:", err)
				return nil
			}
		
		tab = append(tab,&Event{
			id,
			value[1],
			value[2],
			tabChoice,
		}) 
	}

	return tab	
}

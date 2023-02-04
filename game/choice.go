package game

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)


type Choice struct {
	Id		int
	Title		string
	Description	string
}

func LoadChoices () []*Choice {
	file, err := os.Open("data/dialog/choices.csv")
	if err != nil {
		fmt.Println("Error while trying opening csv file (choices.csv):", err)
		return nil
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		fmt.Println("Error reading CSV (choices.csv):", err)
		return nil
	}	
	
	tab := []*Choice{}

	for _, value := range records {
		id, err := strconv.Atoi(value[0])
			if err != nil {
				fmt.Println("Error converting string to int:", err)
				return nil
			}
		
		tab = append(tab,&Choice{
			id,
			value[1],
			value[2],
		}) 
	}

	return tab	
}

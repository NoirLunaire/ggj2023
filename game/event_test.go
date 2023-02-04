package game

import (
	"testing"
	"fmt"
)

func TestParser(t *testing.T) {
    //t.Errorf("Abs(-1) = %d; want 1", got)
	tabEvent := LoadEvents()
	fmt.Println(tabEvent[0])

	tabChoice := LoadChoices()
	fmt.Println(tabChoice[0])
	if len(tabEvent[0].Choices) < 2  {
		t.Errorf("Wrong length of choice %d", tabEvent[0].Id)
		return
	}
	
}
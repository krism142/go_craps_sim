package player

import (
	"github.com/krism142/go_craps_sim/dice"
	"os"
	"fmt"
)

var dice1,dice2,result,previous int
var comingOut = true
var err error

func TakeTurn() error{
	dice1,dice2,err = dice.CrapsRoll()
	result = dice1+dice2
	if (result == 7 || result == 11) && comingOut {
		fmt.Fprintf(os.Stdout, "Winner winner! You rolled a %v on the come out\n", result)
		comingOut = true
		return nil
	} else if result == 7 && !comingOut {
		fmt.Fprintf(os.Stdout, "Looks like your turn is over, you rolled a 7\n")
		comingOut = true
		previous = 0
		return nil
	} else if (result == 2 || result == 3 || result == 12) && comingOut {
		fmt.Fprintf(os.Stdout, "Craps! you rolled a %v on the come out\n", result)
		comingOut = true
		return nil
	} else {
		if comingOut {
			fmt.Fprintf(os.Stdout, "You rolled a %v, let's see if you can get another before a 7!\n", result)
			comingOut = false
			previous = result
		} else {
			if previous == result {
				fmt.Fprintf(os.Stdout, "Congrats we have a winner you rolled another %v!\n", result)
				comingOut = true
			} else {
				fmt.Fprintf(os.Stdout, "No winner this time you rolled a %v, try again.\n", result)
			}
		}
		return nil
	}
}

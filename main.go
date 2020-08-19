package main

import (
  "os"
	"fmt"
	"github.com/krism142/go_craps_sim/player"
)

func main() {
  var stillTurn = true
	var err error
	for {
		if stillTurn {
			stillTurn,err = player.TakeTurn()
		}else{
			break
		}
		if err != nil {
			fmt.Fprintf(os.Stdout, "There was an error, it was \n%v\n", err)
		}
	}
}

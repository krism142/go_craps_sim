package main

import (
  "fmt"
  "os"
	"strconv"
  "github.com/krism142/go_craps_sim/dice"
)

func main() {
	rolls,err := strconv.Atoi(os.Getenv("dice_rolls"))
	if err != nil {
		return
	}
	var i,dice1,dice2,result int
	var nums [13]int
	for i=0; i<int(rolls); i++ {
		dice1,dice2,err = dice.CrapsRoll()
		result = dice1+dice2
		nums[result] = nums[result]+1
		//fmt.Fprintf(os.Stdout, "Dice1: %v, Dice2: %v, Total: %v\n", dice1, dice2, result)
	}
	for i=0; i<13; i++ {
		fmt.Fprintf(os.Stdout, "%v : %v\n", i, nums[i])
	}
}

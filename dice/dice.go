package dice

import (
  "time"
  "math/rand"
	"os"
	"strconv"
)

func roll() (int,error) {
	sides, err := strconv.Atoi(os.Getenv("dice_sides"))
	if err != nil {
		return 0, err
	}
  rand.Seed(time.Now().UnixNano())
  num := rand.Intn(sides)+1
  return num,nil
}

func CrapsRoll() (int, int, error) {
  dice1,err := roll()
	if err != nil {
		return 0, 0, err
	}
  dice2,err := roll()
	if err != nil {
		return 0, 0, err
	}

  return dice1,dice2,nil
}

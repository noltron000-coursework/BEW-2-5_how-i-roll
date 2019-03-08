package roller

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

// parseDice input should be a string `XdY`.
func parseDice(dice string) (int, int) {

	var diceNum string
	var diceSize string
	var passType bool

	for _, char := range dice {
		if char == 'd' && !passType {
			passType = true
		} else if char == 'd' && passType {
			fmt.Println("ERROR")
		} else {
			// check if item is number
			var found = false
			for _, num := range "0123456789" {
				if num == char {
					if passType {
						diceSize += string(char)
					} else {
						diceNum += string(char)
					}
					found = true
					break
				}
			}
			if !found {
				fmt.Println("ERROR")
			}
		}
	}

	var outNum, _ = strconv.Atoi(diceNum)
	var outSize, _ = strconv.Atoi(diceSize)
	return outNum, outSize
}

// randDice inputs two ints & outputs one random int
func randDice(diceNum int, diceSize int) int {
	rand.Seed(time.Now().UnixNano())

	var counter int
	var result int

	for counter < diceNum {
		var diceRoll = rand.Intn(diceSize) + 1
		result += diceRoll
		counter++
	}
	return result
}

// RollDice parses dice string, then computes the roll.
func RollDice(dice string) int {
	var result = randDice(parseDice(dice))
	return result
}

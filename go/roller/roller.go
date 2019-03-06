package roller

import (
	"fmt"
	"math/rand"
	"strconv"
)

// parseDice input should be a string `XdY`.
func parseDice(dice string) (int, int) {

	var diceNum string
	var diceSize string
	var passType bool

	fmt.Println("boolean:")
	fmt.Println(passType)

	for _, char := range dice {
		if char == 'd' && !passType {
			passType = true
		} else if char == 'd' && passType {
			return 0, 0
		} else {
			// check if item is number
			for _, num := range "0123456789" {
				if num == char {
					if passType {
						diceSize += string(char)
					} else {
						diceNum += string(char)
					}
					break
				} else {
					return 0, 0
				}
			}
		}
	}

	var outNum, _ = strconv.Atoi(diceNum)
	var outSize, _ = strconv.Atoi(diceSize)
	return outNum, outSize
}

// rollDice inputs two ints & outputs one random int.
func rollDice(diceNum int, diceSize int) int {

	var counter int
	var result int

	for counter < diceNum {
		var diceRoll = rand.Intn(diceSize)
		result += diceRoll
		counter++
	}
	return result
}

// main runs rollDice on parseDice of input.
func main(dice string) int {
	return rollDice(parseDice(dice))
}

package main

import (
	"fmt"
	"utils"
)

type Opcode int

const (
	ADD  Opcode = 1
	MULT Opcode = 2
	STOP Opcode = 99
)

func main() {
	var input = utils.GetInput("src/day2/input.txt")

	for noun := 0; noun <= 99; noun++ {
		for verb := 0; verb <= 99; verb++ {
			attempt := make([]int, len(input))
			copy(attempt, input)
			attempt[1] = noun
			attempt[2] = verb

			output := getOutput(attempt)

			if output == 19690720 {
				fmt.Println(output)
				fmt.Println(100*noun + verb)
				break
			}
		}
	}
}

func getOutput(input []int) int {
inputLoop:
	for i := 0; i < len(input); i += 4 {
		var opcode = Opcode(input[i])

		switch opcode {
		case ADD:
			aValue, bValue, cValue := getValues(input, i)
			input[cValue] = aValue + bValue

		case MULT:
			aValue, bValue, cValue := getValues(input, i)
			input[cValue] = aValue * bValue

		case STOP:
			break inputLoop

		default:
			panic(fmt.Sprintf("Unknown opcode %v", opcode))
		}
	}

	return input[0]
}

func getValues(input []int, index int) (int, int, int) {
	var a = input[index+1]
	var b = input[index+2]
	var aValue = input[a]
	var bValue = input[b]
	var cValue = input[index+3]

	return aValue, bValue, cValue
}

// 4023471

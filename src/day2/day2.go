package main

import (
	"fmt"
	"utils"
)

type Opcode int

var input = utils.GetInput("src/day2/input.txt")

const (
	ADD  Opcode = 1
	MULT Opcode = 2
	STOP Opcode = 99
)

func main() {
inputLoop:
	for i := 0; i < len(input); i += 4 {
		var opcode = Opcode(input[i])

		switch opcode {
		case ADD:
			aValue, bValue, cValue := getValues(i)
			input[cValue] = aValue + bValue

		case MULT:
			aValue, bValue, cValue := getValues(i)
			input[cValue] = aValue * bValue

		case STOP:
			break inputLoop

		default:
			panic(fmt.Sprintf("Unknown opcode %v", opcode))
		}
	}

	fmt.Println(input[0])
}

func getValues(index int) (int, int, int) {
	var a = input[index+1]
	var b = input[index+2]
	var aValue = input[a]
	var bValue = input[b]
	var cValue = input[index+3]

	return aValue, bValue, cValue
}

// 4023471

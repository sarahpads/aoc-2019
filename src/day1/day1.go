package main

import (
	"fmt"
	"utils"
)

func main() {
	modules := utils.GetInput("src/day1/input.txt")
	var totalFuel int

	for _, element := range modules {
		var requiredFuel = getRequiredFuel(element)

		totalFuel += requiredFuel
	}

	fmt.Println(totalFuel)
}

func getRequiredFuel(mass int) int {
	var fuel = getFuelByMass(mass)
	var remainingMass = fuel

	for remainingMass > 0 {
		var additionalFuel = getFuelByMass(remainingMass)

		if additionalFuel <= 0 {
			break
		}

		fuel += additionalFuel
		remainingMass = additionalFuel
	}

	return fuel
}

func getFuelByMass(mass int) int {
	return int(mass/3) - 2
}

// 5093620

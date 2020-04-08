package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	"strconv"
)

func main() {
	file, err := ioutil.ReadFile("../input.txt")

	if err != nil {
		panic(err)
	}

	modules := strings.Split(string(file), "\n")

	var totalFuel int

	for _, element := range modules {
		var mass, _ = strconv.Atoi(element)
		var requiredFuel = getRequiredFuel(mass)

		totalFuel += requiredFuel
	}

	fmt.Println(totalFuel)
}

func getRequiredFuel(mass int) int {
	var fuel = getFuelByMass(mass);
	var remainingMass = fuel

	for remainingMass > 0 {
		var additionalFuel = getFuelByMass(remainingMass)

		if additionalFuel <= 0 {
			break
		}

		fuel += additionalFuel;
		remainingMass = additionalFuel
	}

	return fuel;
}

func getFuelByMass(mass int) int {
	return int(mass / 3) - 2
}
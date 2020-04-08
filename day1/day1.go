package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	"strconv"
)

func main() {
	file, _ := ioutil.ReadFile("../input.txt")
	modules := strings.Split(string(file), "\n")

	var total int

	for _, element := range modules {
		var fuel, _ = strconv.Atoi(element)
		total += int(fuel / 3) - 2
	}

	fmt.Print(total)
}
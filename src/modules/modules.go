package modules

import (
	"io/ioutil"
	"strconv"
	"strings"
)

func GetModules() []int {
	file, err := ioutil.ReadFile("input.txt")

	if err != nil {
		panic(err)
	}

	rows := strings.Split(string(file), "\n")
	modules := make([]int, len(rows))

	for index, element := range rows {
		module, _ := strconv.Atoi(element)
		modules[index] = module
	}

	return modules
}

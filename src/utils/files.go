package utils

import (
	"io/ioutil"
	"strconv"
	"strings"
)

func GetInput(path string) []int {
	file, err := ioutil.ReadFile(path)

	if err != nil {
		panic(err)
	}

	rows := strings.FieldsFunc(string(file), split)
	values := make([]int, len(rows))

	for index, input := range rows {
		value, _ := strconv.Atoi(strings.TrimSpace(input))
		values[index] = value
	}

	return values
}

func split(r rune) bool {
	return strings.ContainsRune(", \n", r)
}

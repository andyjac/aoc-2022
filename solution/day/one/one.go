package one

import (
	"sort"
	"strconv"
	"strings"
)

func calcCalories(input string) ([]int, error) {
	var calories []int
	var total int
	lines := strings.Split(input, "\n")

	for _, l := range lines {
		if l == "" {
			calories = append(calories, total)
			total = 0
		} else {
			count, err := strconv.Atoi(l)
			if err != nil {
				return calories, err
			}
			total += count
		}
	}

	return calories, nil
}

func PartOne(input string) (string, error) {
	calories, err := calcCalories(input)
	if err != nil {
		return "", err
	}

	sort.Sort(sort.Reverse(sort.IntSlice(calories)))
	return strconv.Itoa(calories[0]), nil
}

func PartTwo(input string) (string, error) {
	calories, err := calcCalories(input)
	if err != nil {
		return "", err
	}

	sort.Sort(sort.Reverse(sort.IntSlice(calories)))
	return strconv.Itoa(calories[0] + calories[1] + calories[2]), nil
}

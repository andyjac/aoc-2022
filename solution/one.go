package solution

import (
	"sort"
	"strconv"
	"strings"

	inputparser "github.com/andyjac/aoc-2022/input_parser"
)

type One struct{}

func (*One) Run() (string, string, error) {
	input, err := inputparser.Parse(ONE)
	if err != nil {
		return "", "", err
	}

	r1, err := one_PartOne(input)
	r2, err := one_PartTwo(input)
	return r1, r2, err
}

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

func one_PartOne(input string) (string, error) {
	calories, err := calcCalories(input)
	if err != nil {
		return "", err
	}

	sort.Sort(sort.Reverse(sort.IntSlice(calories)))
	return strconv.Itoa(calories[0]), nil
}

func one_PartTwo(input string) (string, error) {
	calories, err := calcCalories(input)
	if err != nil {
		return "", err
	}

	sort.Sort(sort.Reverse(sort.IntSlice(calories)))
	return strconv.Itoa(calories[0] + calories[1] + calories[2]), nil
}

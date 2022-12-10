package solution

import inputparser "github.com/andyjac/aoc-2022/input_parser"

type Two struct{}

func (*Two) Run() (string, string, error) {
	input, err := inputparser.Parse(TWO)
	if err != nil {
		return "", "", err
	}

	r1, err := two_PartOne(input)
	r2, err := two_PartTwo(input)

	return r1, r2, err
}

func two_PartOne(input string) (string, error) {
	return "", nil
}

func two_PartTwo(input string) (string, error) {
	return "", nil
}

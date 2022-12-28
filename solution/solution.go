package solution

import (
	"errors"
	"fmt"

	inputparser "github.com/andyjac/aoc-2022/input_parser"
	"github.com/andyjac/aoc-2022/solution/day"
	"github.com/andyjac/aoc-2022/solution/day/five"
	"github.com/andyjac/aoc-2022/solution/day/four"
	"github.com/andyjac/aoc-2022/solution/day/one"
	"github.com/andyjac/aoc-2022/solution/day/seven"
	"github.com/andyjac/aoc-2022/solution/day/six"
	"github.com/andyjac/aoc-2022/solution/day/three"
	"github.com/andyjac/aoc-2022/solution/day/two"
)

var ErrDayNotImplemented = errors.New("day is not implmeneted")

type PartRunner func(input string) (string, error)

type Output struct {
	result string
	err    error
}

var partRunners = map[string][]PartRunner{
	day.One:   {one.PartOne, one.PartTwo},
	day.Two:   {two.PartOne, two.PartTwo},
	day.Three: {three.PartOne, three.PartTwo},
	day.Four:  {four.PartOne, four.PartTwo},
	day.Five:  {five.PartOne, five.PartTwo},
	day.Six:   {six.PartOne, six.PartTwo},
	day.Seven: {seven.PartOne, seven.PartTwo},
}

func Run(day string) {
	output := []Output{}
	parts := partRunners[day]
	if parts == nil {
		fmt.Println(ErrDayNotImplemented)
		return
	}

	input, err := inputparser.Parse(day)
	if err != nil {
		fmt.Println(fmt.Errorf("Failed to parse input: %w", err))
		return
	}

	for _, run := range parts {
		result, err := run(input)

		output = append(output, Output{
			result: result,
			err:    err,
		})
	}

	for i, output := range output {
		if output.err != nil {
			fmt.Printf("error in day %s, part %d: %v\n", day, i+1, output.err)
		} else {
			fmt.Printf("result for day %s, part %d: %s\n", day, i+1, output.result)
		}
	}
}

package four

import (
	"strconv"
	"strings"
)

type Four struct{}

func parseSection(section string) ([]int, error) {
	idRange := strings.Split(section, "-")
	start, err := strconv.Atoi(idRange[0])
	if err != nil {
		return []int{}, err
	}

	end, err := strconv.Atoi(idRange[1])
	if err != nil {
		return []int{}, err
	}

	return []int{start, end}, nil
}

func max(int1 int, int2 int) int {
	if int1 > int2 {
		return int1
	}

	return int2
}

func min(int1 int, int2 int) int {
	if int1 < int2 {
		return int1
	}

	return int2
}

func doRangesOverlap(range1 []int, range2 []int) bool {
	maxStart := max(range1[0], range2[0])
	minEnd := min(range1[1], range2[1])

	return maxStart <= minEnd
}

func isRangeContained(range1 []int, range2 []int) bool {
	if range1[0] <= range2[0] && range1[1] >= range2[1] {
		return true
	}
	if range2[0] <= range1[0] && range2[1] >= range1[1] {
		return true
	}

	return false
}

func PartOne(input string) (string, error) {
	var total int
	lines := strings.Split(input, "\n")

	for _, l := range lines {
		if l == "" {
			continue
		}

		sections := strings.Split(l, ",")
		range1, err := parseSection(sections[0])
		if err != nil {
			return "", err
		}

		range2, err := parseSection(sections[1])
		if err != nil {
			return "", err
		}

		if isRangeContained(range1, range2) {
			total += 1
		}
	}

	return strconv.Itoa(total), nil
}

func PartTwo(input string) (string, error) {
	var total int
	lines := strings.Split(input, "\n")

	for _, l := range lines {
		if l == "" {
			continue
		}

		sections := strings.Split(l, ",")
		range1, err := parseSection(sections[0])
		if err != nil {
			return "", err
		}

		range2, err := parseSection(sections[1])
		if err != nil {
			return "", err
		}

		if doRangesOverlap(range1, range2) {
			total += 1
		}
	}

	return strconv.Itoa(total), nil
}

package three

import (
	"errors"
	"strconv"
	"strings"
)

type Three struct{}

var charPriorityMap = map[string]int{
	"a": 1,
	"b": 2,
	"c": 3,
	"d": 4,
	"e": 5,
	"f": 6,
	"g": 7,
	"h": 8,
	"i": 9,
	"j": 10,
	"k": 11,
	"l": 12,
	"m": 13,
	"n": 14,
	"o": 15,
	"p": 16,
	"q": 17,
	"r": 18,
	"s": 19,
	"t": 20,
	"u": 21,
	"v": 22,
	"w": 23,
	"x": 24,
	"y": 25,
	"z": 26,
	"A": 27,
	"B": 28,
	"C": 29,
	"D": 30,
	"E": 31,
	"F": 32,
	"G": 33,
	"H": 34,
	"I": 35,
	"J": 36,
	"K": 37,
	"L": 38,
	"M": 39,
	"N": 40,
	"O": 41,
	"P": 42,
	"Q": 43,
	"R": 44,
	"S": 45,
	"T": 46,
	"U": 47,
	"V": 48,
	"W": 49,
	"X": 50,
	"Y": 51,
	"Z": 52,
}

func PartOne(input string) (string, error) {
	lines := strings.Split(input, "\n")
	var total int

	for _, l := range lines {
		rucksack1 := l[0 : len(l)/2]
		rucksack2 := l[len(l)/2:]

		for _, c := range strings.Split(rucksack1, "") {
			if strings.Contains(rucksack2, c) {
				total += charPriorityMap[c]
				break
			}
		}
	}

	return strconv.Itoa(total), nil
}

func chunkSlice(slice []string, chunkSize int) [][]string {
	var chunks [][]string

	for i := 0; i < len(slice); i += chunkSize {
		end := i + chunkSize

		if end > len(slice) {
			end = len(slice)
		}

		chunks = append(chunks, slice[i:end])
	}

	return chunks
}

func getGroupBadge(rucksacks []string) (string, error) {
	rucksackBadgeMap := map[int]map[string]bool{
		0: {},
		1: {},
		2: {},
	}

	for i, r := range rucksacks {
		chars := strings.Split(r, "")

		for _, c := range chars {
			rucksackBadgeMap[i][c] = true

			if rucksackBadgeMap[0][c] && rucksackBadgeMap[1][c] && rucksackBadgeMap[2][c] {
				return c, nil
			}
		}
	}

	return "", errors.New("Badge not found")
}

func PartTwo(input string) (string, error) {
	lines := strings.Split(input, "\n")
	var total int

	groups := chunkSlice(lines, 3)
	for _, rucksacks := range groups {
		badge, err := getGroupBadge(rucksacks)

		if err == nil {
			total += charPriorityMap[badge]
		}
	}

	return strconv.Itoa(total), nil
}

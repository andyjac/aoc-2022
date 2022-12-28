package six

import (
	"strconv"
	"strings"
)

type Six struct{}

func isUniqueSegment(segment []string) bool {
	charMap := map[string]bool{}

	for _, c := range segment {
		if charMap[c] {
			return false
		}

		charMap[c] = true
	}

	return true
}

func getUniqueSegmentMarker(input string, size int) int {
	chars := strings.Split(input, "")
	left := 0
	right := size

	for right <= len(chars) {
		segment := chars[left:right]
		if isUniqueSegment(segment) {
			break
		}

		left++
		right++
	}

	return right

}

func PartOne(input string) (string, error) {
	marker := getUniqueSegmentMarker(input, 4)
	return strconv.Itoa(marker), nil
}

func PartTwo(input string) (string, error) {
	marker := getUniqueSegmentMarker(input, 14)
	return strconv.Itoa(marker), nil
}

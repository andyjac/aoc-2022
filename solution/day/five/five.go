package five

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type Five struct{}

type Instruction struct {
	count int
	src   int
	dest  int
}

func createGrid(input string) [][]string {
	grid := [][]string{}
	lines := strings.Split(input, "\n")

	for _, l := range lines {
		// Exit case: End of map section
		if l == "" {
			break
		}

		chars := strings.Split(l, "")
		grid = append(grid, chars)
	}

	return grid
}

func parseId(v string) (int, error) {
	id, err := strconv.Atoi(v)
	if err != nil {
		return id, err
	}

	return id, nil
}

func getMaybeCrateAtIdx(row []string, idx int) string {
	if len(row) >= idx {
		crate := row[idx]

		if strings.TrimSpace(crate) != "" {
			return crate
		}
	}

	return ""
}

func parseCrateMap(input string) map[int][]string {
	var crateMap = map[int][]string{}

	grid := createGrid(input)
	idRowIdx := len(grid) - 1
	idRow := grid[idRowIdx]

	for i, v := range idRow {
		id, err := parseId(v)
		if err != nil {
			continue
		}

		crateMap[id] = []string{}

		rowIdx := idRowIdx - 1
		for rowIdx >= 0 {
			row := grid[rowIdx]
			crate := getMaybeCrateAtIdx(row, i)
			if crate != "" {
				crateMap[id] = append(crateMap[id], crate)
			}

			rowIdx--
		}
	}

	return crateMap
}

func parseInstructions(input string) []string {
	var instructionStartIdx int

	lines := strings.Split(input, "\n")
	for i, l := range lines {
		if l == "" {
			instructionStartIdx = i + 1
			break
		}
	}

	return lines[instructionStartIdx:]
}

func parseInstruction(line string) (*Instruction, error) {
	if line == "" {
		return &Instruction{}, errors.New("line is blank")
	}

	//       *            *        *
	// [0]  [1]     [2]  [3]   [4][5]
	// move <count> from <src> to <dest>
	parts := strings.Split(line, " ")
	count, err := strconv.Atoi(parts[1])
	if err != nil {
		return &Instruction{}, err
	}

	src, err := strconv.Atoi(parts[3])
	if err != nil {
		return &Instruction{}, err
	}

	dest, err := strconv.Atoi(parts[5])
	if err != nil {
		return &Instruction{}, err
	}

	return &Instruction{
		count: count,
		src:   src,
		dest:  dest,
	}, nil
}

func top(crates []string) string {
	return crates[len(crates)-1]
}

func pop(crates []string) (string, []string) {
	return top(crates), crates[:len(crates)-1]
}

func getCrateMapKeys(crateMap map[int][]string) []int {
	keys := make([]int, len(crateMap))

	for k := range crateMap {
		keys[k-1] = k
	}

	return keys
}

func getTopCrates(crateMap map[int][]string) []string {
	keys := getCrateMapKeys(crateMap)
	crates := []string{}

	for _, k := range keys {
		crates = append(crates, top(crateMap[k]))
	}

	return crates
}

func printCrateMap(crateMap map[int][]string) {
	var output string
	keys := getCrateMapKeys(crateMap)

	for _, k := range keys {
		output += fmt.Sprintf("%d:", k)

		for _, c := range crateMap[k] {
			output += fmt.Sprintf(" [%s]", c)
		}

		output += "\n"
	}

	fmt.Print(output)
}

func PartOne(input string) (string, error) {
	crateMap := parseCrateMap(input)
	instructions := parseInstructions(input)

	for _, line := range instructions {
		instruction, err := parseInstruction(line)
		if err != nil {
			continue
		}

		c := instruction.count
		for c > 0 {
			crate, newStack := pop(crateMap[instruction.src])
			crateMap[instruction.src] = newStack
			crateMap[instruction.dest] = append(crateMap[instruction.dest], crate)

			c--
		}
	}

	topCrates := getTopCrates(crateMap)
	return strings.Join(topCrates, ""), nil
}

func PartTwo(input string) (string, error) {
	crateMap := parseCrateMap(input)
	instructions := parseInstructions(input)

	for _, line := range instructions {
		instruction, err := parseInstruction(line)
		if err != nil {
			continue
		}

		// The stack to move _from_
		srcStack := crateMap[instruction.src]
		// The stack to move _to_
		destStack := crateMap[instruction.dest]
		// The index from which to move the src crates
		moveIdx := len(srcStack) - instruction.count
		// The top `instruction.count` crates on the stack
		cratesToMove := srcStack[moveIdx:]

		// Pull the required number of crates _off_ of the src
		crateMap[instruction.src] = srcStack[:moveIdx]
		// Put the required number of crates _on_ the dest
		crateMap[instruction.dest] = append(destStack, cratesToMove...)
	}

	topCrates := getTopCrates(crateMap)
	return strings.Join(topCrates, ""), nil
}

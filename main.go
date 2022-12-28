package main

import (
	"fmt"

	"github.com/andyjac/aoc-2022/solution"
)

func main() {
	var day string
	var err error

	fmt.Print("Choose a day: ")
	_, err = fmt.Scanln(&day)

	if err != nil {
		fmt.Println("failed to read input:", err)
		return
	}

	if day == "" {
		fmt.Println("day must not be empty")
		return
	}

	solution.Run(day)
}

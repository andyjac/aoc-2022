package main

import (
	"fmt"

	"github.com/andyjac/aoc-2022/solution"
)

func main() {
	var p string
	var err error

	fmt.Print("Choose a problem: ")
	_, err = fmt.Scanln(&p)

	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}

	if p == "" {
		fmt.Println("Problem must not be empty")
		return
	}

	s := solution.Solutions[p]
	if s == nil {
		fmt.Println("Solution not found:", p)
		return
	}

	p1, p2, err := s.Run()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("%s.1: %s\n", p, p1)
	fmt.Printf("%s.2: %s\n", p, p2)
}

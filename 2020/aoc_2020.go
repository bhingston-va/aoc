package main

import (
	"fmt"
	"github.com/bhingston-va/aoc/2020/q2"
	"os"
)

func main() {
	input, err := q2.Input("")
	if err != nil {
		fmt.Errorf("failed to get input: %s", err.Error())
		os.Exit(-1)
	}
	a2a := q2.Part1(input)
	fmt.Println("answer to Q2 part 1:")
	fmt.Println(a2a)
	a2b := q2.Part2(input)
	fmt.Println("answer to Q2 part 2:")
	fmt.Println(a2b)

	os.Exit(0)
}

package main

import (
	"fmt"
	"github.com/bhingston-va/aoc/2020/q2"
	"github.com/bhingston-va/aoc/2020/q3"
	"os"
)

func question2() {
	input, err := q2.Input("")
	if err != nil {
		fmt.Errorf("failed to get input: %s", err.Error())
		os.Exit(-1)
	}
	a := q2.Part1(input)
	fmt.Println("answer to Q2 part 1:")
	fmt.Println(a)
	b := q2.Part2(input)
	fmt.Println("answer to Q2 part 2:")
	fmt.Println(b)
}

func question3() {
	input, err := q3.Input("")
	if err != nil {
		fmt.Errorf("failed to get input: %s", err.Error())
		os.Exit(-1)
	}
	fmt.Println("answer to Q3 part 1:")
	a := q3.Part1(input)
	fmt.Println(a)
	fmt.Println("answer to Q3 part 2:")
	b := q3.Part2(input)
	fmt.Println(b)
}

func main() {
	//question2()
	question3()

	os.Exit(0)
}

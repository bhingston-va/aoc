package q2

import (
	"bufio"
	"fmt"
	"github.com/bhingston-va/aoc/2020/fileio"
	"os"
	"strconv"
	"strings"
)

/**
--- Day 2: Password Philosophy ---
Your flight departs in a few days from the coastal airport; the easiest way down
to the coast from here is via toboggan.

The shopkeeper at the North Pole Toboggan Rental Shop is having a bad day.
"Something's wrong with our computers; we can't log in!" You ask if you can take
a look.

Their password database seems to be a little corrupted: some of the passwords
wouldn't have been allowed by the Official Toboggan Corporate Policy that was in
effect when they were chosen.

To try to debug the problem, they have created a list (your puzzle input) of
passwords (according to the corrupted database) and the corporate policy when
that password was set.

For example, suppose you have the following list:
1-3 a: abcde
1-3 b: cdefg
2-9 c: ccccccccc

Each line gives the password policy and then the password. The password policy
indicates the lowest and highest number of times a given letter must appear for
the password to be valid. For example, 1-3 a means that the password must
contain a at least 1 time and at most 3 times.

In the above example, 2 passwords are valid. The middle password, cdefg, is not;
it contains no instances of b, but needs at least 1. The first and third
passwords are valid: they contain one a or nine c, both within the limits of
their respective policies.

How many passwords are valid according to their policies?


--- Part Two ---
While it appears you validated the passwords correctly, they don't seem to be
what the Official Toboggan Corporate Authentication System is expecting.

The shopkeeper suddenly realizes that he just accidentally explained the
password policy rules from his old job at the sled rental place down the street!
The Official Toboggan Corporate Policy actually works a little differently.

Each policy actually describes two positions in the password, where 1 means the
first character, 2 means the second character, and so on. (Be careful; Toboggan
Corporate Policies have no concept of "index zero"!) Exactly one of these
positions must contain the given letter. Other occurrences of the letter are
irrelevant for the purposes of policy enforcement.

Given the same example list from above:
1-3 a: abcde is valid: position 1 contains a and position 3 does not.
1-3 b: cdefg is invalid: neither position 1 nor position 3 contains b.
2-9 c: ccccccccc is invalid: both position 2 and position 9 contain c.

How many passwords are valid according to the new interpretation of the
policies?
*/

var valid_1 = passwordPolicy{Min: 1, Max: 3, Char: "a", Pswd: "abcde"}
var invalid = passwordPolicy{Min: 1, Max: 3, Char: "b", Pswd: "cdefg"}
var valid_2 = passwordPolicy{Min: 2, Max: 9, Char: "c", Pswd: "ccccccccc"}

func testInput() []passwordPolicy {
	return []passwordPolicy{valid_1, invalid, valid_2}
}

type passwordPolicy struct {
	Min  int
	Max  int
	Char string
	Pswd string
}

// line format "{min}-{max} {char}: {password}"
func lineToPasswordPolicy(line string) passwordPolicy {
	parts := strings.Split(line, " ")

	minMax := strings.Split(parts[0], "-")
	min, _ := strconv.Atoi(minMax[0])
	max, _ := strconv.Atoi(minMax[1])

	letter := strings.ReplaceAll(parts[1], ":", "")
	pswd := parts[2]

	return passwordPolicy{
		Min:  min,
		Max:  max,
		Char: letter,
		Pswd: pswd,
	}
}

// readQ1Lines is to gut check the fileio go package (anwser done in python)
func readQ1Lines() ([]int, error) {
	return fileio.ReadLinesToInt("./q1/input-a.txt", strconv.Atoi)
}

func TestFileReader() error {
	q1_input, err := readQ1Lines()
	if err != nil {
		return err
	}
	fmt.Printf("first input: %d\n", q1_input[0])
	return nil
}

func isValid(pp passwordPolicy) bool {
	c := strings.Count(pp.Pswd, pp.Char)
	return c >= pp.Min && c <= pp.Max
}

func isValid2(pp passwordPolicy) bool {
	pos1 := pp.Min - 1
	pos2 := pp.Max - 1
	p1IsChar := string(pp.Pswd[pos1]) == pp.Char
	p2IsChar := string(pp.Pswd[pos2]) == pp.Char
	if p1IsChar && p2IsChar {
		return false
	}
	if p1IsChar || p2IsChar {
		return true
	}
	return false
}

func Input(path string) ([]passwordPolicy, error) {
	p := "./q2/input.txt"
	if path != "" {
		p = path
	}
	f, err := os.Open(p)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	s := bufio.NewScanner(f)
	var ls []passwordPolicy
	for s.Scan() {
		ln := s.Text()
		pp := lineToPasswordPolicy(ln)
		ls = append(ls, pp)
	}
	return ls, s.Err()
}

// How many passwords are valid?
func Part1(input []passwordPolicy) int {
	n := 0
	for _, v := range input {
		if isValid(v) {
			n += 1
		}
	}

	return n
}

//How many passwords are valid according to the new interpretation of the policies?
func Part2(input []passwordPolicy) int {
	n := 0
	for _, v := range input {
		if isValid2(v) {
			n += 1
		}
	}

	return n
}

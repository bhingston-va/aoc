package q3

import (
	"bufio"
	"os"
)

func testInput() Navigator {
	return &Grid{
		rows: []row{
			{tSeq: "..##......."},
			{tSeq: "#...#...#.."},
			{tSeq: ".#....#..#."},
			{tSeq: "..#.#...#.#"},
			{tSeq: ".#...##..#."},
			{tSeq: "..#.##....."},
			{tSeq: ".#.#.#....#"},
			{tSeq: ".#........#"},
			{tSeq: "#.##...#..."},
			{tSeq: "#...##....#"},
			{tSeq: ".#..#...#.#"},
		},
	}
}

func Input(path string) (Navigator, error) {
	p := "./q3/input.txt"
	if path != "" {
		p = path
	}
	f, err := os.Open(p)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	var rows []row
	s := bufio.NewScanner(f)
	for s.Scan() {
		ln := s.Text()
		seq := row{tSeq: ln}
		rows = append(rows, seq)
	}

	g := &Grid{rows: rows}
	return g, s.Err()
}

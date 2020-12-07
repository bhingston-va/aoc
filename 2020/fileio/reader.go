package fileio

import (
	"bufio"
	"fmt"
	"os"
)

type Atoi func(line string) (out int, err error)

type scanner struct {
	buffer *bufio.Scanner
}

func newReader(path string) (*scanner, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	return &scanner{
		buffer: bufio.NewScanner(f),
	}, nil
}

func ReadLines(path string) ([]string, error) {
	r, err := newReader(path)
	if err != nil {
		return nil, err
	}

	var ls []string
	for r.buffer.Scan() {
		ls = append(ls, r.buffer.Text())
	}
	return ls, r.buffer.Err()
}

func ReadLinesToInt(path string, atoi Atoi) ([]int, error) {
	r, err := newReader(path)
	if err != nil {
		return nil, err
	}

	var ls []int
	for r.buffer.Scan() {
		ln := r.buffer.Text()
		out, err := atoi(ln)
		if err != nil {
			return nil, fmt.Errorf("failed to convert line: %w", err)
		}
		ls = append(ls, out)
	}
	return ls, r.buffer.Err()

}

func Scan(path string) (*bufio.Scanner, error) {
	r, err := newReader(path)
	if err != nil {
		return nil, err
	}

	return r.buffer, nil
}

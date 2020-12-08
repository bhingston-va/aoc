package q3

/**
--- Day 3: Toboggan Trajectory ---
With the toboggan login problems resolved, you set off toward the airport. While
travel by toboggan might be easy, it's certainly not safe: there's very minimal
steering and the area is covered in trees. You'll need to see which angles will
take you near the fewest trees.

Due to the local geology, trees in this area only grow on exact integer
coordinates in a grid. You make a map (your puzzle input) of the open squares
(.) and trees (#) you can see. For example:

..##.......
#...#...#..
.#....#..#.
..#.#...#.#
.#...##..#.
..#.##.....
.#.#.#....#
.#........#
#.##...#...
#...##....#
.#..#...#.#

These aren't the only trees, though; due to something you read about once
involving arboreal genetics and biome stability, the same pattern repeats to the
right many times:

..##.........##.........##.........##.........##.........##.......  --->
#...#...#..#...#...#..#...#...#..#...#...#..#...#...#..#...#...#..
.#....#..#..#....#..#..#....#..#..#....#..#..#....#..#..#....#..#.
..#.#...#.#..#.#...#.#..#.#...#.#..#.#...#.#..#.#...#.#..#.#...#.#
.#...##..#..#...##..#..#...##..#..#...##..#..#...##..#..#...##..#.
..#.##.......#.##.......#.##.......#.##.......#.##.......#.##.....  --->
.#.#.#....#.#.#.#....#.#.#.#....#.#.#.#....#.#.#.#....#.#.#.#....#
.#........#.#........#.#........#.#........#.#........#.#........#
#.##...#...#.##...#...#.##...#...#.##...#...#.##...#...#.##...#...
#...##....##...##....##...##....##...##....##...##....##...##....#
.#..#...#.#.#..#...#.#.#..#...#.#.#..#...#.#.#..#...#.#.#..#...#.#  --->
You start on the open square (.) in the top-left corner and need to reach the
bottom (below the bottom-most row on your map).

The toboggan can only follow a few specific slopes (you opted for a cheaper
model that prefers rational numbers); start by counting all the trees you would
encounter for the slope right 3, down 1:

From your starting position at the top-left, check the position that is right 3
and down 1. Then, check the position that is right 3 and down 1 from there, and
so on until you go past the bottom of the map.

The locations you'd check in the above example are marked here with O where
there was an open square and X where there was a tree:

..##.........##.........##.........##.........##.........##.......  --->
#..O#...#..#...#...#..#...#...#..#...#...#..#...#...#..#...#...#..
.#....X..#..#....#..#..#....#..#..#....#..#..#....#..#..#....#..#.
..#.#...#O#..#.#...#.#..#.#...#.#..#.#...#.#..#.#...#.#..#.#...#.#
.#...##..#..X...##..#..#...##..#..#...##..#..#...##..#..#...##..#.
..#.##.......#.X#.......#.##.......#.##.......#.##.......#.##.....  --->
.#.#.#....#.#.#.#.O..#.#.#.#....#.#.#.#....#.#.#.#....#.#.#.#....#
.#........#.#........X.#........#.#........#.#........#.#........#
#.##...#...#.##...#...#.X#...#...#.##...#...#.##...#...#.##...#...
#...##....##...##....##...#X....##...##....##...##....##...##....#
.#..#...#.#.#..#...#.#.#..#...X.#.#..#...#.#.#..#...#.#.#..#...#.#  --->

In this example, traversing the map using this slope would cause you to
encounter 7 trees.

Starting at the top-left corner of your map and following a slope of right 3 and
down 1, how many trees would you encounter?
*/

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

type Terrain = string

const (
	tree Terrain = "#"
	open Terrain = "."
)

type Grid struct {
	rows []row
}

type row struct {
	tSeq string
}

func (g *Grid) GetTerrain(crd Coord) Terrain {
	row := g.rows[crd.y]
	x := crd.x % len(g.rows[0].tSeq)
	ter := string(row.tSeq[x])
	return ter
}

func (g *Grid) Height() int {
	return len(g.rows)
}

type Grid2 struct {
	cells      [][]cell
	patternLen int
	height     int
}

type cell struct {
	t Terrain
}

func (g2 *Grid2) GetTerrain(crd Coord) Terrain {
	x := crd.x % g2.patternLen
	cel := g2.cells[x][crd.y]
	return cel.t
}

func (g2 *Grid2) Height() int {
	return g2.height
}

type Navigator interface {
	GetTerrain(crd Coord) Terrain
	Height() int
}

type Coord struct {
	x, y int
}

type Trip struct {
	nav              Navigator
	loc              Coord
	treesEncountered int
}

func NewTrip(nav Navigator) *Trip {
	return &Trip{nav: nav}
}

func (tp *Trip) move() {
	tp.loc.x += 3
	tp.loc.y += 1
	if ter := tp.nav.GetTerrain(tp.loc); ter == tree {
		tp.treesEncountered += 1
	}
}

func (tp *Trip) done() bool {
	return tp.loc.y >= tp.nav.Height()-1
}

func (tp *Trip) Sled() {
	for !tp.done() {
		tp.move()
	}
}

func Part1(input Navigator) int {
	t := NewTrip(input)
	t.Sled()
	return t.treesEncountered
}

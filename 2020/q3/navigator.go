package q3

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

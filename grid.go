package mazer

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

type Grid struct {
	mazeGrid   []*Cell
	rows, cols int
	rng        *rand.Rand
}

func NewGrid(rows, cols int) *Grid {
	g := &Grid{}
	g.rows = rows
	g.cols = cols

	g.mazeGrid = g.createGrid()
	g.initGrid()

	g.rng = rand.New(rand.NewSource(time.Now().Unix()))

	return g
}

func (g *Grid) IJToPos(i, j int) int {
	return i*g.cols + j
}

func (g *Grid) PosToIJ(pos int) (i, j int) {
	return pos / g.cols, pos % g.cols
}

func (g *Grid) createGrid() []*Cell {
	maze := make([]*Cell, g.rows*g.cols)

	for i := 0; i < g.rows; i++ {
		for j := 0; j < g.cols; j++ {
			maze[i*g.cols+j] = NewCell(i, j)
		}
	}
	return maze
}

func (g *Grid) GetCell(row, col int) *Cell {
	if row < 0 || row >= g.rows || col < 0 || col >= g.cols {
		return nil
	}
	return g.mazeGrid[g.IJToPos(row, col)]
}

func (g *Grid) initGrid() {
	for _, c := range g.mazeGrid {
		row, col := c.row, c.col

		c.North = g.GetCell(row-1, col)
		c.East = g.GetCell(row, col+1)
		c.South = g.GetCell(row+1, col)
		c.West = g.GetCell(row, col-1)
	}
}

func (g *Grid) Size() int {
	return g.rows * g.cols
}

func (g *Grid) RandomCell() *Cell {
	row := g.rng.Int() % g.rows
	col := g.rng.Int() % g.cols

	return g.GetCell(row, col)
}

func (g *Grid) Apply(f func(*Cell)) {
	for _, c := range g.mazeGrid {
		f(c)
	}
}

func (g *Grid) PrintASCII() {
	var b strings.Builder

	//Header
	b.WriteString("+")
	for j := 0; j < g.cols; j++ {
		b.WriteString("---+")
	}
	b.WriteString("\n")

	for i := 0; i < g.rows; i++ {
		var top strings.Builder
		var bottom strings.Builder
		top.WriteString("|")
		bottom.WriteString("+")

		for j := 0; j < g.cols; j++ {
			top.WriteString("   ")
			cell := g.GetCell(i, j)

			if cell != nil {
				if cell.IsLinked(cell.East) {
					top.WriteString(" ")
				} else {
					top.WriteString("|")
				}
			} else {
				top.WriteString("|")
			}

			if cell != nil {
				if cell.IsLinked(cell.South) {
					bottom.WriteString("   ")
				} else {
					bottom.WriteString("---")
				}
			} else {
				bottom.WriteString("---")
			}
			bottom.WriteString("+")
		}
		b.WriteString(top.String())
		b.WriteString("\n")
		b.WriteString(bottom.String())
		b.WriteString("\n")
	}
	fmt.Println(b.String())
}

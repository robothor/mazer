package algorithms

import (
	"math/rand"

	"github.com/robothor/mazer"
)

type BinaryTreeMazer struct {
	Rng *rand.Rand
}

func (b *BinaryTreeMazer) binaryTree(cell *mazer.Cell) {
	if b.Rng == nil {
		panic("Must set rng")
	}
	neighbors := []*mazer.Cell{}

	if cell.North != nil {
		neighbors = append(neighbors, cell.North)
	}
	if cell.East != nil {
		neighbors = append(neighbors, cell.East)
	}

	l := len(neighbors)
	var index int
	if l != 0 {
		index = b.Rng.Int() % l
		neighbor := neighbors[index]
		cell.Link(neighbor, true)
	}
}

func (b *BinaryTreeMazer) On(g *mazer.Grid) {
	g.Apply(b.binaryTree)
}

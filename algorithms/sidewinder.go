package algorithms

import (
	"math/rand"

	"github.com/robothor/mazer"
)

type Sidewinder struct {
	Rng *rand.Rand
}

func (s *Sidewinder) sidewinder(cell *mazer.Cell) {
	if s.Rng == nil {
		panic("Must set rng")
	}

}

func (s *Sidewinder) On(g *mazer.Grid) {
	g.Apply(s.sidewinder)
}

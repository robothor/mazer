package mazer

import "testing"

func TestGridSetup(t *testing.T) {
	g := NewGrid(5, 5)
	if g == nil {
		t.Fatal("Unable to construct grid")
	}

}

func TestCoords(t *testing.T) {
	rows := 5
	cols := 5
	g := NewGrid(rows, cols)

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			pos := g.IJToPos(i, j)
			i1, j1 := g.PosToIJ(pos)

			if i != i1 || j != j1 {
				t.Fatalf("Error converting %d, %d -> %d -> %d, %d", i, j, pos, i1, j1)
			}
		}
	}
}

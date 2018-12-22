package mazer

type Cell struct {
	links                    map[*Cell]struct{}
	row, col                 int
	North, East, South, West *Cell
}

func NewCell(row, col int) *Cell {
	c := &Cell{}
	c.row = row
	c.col = col
	c.links = map[*Cell]struct{}{}
	return c
}

func (c *Cell) Link(to *Cell, bidi bool) {
	c.links[to] = struct{}{}
	if bidi {
		to.Link(c, false)
	}
}

func (c *Cell) Unlink(from *Cell, bidi bool) {
	delete(c.links, from)
	if bidi {
		from.Unlink(c, false)
	}
}

func (c *Cell) GetLinks() []*Cell {
	l := []*Cell{}
	for link := range c.links {
		l = append(l, link)
	}
	return l
}

func (c *Cell) IsLinked(to *Cell) bool {
	_, ok := c.links[to]
	return ok
}

func (c *Cell) Row() int {
	return c.row
}

func (c *Cell) Col() int {
	return c.col
}

func (c *Cell) Neighbors() []*Cell {
	l := []*Cell{}
	dirs := []*Cell{c.North, c.East, c.South, c.West}
	for _, dir := range dirs {
		if dir != nil {
			l = append(l, dir)
		}
	}
	return l
}

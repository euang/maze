package maze

type Cell struct {
	Row    int
	Column int
	North  *Cell
	South  *Cell
	West   *Cell
	East   *Cell
	Links  map[*Cell]*Cell
}

func NewCell(row, column int) *Cell {
	c := new(Cell)
	c.Row = row
	c.Column = column
	c.Links = make(map[*Cell]*Cell)
	return c
}

func (c *Cell) Neighbours() []*Cell {
	var neighbours []*Cell
	if c.North != nil {
		neighbours = append(neighbours, c.North)
	}
	if c.South != nil {
		neighbours = append(neighbours, c.South)
	}
	if c.West != nil {
		neighbours = append(neighbours, c.West)
	}
	if c.East != nil {
		neighbours = append(neighbours, c.East)
	}
	return neighbours
}

func (c *Cell) Link(cell *Cell) {
	c.Links[cell] = cell
	cell.LinkOneWay(c)
}

func (c *Cell) LinkOneWay(cell *Cell) {
	c.Links[cell] = cell
}

func (c *Cell) UnLink(cell *Cell) {
	delete(c.Links, cell)

}

func (c *Cell) IsLinked(cell *Cell) bool {
	_, ok := c.Links[cell]
	return ok
}

func (c *Cell) Distances() *Distances {
	distances := NewDistances(c)
	frontier := []*Cell{c}

	for len(frontier) > 0 {
		new_frontier := make([]*Cell, 0)

		for _, cell := range frontier {

			for _, linked := range cell.Links {

				if _, present := distances.Cells[linked]; !present {
					distances.Cells[linked] = distances.Cells[cell] + 1
					new_frontier = append(new_frontier, linked)
				}
			}
		}

		frontier = new_frontier
	}

	return distances
}

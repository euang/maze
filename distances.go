package maze

type Distances struct {
	Cells map[*Cell]int
	root  *Cell
}

func NewDistances(root *Cell) *Distances {
	d := new(Distances)
	d.root = root
	d.Cells = make(map[*Cell]int)
	d.Cells[root] = 0
	return d
}

/*
def [](cell)
@cells[cell]
end

def []=(cell, distance)
@cells[cell] = distance
end
*/
func (d *Distances) AllCells() []*Cell {
	keys := make([]*Cell, len(d.Cells))

	i := 0
	for k := range d.Cells {
		keys[i] = k
		i++
	}

	return keys
}

func (d *Distances) path_to(goal *Cell) *Distances {
	current := goal

	breadcrumbs := NewDistances(d.root)
	breadcrumbs.Cells[current] = d.Cells[current]

	for current != d.root {

		for neighbour := range current.Links {

			if d.Cells[neighbour] < d.Cells[current] {

				breadcrumbs.Cells[neighbour] = d.Cells[neighbour]
				current = neighbour
				break
			}
		}
	}

	return breadcrumbs
}

func (d *Distances) Max() (*Cell, int) {
	max_distance := 0
	max_cell := d.root

	for cell, distance := range d.Cells {
		if distance > max_distance {
			max_cell = cell
			max_distance = distance
		}
	}

	return max_cell, max_distance
}

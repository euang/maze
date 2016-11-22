package maze

import (
	"fmt"
	"math/rand"
	"strings"
)

type Grid struct {
	rows    int
	columns int
	grid    [][]*Cell
}

func NewGrid(rows, cols int) *Grid {
	g := new(Grid)
	g.rows = rows
	g.columns = cols
	g.prepareGrid()
	g.configureCells()
	return g
}

func (g *Grid) Size() int {
	return g.columns * g.rows
}

func (g *Grid) prepareGrid() {
	g.grid = make([][]*Cell, g.rows)
	for i := range g.grid {
		g.grid[i] = make([]*Cell, g.columns)
		for y := range g.grid[i] {
			g.grid[i][y] = NewCell(i, y)
		}
	}
}

func (g *Grid) Cell(row, column int) *Cell {
	if row < 0 || row > g.rows-1 {
		return nil
	}

	if column < 0 || column > g.columns-1 {
		return nil
	}

	return g.grid[row][column]
}

func (g *Grid) configureCells() {
	for i := range g.grid {
		for y := range g.grid[i] {
			row := g.grid[i][y].Row
			col := g.grid[i][y].Column

			g.grid[i][y].North = g.Cell(row-1, col)
			g.grid[i][y].South = g.Cell(row+1, col)
			g.grid[i][y].West = g.Cell(row, col-1)
			g.grid[i][y].East = g.Cell(row, col+1)
		}
	}
}

func (g *Grid) RandomCell() *Cell {
	return g.Cell(rand.Intn(g.rows), rand.Intn(g.columns))
}

func (g *Grid) AllCells() []*Cell {
	cells := make([]*Cell, 0)
	for i := range g.grid {
		for y := range g.grid[i] {
			cells = append(cells, g.grid[i][y])
		}
	}

	return cells
}

func (g *Grid) PrintOut() string {
	output := "+" + strings.Repeat("---+", g.columns) + "\n"

	for r := range g.grid {

		top := "|"
		bottom := "+"

		for c := range g.grid[r] {
			var cell *Cell
			if g.grid[r][c] == nil {
				cell = NewCell(-1, -1)
			} else {
				cell = g.grid[r][c]
			}
			fmt.Println(cell.Row, cell.Column)
			body := "   " //<-- that's THREE(3)spaces!
			east_boundary := "|"
			if cell.IsLinked(cell.East) {
				east_boundary = " "
			}
			top += body + east_boundary

			south_boundary := "---"
			if cell.IsLinked(cell.South) {
				south_boundary = "   "
			}
			const corner = "+"
			bottom += south_boundary + corner
		}
		output += top + "\n"
		output += bottom + "\n"
	}

	return output

}

func (g *Grid) PrintOutCleaner(i Contents) string {
	output := "\u250C"
	//+ strings.Repeat("---+", g.columns) + "\n"

	//do top row
	for _, cell := range g.grid[0] {
		if cell.East == nil {
			output += "\u2500\u2500\u2500\u2510"
		} else {
			if cell.IsLinked(cell.East) {
				output += "\u2500\u2500\u2500\u2500"
			} else {
				output += "\u2500\u2500\u2500\u252C"
			}
		}
	}
	output += "\n"

	for r := range g.grid {

		top := "\u2502"

		cell := g.grid[r][0]
		var row string
		if cell.IsLinked(cell.South) {
			row = "\u2502" //│
		} else {
			if cell.South == nil {
				row = "\u2514" //└
			} else {
				row = "\u251C" //├
			}
		}

		for y := range g.grid[r] {
			if g.grid[r][y] == nil {
				g.grid[r][y] = NewCell(-1, -1)
			}
			cell = g.grid[r][y]

			var eastBoundary string
			if cell.IsLinked(cell.East) {
				eastBoundary = " "
			} else {
				eastBoundary = "\u2502" //│
			}
			body := i.Contents_of(cell)
			top += body + eastBoundary

			//// three spaces below, too >>-------------->> >...<
			var southBoundary string
			if cell.IsLinked(cell.South) {
				southBoundary = "   "
			} else {
				southBoundary = "\u2500\u2500\u2500"
			}

			var up, down, left, right bool = false, true, false, false

			if cell.South == nil {
				left = true
				down = false
			} else {
				left = !cell.IsLinked(cell.South)
			}

			if cell.East == nil {
				up = true
				right = false
			} else {
				up = !cell.IsLinked(cell.East)
				if cell.East.South == nil {
					right = true
					down = false
				} else {
					right = !cell.East.IsLinked(cell.East.South)
					down = !cell.South.IsLinked(cell.East.South)
				}
			}
			corner := " "
			if left && right && up && down {
				corner = "\u253C" //┼
			}

			if left && right && up && !down {
				corner = "\u2534" //┴
			}

			if left && right && !up && down {
				corner = "\u252C" //┬
			}

			if left && right && !up && !down {
				corner = "\u2500" //─
			}

			if !left && right && up && down {
				corner = "\u251C" //├
			}

			if !left && right && up && !down {
				corner = "\u2514" //└
			}

			if !left && right && !up && down {
				corner = "\u250C" //└
			}

			if left && !right && up && down {
				corner = "\u2524" //┤
			}

			if !left && !right && up && down {
				corner = "\u2502" //│
			}

			if left && !right && up && !down {
				corner = "\u2518" //┘
			}

			if left && right && !up && !down {
				corner = "\u2500"
			}

			if left && !right && !up && !down {
				corner = "\u2500"
			}

			if left && !right && !up && down {
				corner = "\u2510"
			}

			row += southBoundary + corner
		}
		output += top + "\n"
		output += row + "\n"

	}

	return output

}

func (g *Grid) DeadEnds() []*Cell {

	list := make([]*Cell, 0)
	for i := range g.grid {
		for y := range g.grid[i] {
			if len(g.grid[i][y].Links) == 1 {
				list = append(list, g.grid[i][y])
			}
		}
	}

	return list
}

func (g *Grid) BraidPartial(p float32) {

	slice := g.DeadEnds()
	for i := range slice {
		j := rand.Intn(i + 1)
		slice[i], slice[j] = slice[j], slice[i]
	}

	for _, cell := range slice {
		if len(cell.Links) != 1 || rand.Float32() > p {

		} else {
			neighbors := Filter(cell.Neighbours(), cell, NotLinked)
			best := make([]*Cell, 0)
			for _, v := range neighbors {
				if len(v.Links) == 1 {
					best = append(best, v)
				}
			}
			if len(best) == 0 {
				best = neighbors
			}
			cell.Link(best[rand.Intn(len(best))])
		}
	}
}

func (g *Grid) Braid() {
	g.BraidPartial(1)
}

// Filter returns a new slice holding only
// the elements of s that satisfy f()
func Filter(s []*Cell, c *Cell, fn func(*Cell, *Cell) bool) []*Cell {
	var p []*Cell // == nil
	for _, v := range s {
		if fn(c, v) {
			p = append(p, v)
		}
	}
	return p
}

func NotLinked(cell, neighbour *Cell) bool {
	return !cell.IsLinked(neighbour)
}

func (g *Grid) Contents_of(*Cell) string {
	return "   "
}

type Contents interface {
	Contents_of(*Cell) string
}

func (g Grid) String() string {
	return g.PrintOutCleaner(&g)
}

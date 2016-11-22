package maze

import (
	"math/rand"
)

func SideWinder(g *Grid) {
	for r := range g.grid {
		run := make([]*Cell, 0)

		for c := range g.grid[r] {
			cell := g.grid[r][c]
			run = append(run, cell)

			at_eastern_boundary := false
			if cell.East == nil {
				at_eastern_boundary = true
			}
			at_northern_boundary := false
			if cell.North == nil {
				at_northern_boundary = true
			}

			should_close_out :=
				at_eastern_boundary || (!at_northern_boundary && rand.Intn(2) == 0)

			if should_close_out {
				member := run[rand.Intn(len(run))]
				if member.North != nil {
					member.Link(member.North)
				}
				run = nil
			} else {
				cell.Link(cell.East)
			}

		}

	}
}

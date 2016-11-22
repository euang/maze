package maze

import (
	"math/rand"
)

func BinaryTree(grid *Grid) {

	for _, cell := range grid.AllCells() {
		neighbours := make([]*Cell, 0)
		if cell.North != nil {
			neighbours = append(neighbours, cell.North)
		}
		if cell.East != nil {
			neighbours = append(neighbours, cell.East)
		}

		if len(neighbours) > 0 {
			index := rand.Intn(len(neighbours))
			neighbor := neighbours[index]

			if neighbor != nil {
				cell.Link(neighbor)
			}
		}
	}
}

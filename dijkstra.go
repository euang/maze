package maze

import (
	"fmt"
)

//require 'binary_tree'
func Dijkstra() {
	grid := DistanceGrid{Grid: *NewGrid(11, 11)}
	BinaryTree(&grid.Grid)

	start := grid.Cell(0, 0)
	distances := start.Distances()

	grid.distances = *distances
	fmt.Println(grid)

	fmt.Println("path from northwest corner to southwest corner:")
	grid.distances = *distances.path_to(grid.Cell(grid.rows-1, 0))
	fmt.Println(grid)
}

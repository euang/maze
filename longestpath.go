package maze

import "fmt"

func LongestPath() {

	grid := DistanceGrid{Grid: *NewGrid(5, 5)}
	BinaryTree(&grid.Grid)

	start := grid.Cell(0, 0)
	distances := start.Distances()

	grid.distances = *distances
	new_start, _ := distances.Max()
	new_distances := new_start.Distances()
	fmt.Println(grid)

	goal, _ := new_distances.Max()

	grid.distances = *new_distances.path_to(goal)
	fmt.Println(grid)
	grid.to_png_v1(10)

}

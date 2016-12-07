package maze

import "fmt"

func Colouring() {
	grid := ColoredGrid{Grid: *NewGrid(25, 25)}
	BinaryTree(&grid.Grid)

	start := grid.Cell(grid.rows/2, grid.columns/2)
	distances := start.Distances()

	grid.SetDistances(*distances)

	fmt.Println(grid)
	grid.toPngV1(10)
	grid.toPngV2(10, &grid)

}

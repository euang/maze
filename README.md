# maze

A Go library for generating and visualising mazes.

## Features

- **Grid**: Rectangular grid of cells with north/south/east/west neighbour links
- **Maze generation algorithms**:
  - **Binary Tree** – simple, fast algorithm with a north-east bias
  - **Sidewinder** – row-by-row algorithm with a northern bias
- **Pathfinding**: Dijkstra's algorithm for finding shortest paths and distances between cells
- **Braiding**: Remove dead ends to create mazes with loops (`Braid` / `BraidPartial`)
- **Rendering**:
  - Text output using Unicode box-drawing characters
  - PNG export with optional colour-coded distance heatmaps

## Usage

### Generate a maze

```go
import "github.com/euang/maze"

// Create a 10x10 grid and apply the Binary Tree algorithm
g := maze.NewGrid(10, 10)
maze.BinaryTree(g)
fmt.Println(g)

// Or use the Sidewinder algorithm
g2 := maze.NewGrid(10, 10)
maze.SideWinder(g2)
fmt.Println(g2)
```

### Find distances / shortest path

```go
// Solve the maze from the top-left cell to the bottom-left cell
maze.Dijkstra()

// Find the longest path in the maze
maze.LongestPath()
```

### Colour-coded PNG output

```go
// Render a 25x25 maze with a distance heatmap
maze.Colouring()
```

This generates `hello.png` and `hello2.png` in the working directory.

### Removing dead ends (braiding)

```go
g := maze.NewGrid(10, 10)
maze.BinaryTree(g)

// Remove all dead ends
g.Braid()

// Remove ~50% of dead ends
g.BraidPartial(0.5)
```

## Package structure

| File | Description |
|---|---|
| `cell.go` | `Cell` type and link/distance logic |
| `grid.go` | `Grid` type, text/PNG rendering, braiding |
| `binarytree.go` | Binary Tree maze generation algorithm |
| `sidewinder.go` | Sidewinder maze generation algorithm |
| `distances.go` | Dijkstra distance map and path tracing |
| `distancesgrid.go` | Grid that renders distance values in cells |
| `coloredgrid.go` | Grid that renders a colour heatmap by distance |
| `dijkstra.go` | Example: solve a maze with Dijkstra's algorithm |
| `longestpath.go` | Example: find the longest path in a maze |
| `colouring.go` | Example: render a colour-coded PNG maze |

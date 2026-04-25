# Copilot Instructions for `euang/maze`

## Project Overview

This is a **Go library** (`package maze`) for generating and visualising mazes. It lives entirely in the root directory—there is no `cmd/` entry point, no `main` package, and no `go.mod` file. All source files declare `package maze`.

The library was developed against an early Go toolchain (the IDE config references Go 1.7) and has no module system. There are no tests.

---

## Repository Layout

```
/                        ← all Go source files (package maze)
├── cell.go              ← Cell type: row/col coordinates, neighbour pointers, link map
├── grid.go              ← Grid type: 2-D slice of *Cell, text/PNG rendering, braiding
├── distances.go         ← Distances type: BFS distance map + path-tracing
├── distancesgrid.go     ← DistanceGrid embeds Grid, renders cell distances as base-36 text
├── coloredgrid.go       ← ColoredGrid embeds Grid, supplies colour-heatmap background
├── binarytree.go        ← BinaryTree() maze-generation algorithm (north/east bias)
├── sidewinder.go        ← SideWinder() maze-generation algorithm (row-by-row, northern bias)
├── dijkstra.go          ← Example: Dijkstra() – solves an 11×11 maze, prints paths
├── longestpath.go       ← Example: LongestPath() – finds the longest path in a 5×5 maze
├── colouring.go         ← Example: Colouring() – renders a 25×25 colour-coded PNG maze
└── .github/
    └── copilot-instructions.md   ← this file
```

---

## Key Types and Their Relationships

### `Cell` (`cell.go`)
- Fields: `Row`, `Column int`; `North`, `South`, `East`, `West *Cell`; `Links map[*Cell]*Cell`
- `Link(cell)` – bidirectional link; `LinkOneWay(cell)` – one-way only; `UnLink(cell)` removes entry
- `IsLinked(cell) bool` – checks the Links map
- `Neighbours() []*Cell` – returns non-nil cardinal neighbours
- `Distances() *Distances` – runs BFS from `self` to build a full distance map

### `Grid` (`grid.go`)
- Fields: `rows`, `columns int`; `grid [][]*Cell`
- `NewGrid(rows, cols)` – allocates and wires up all Cell neighbours
- `Cell(row, col)` – safe accessor (returns nil for out-of-bounds)
- `AllCells() []*Cell` – flat slice of every cell
- `RandomCell() *Cell`
- `DeadEnds() []*Cell` – cells with exactly one link
- `Braid()` / `BraidPartial(p float32)` – remove dead ends to create loops
- `PrintOut()` – ASCII art using `+`, `|`, `-` (has debug `fmt.Println` calls)
- `PrintOutCleaner(i Contents)` – Unicode box-drawing art; delegates cell contents via `Contents` interface
- `String()` – calls `PrintOutCleaner(&g)` so `fmt.Println(grid)` works
- `toPngV1(cell_size int)` – writes `hello.png` (walls only)
- `toPngV2(cell_size int, bg BackgroundColor)` – writes `hello2.png` (coloured cells + walls)

### Interfaces defined in `grid.go`
- `Contents` – `Contents_of(*Cell) string` (3-char cell body for text rendering)
- `BackgroundColor` – `BackgroundColorFor(*Cell) color.RGBA`
- `Grid` implements both interfaces with default implementations

### `Distances` (`distances.go`)
- `Cells map[*Cell]int` – public; maps cell pointer to BFS distance from root
- `root *Cell` – private starting cell
- `path_to(goal *Cell) *Distances` – traces shortest path back to root
- `Max() (*Cell, int)` – returns the farthest cell and its distance
- `AllCells() []*Cell` – all keys

### `DistanceGrid` (`distancesgrid.go`)
- Embeds `Grid`; has `distances Distances` field
- Overrides `Contents_of` to render distances in base-36 (3-char wide)
- Overrides `String()` so `fmt.Println` uses distance-annotated rendering

### `ColoredGrid` (`coloredgrid.go`)
- Embeds `Grid`; has `distances Distances` and `maximum int` fields
- `SetDistances(distances Distances)` – stores distances and computes maximum
- Overrides `BackgroundColorFor` to return a green heatmap colour

---

## Build System

**There is no `go.mod`** – the project predates Go modules. To work with it:

```sh
# Set GOPATH to a parent directory and place the repo under src/github.com/euang/maze
export GOPATH=/some/path
mkdir -p $GOPATH/src/github.com/euang
ln -s /path/to/repo $GOPATH/src/github.com/euang/maze
cd $GOPATH/src/github.com/euang/maze
go build github.com/euang/maze
```

Alternatively, initialise a module to modernise the project:

```sh
go mod init github.com/euang/maze
go mod tidy   # fetches github.com/llgcode/draw2d
```

### External dependency
The only external dependency is `github.com/llgcode/draw2d` (used in `grid.go` for PNG rendering via `draw2dimg` and `draw2dkit`).

### No tests exist
There are no `*_test.go` files. When adding new functionality, create `_test.go` files in `package maze` (same package, no `_test` suffix needed for white-box access).

---

## Known Issues / Quirks to Be Aware Of

1. **Debug `fmt.Println` left in production code** – `grid.go` has several `fmt.Println("cell", ...)` and `fmt.Println(x1, y1, x2, y2)` calls in `PrintOut()` and `toPngV1/V2`. These produce noisy output but are not bugs per se; remove them only if explicitly asked.

2. **`PrintOut()` vs `PrintOutCleaner()`** – `PrintOut()` uses ASCII art and calls the debug prints above. `PrintOutCleaner()` uses Unicode box-drawing and is what `String()` delegates to. Prefer `PrintOutCleaner`.

3. **`dijkstra.go`, `longestpath.go`, `colouring.go` are example drivers, not algorithms** – they contain package-level functions (`Dijkstra()`, `LongestPath()`, `Colouring()`) that hard-code grid sizes and write files. They are not idiomatic library code; treat them as demo/example code.

4. **PNG output is side-effectful** – `toPngV1` always writes `hello.png`; `toPngV2` always writes `hello2.png` to the current working directory. There is no way to customise the output path without modifying the source.

5. **`UnLink` is one-directional** – unlike `Link`, `UnLink` only removes the entry from one cell's map. If bidirectional unlinking is needed, call `UnLink` on both cells.

6. **`BraidPartial` uses a shuffle then skips even-odds** – the inner `if` has an empty branch (`} else {` … `}`), which is intentional flow-control but looks like a bug; don't remove the empty branch without understanding the logic.

7. **No `go.mod`** – running `go build ./...` directly in the repo root will fail with `pattern ./...: directory prefix . does not contain main module`. Add a `go.mod` first.

---

## Development Workflow

1. **Adding a new maze algorithm**: Create a new `.go` file in `package maze`. Accept `*Grid` as the parameter. Call `cell.Link(neighbour)` to carve passages.

2. **Adding a new renderer**: Implement the `Contents` interface (for text) or the `BackgroundColor` interface (for PNG). Embed `Grid` and override the relevant method, then override `String()` to call `g.Grid.PrintOutCleaner(&yourType)`.

3. **Running example functions**: Since there is no `main` package, call `maze.Dijkstra()`, `maze.LongestPath()`, or `maze.Colouring()` from an external `main.go`.

4. **Checking dead ends and braiding**: `g.DeadEnds()` returns dead-end cells; `g.Braid()` removes all of them; `g.BraidPartial(0.5)` removes roughly half at random.

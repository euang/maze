package maze

import (
	"fmt"
	"image/color"
	"math"
)

type ColoredGrid struct {
	Grid
	distances Distances
	maximum   int
}

func (g *ColoredGrid) SetDistances(distances Distances) {

	g.distances = distances
	_, g.maximum = distances.Max()
}

func (g *ColoredGrid) BackgroundColorFor(cell *Cell) color.RGBA {

	distance := g.distances.Cells[cell]
	fmt.Println("dist:", distance)
	intensity := float64(g.maximum-distance) / float64(g.maximum)
	dark := uint8(math.Ceil(255.0 * intensity))
	bright := 128 + uint8(math.Ceil(127.0*intensity))

	return color.RGBA{dark, bright, dark, 0xff}

}

package maze

import (
	"fmt"
	"strconv"
)

type DistanceGrid struct {
	Grid
	distances Distances
}

func (d *DistanceGrid) Contents_of(cell *Cell) string {

	if _, ok := d.distances.Cells[cell]; ok {
		return fmt.Sprintf("%3s", strconv.FormatInt(int64(d.distances.Cells[cell]), 36))
	} else {
		return d.Grid.Contents_of(cell)
	}
}

func (d DistanceGrid) String() string {
	return d.Grid.PrintOutCleaner(&d)
}

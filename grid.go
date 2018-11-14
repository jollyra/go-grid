package grid

import (
	"fmt"
	"strings"
)

// Grid represents a 2D grid of size dx x dy.
type Grid struct {
	dx   int
	dy   int
	Grid [][]byte
}

// NewGrid initializes a new grid of size dx x dy.
func NewGrid(dx, dy int) *Grid {
	grid := make([][]byte, dy)
	for y := range grid {
		grid[y] = make([]byte, dx)
	}

	return &Grid{
		dx:   dx,
		dy:   dy,
		Grid: grid,
	}
}

// Copy returns a deep copy of Grid g.
func (g *Grid) Copy() *Grid {
	copy := NewGrid(g.dx, g.dy)
	for y := range g.Grid {
		for x := range g.Grid[y] {
			copy.Grid[y][x] = g.Grid[y][x]
		}
	}
	return copy
}

// IsInBounds returns true if point (x, y) is in the bounds of Grid g,
// false otherwise.
func (g *Grid) IsInBounds(x, y int) bool {
	if 0 <= x && x < g.dx && 0 <= y && y < g.dy {
		return true
	}
	return false
}

// Neighbours8 returns all neighbour of point (x, y) including diagonals.
func (g *Grid) Neighbours8(x, y int) []Point {
	points := []Point{
		Point{x + 1, y},
		Point{x + 1, y + 1},
		Point{x, y + 1},
		Point{x - 1, y + 1},
		Point{x - 1, y},
		Point{x - 1, y - 1},
		Point{x, y - 1},
		Point{x + 1, y - 1},
	}

	pointsInBounds := make([]Point, 0)
	for _, p := range points {
		if g.IsInBounds(p.X, p.Y) {
			pointsInBounds = append(pointsInBounds, p)
		}
	}

	return pointsInBounds
}

// String returns a string representation of Grid g.
func (g *Grid) String() string {
	var b strings.Builder
	for _, row := range g.Grid {
		for _, x := range row {
			fmt.Fprintf(&b, "%c", x)
		}
		fmt.Fprintf(&b, "\n")
	}
	return b.String()
}

package grid

// Point represents a simple 2D point.
type Point struct{ X, Y int }

// Equals returns true if p and q are the same.
func (p *Point) Equals(q Point) bool {
	if p.X == q.X && p.Y == q.Y {
		return true
	}
	return false
}

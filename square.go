package square

import "math"

type Point struct {
	x, y int
}

type Square struct {
	start Point
	a     uint
}

func (s Square) End() Point {
	return Point{s.start.x + int(s.a), s.start.y + int(s.a)}
}

func (s Square) Area(x2, y2 int) uint {
	sq := math.Abs(float64((x2 - s.start.x) * (y2 - s.start.y)))
	return uint(sq)
}

func (s Square) Perimeter(x2, y2 int) uint {
	p := 2*math.Abs(float64(x2-s.start.x)) + 2*math.Abs(float64(y2-s.start.y))
	return uint(p)
}

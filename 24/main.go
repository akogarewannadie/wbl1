package main

import (
	"fmt"
	"math"
)

type Point struct {
	x, y float64
}

func NewPoint(x, y float64) Point {
	return Point{x, y}
}

func (p Point) DistanceTo(other Point) float64 {
	dx := p.x - other.x
	dy := p.y - other.y
	return math.Sqrt(dx*dx + dy*dy)
}

func main() {

	p1 := NewPoint(0, 0)
	p2 := NewPoint(3, 4)

	distance := p1.DistanceTo(p2)

	fmt.Printf("Расстояние между точками (%.1f, %.1f) и (%.1f, %.1f) равно %.2f\n", p1.x, p1.y, p2.x, p2.y, distance)
}

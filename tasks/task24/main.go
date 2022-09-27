package main

import (
	"fmt"
	"math"
)

func main() {
	p1 := newPoint(0.0, 3.0)
	p2 := newPoint(4.0, 0.0)
	fmt.Printf("%.2f", findDistance(p1, p2))
}

type Point struct {
	x float64
	y float64
}

func newPoint(x, y float64) *Point {
	return &Point{x: x, y: y}
}

func findDistance(p1 *Point, p2 *Point) float64 {
	//по теореме Пифагора находим расстояние между точками
	squareX := (p1.x - p2.x) * (p1.x - p2.x)
	squareY := (p1.y - p2.y) * (p1.y - p2.y)

	return math.Sqrt(squareX + squareY)
}

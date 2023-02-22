package main

import (
	"fmt"
	"test/reccount"
)

type Point = reccount.Point

func main() {
	points := []Point{{X: 1, Y: 2}, {X: 3, Y: 2}, {X: 3, Y: 4}, {X: 1, Y: 4}, {X: 3, Y: 5}, {X: 1, Y: 5}}
	count := reccount.HowManyRectangles(points)
	fmt.Printf("There are %v rectangles", count)
}

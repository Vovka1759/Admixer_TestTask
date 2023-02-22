package reccount

import (
	//"fmt"
	"math"
)

type Point struct {
	X float64
	Y float64
}

func hasNoDuplicats(points []Point) bool {
	for i := 0; i < len(points); i++ {
		for j := 0; j < len(points); j++ {
			if i != j && points[i] == points[j] {
				return false
			}
		}
	}
	return true
}

func isItRectangle(points [4]Point) bool {
	if !hasNoDuplicats(points[:]) {
		return false
	}

	centerX := (points[0].X + points[1].X + points[2].X + points[3].X) / 4
	centerY := (points[0].Y + points[1].Y + points[2].Y + points[3].Y) / 4
	distance1 := math.Sqrt(math.Abs(centerX-points[0].X)) + math.Sqrt(math.Abs(centerY-points[0].Y))
	distance2 := math.Sqrt(math.Abs(centerX-points[1].X)) + math.Sqrt(math.Abs(centerY-points[1].Y))
	distance3 := math.Sqrt(math.Abs(centerX-points[2].X)) + math.Sqrt(math.Abs(centerY-points[2].Y))
	distance4 := math.Sqrt(math.Abs(centerX-points[3].X)) + math.Sqrt(math.Abs(centerY-points[3].Y))

	return distance1 == distance2 && distance1 == distance3 && distance1 == distance4
}

func HowManyRectangles(points []Point) int {
	counter := 0
	for a := 0; a < len(points); a++ {
		for b := a + 1; b < len(points); b++ {
			for c := b + 1; c < len(points); c++ {
				for d := c + 1; d < len(points); d++ {
					if isItRectangle([4]Point{points[a], points[b], points[c], points[d]}) {
						//fmt.Println([4]Point{points[a], points[b], points[c], points[d]},"is a rectangle")
						counter++
					}
				}
			}
		}
	}
	return counter
}

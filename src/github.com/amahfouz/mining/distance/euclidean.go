package distance

import "math"
import myMath "github.com/amahfouz/util/math"

// functions to compute Euclidean distance

func ComputeL1Dist(a, b myMath.Point) float64 {
	return math.Abs(float64(a.X-b.X)) + math.Abs(float64(a.Y-b.Y))
}

func ComputeL2Dist(a, b myMath.Point) float64 {
	xDiff := a.X - b.X
	yDiff := a.Y - b.Y

	return math.Sqrt(float64(xDiff*xDiff + yDiff*yDiff))
}

func NearestDistance(reference []myMath.Point,
	points []myMath.Point,
	dist func(myMath.Point, myMath.Point) float64) []myMath.Point {
	result := make([]myMath.Point, len(points))

	for j, p := range points {
		minDist := math.Inf(1)
		for i, ref := range reference {
			curDist := dist(p, ref)
			if curDist < minDist {
				minDist = curDist
				result[j] = reference[i]
			}
		}
	}
	return result
}

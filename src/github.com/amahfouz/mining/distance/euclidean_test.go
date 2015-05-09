package distance

import (
	"strconv"
	"testing"
)

import myMath "github.com/amahfouz/util/math"

func TestL1Dist(t *testing.T) {
	p1 := myMath.Point{1, 1}
	p2 := myMath.Point{5, 4}

	dist := ComputeL1Dist(p1, p2)
	if dist != 7.0 {
		t.Error("Expected 7 but got a distance " + strconv.FormatFloat(dist, 'f', 2, 64))
	}
}

func TestL2Dist(t *testing.T) {
	p1 := myMath.Point{1, 1}
	p2 := myMath.Point{5, 4}

	dist := ComputeL2Dist(p1, p2)
	if dist != 5.0 {
		t.Error("Expected 5 but got a distance " + strconv.FormatFloat(dist, 'f', 2, 64))
	}
}

package main

import "fmt"

import "github.com/amahfouz/util/math"
import "github.com/amahfouz/mining/distance"

func main() {
	// Week4A (Basic) - Question #2
	
	factors := []float64 {0, 0.5, 1, 2}
	
	for _,factor := range factors {
		a := math.NewIntVector([]int64 {1, 0, 1, 0, 1, 2})
		b := math.NewIntVector([]int64 {1, 1, 0, 0,	1, 6})
		c := math.NewIntVector([]int64 {0, 1, 0, 1,	0, 2})

		a.Items[5] = int64(float64(a.Items[5]) * factor)
		b.Items[5] = int64(float64(b.Items[5]) * factor)
		c.Items[5] = int64(float64(c.Items[5]) * factor)
		
		fmt.Printf("----- Factor = %f\n", factor)
		fmt.Printf("Dist(a, b) = %f\n", distance.CosineDistance(*a, *b))
		fmt.Printf("Dist(a, c) = %f\n", distance.CosineDistance(*a, *c))
		fmt.Printf("Dist(c, b) = %f\n", distance.CosineDistance(*c, *b))
	}
}

func old_main() {
	refs := []math.Point{math.Point{0, 0}, math.Point{100, 40}}
	pts := []math.Point{math.Point{61, 10}, math.Point{59, 10}, math.Point{52, 13}, math.Point{54, 8}}

	closest := distance.NearestDistance(refs, pts, distance.ComputeL1Dist)

	for _, p := range closest {
		fmt.Printf("%d , %d\n", p.X, p.Y)
	}

	closest = distance.NearestDistance(refs, pts, distance.ComputeL2Dist)

	for _, p := range closest {
		fmt.Printf("%d , %d\n", p.X, p.Y)
	}

}

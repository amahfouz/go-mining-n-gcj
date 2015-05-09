package distance

import myMath "github.com/amahfouz/util/math"

func CosineSimilarity(a, b myMath.IntVector) float64 {
	return float64(a.Dot(b)) / (a.Magnitude() * b.Magnitude()) 	
}

func CosineDistance(a, b myMath.IntVector) float64 {
    return 1 - CosineSimilarity(a, b)
}

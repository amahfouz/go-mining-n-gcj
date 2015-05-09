package distance
import "testing"

import "math"
import myMath "github.com/amahfouz/util/math"

func TestOrthogonal(t *testing.T) {
    a := myMath.NewIntVector([]int64 {3, 0})
    b := myMath.NewIntVector([]int64 {0, 4})
    
    if CosineSimilarity(*a, *b) != 0.0 {
        t.Error("TestOrthogonal failed.");
    }
}

func TestCosineSim(t *testing.T) {
    a := myMath.NewIntVector([]int64 {3, 4})
    b := myMath.NewIntVector([]int64 {4, 3})
    
    if math.Abs(CosineSimilarity(*a, *b) - 0.96) > 0.01 {
        t.Error("TestCosineDist failed.");
    }
}
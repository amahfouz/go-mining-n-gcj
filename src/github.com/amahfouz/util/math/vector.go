package math

import "math"

type IntVector struct {
    Items[] int64
}

func NewIntVector(data []int64) *IntVector {
    v := new(IntVector)
    
    if data == nil {
        data = make([]int64, 0)
    }
    v.Items = data
    
    return v
}

func (vector IntVector) Size() int {
    return len(vector.Items)
}

func (vector IntVector) Magnitude() float64 {
    var sumSquares int64 = 0
    
    for i := range vector.Items {
		sumSquares += vector.Items[i] * vector.Items[i]         
    }
    return math.Sqrt(float64(sumSquares))
}

func (v IntVector) Dot(u IntVector) int64 {
    
    if v.Size() != u.Size() {
        panic("Vectors have to be of same size for dot product.")
    }
    
    var sum int64 = 0
    
    for i := range v.Items {
        sum += v.Items[i] * u.Items[i]
    }
    return sum
}  
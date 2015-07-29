package algo

import (
	"testing"
)

var (
	size           = 500
	matrix [][]int = createRandomMatrix(size)
	natrix [][]int = createRandomMatrix(size)
)

func zTestCorrectness(t *testing.T) {
	result1 := serialSquareMatrixMul(matrix, natrix)
	result2 := ParalellSquareMatrixMul(matrix, natrix)

	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			if result1[i][j] != result2[i][j] {
				t.Error("Not equal!")
			}
		}
	}
}

func TestSerial(t *testing.T) {
	serialSquareMatrixMul(matrix, natrix)
}

func TestParallel(t *testing.T) {
	ParalellSquareMatrixMul(matrix, natrix)
}

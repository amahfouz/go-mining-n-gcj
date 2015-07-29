package algo

import (
	"math/rand"
	"time"
)

func ParalellSquareMatrixMul(a, b [][]int) [][]int {
	size := len(a)
	result := createMatrix(size)

	semaphore := make(chan bool)
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			go mulVectors(a, b, result, i, j, semaphore)
		}
	}

	for wait := 0; wait < size*size; wait++ {
		<-semaphore
	}
	return result
}

func mulVectors(a, b, result [][]int, i, j int, c chan bool) {
	size := len(a)
	for k := 0; k < size; k++ {
		result[i][j] += a[i][k] * a[k][j]
	}
	c <- true
}

// assumes square matrix - no error checking
func serialSquareMatrixMul(a, b [][]int) [][]int {
	size := len(a)
	result := createMatrix(size)
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			for k := 0; k < size; k++ {
				result[i][j] += a[i][k] * a[k][j]
			}
		}
	}
	return result
}

// utilities

func createMatrix(dim int) [][]int {
	matrix := make([][]int, dim)
	for i := range matrix {
		matrix[i] = make([]int, dim)
	}
	return matrix
}

func createRandomMatrix(dim int) [][]int {
	matrix := make([][]int, dim)

	randomizer := createRandomizer()

	for i := range matrix {
		matrix[i] = make([]int, dim)

		for j := range matrix[i] {
			matrix[i][j] = randomizer.Int()
		}
	}
	return matrix
}

func createRandomizer() *rand.Rand {
	return rand.New(rand.NewSource(time.Now().Unix()))
}

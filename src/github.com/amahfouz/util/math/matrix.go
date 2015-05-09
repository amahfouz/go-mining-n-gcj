package math

import "fmt"

type IntMatrix struct  {
    M[][] int64
    Rows int
    Columns int
}

func NewIntMatrix(data [][]int64) *IntMatrix {
    m := new(IntMatrix)
    m.M = data
    m.Rows = len(data)
    if m.Rows > 0 {
    	m.Columns = len(m.M[0])    
    } else {
     	m.Columns = 0
    }
    return m
}

func (m IntMatrix) String() string {
    return fmt.Sprintln(m.M)
}
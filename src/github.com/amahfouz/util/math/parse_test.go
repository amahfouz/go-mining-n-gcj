package math

import (	
    "testing"
    "bufio"
    "os"
)

func TestParseMatrix(t *testing.T) {
    var file, err = os.Open(os.Getenv("GOPATH") + "/resource/test/util/math/6by7.matrix")
    if err != nil {
        t.Error(err)
    }
    defer file.Close()
    
 	matrix := ParseMatrix(bufio.NewReader(file))
    
    if matrix.Rows != 6 {
        t.Error("Wrong nnumber of rows " + string(matrix.Rows))
    }
    if matrix.Columns != 7 {
        t.Error("Wrong nnumber of columns " + string(matrix.Columns))
    }
    if matrix.M[2][2] != 2 {
        t.Error("Wrong value in matrix")
    }
}
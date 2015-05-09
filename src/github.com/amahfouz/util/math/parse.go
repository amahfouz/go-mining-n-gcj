// parser for input to be interpreted as matrices and vectors

package math

import ( 
	"fmt"
	"strings"
	"strconv"
	"io"
	"bufio"
)

// parses a string made of comma-separated list of ints
func parseCsInts(s string) []int64 {
    splits := strings.Split(s, ",")

    var result []int64 = make([]int64, len(splits))
    for i := range result {
        parsed, err := strconv.ParseInt(strings.TrimSpace(splits[i]), 10, 64)
        
        if err == nil {
        	result[i] = parsed  
        } else {
            panic(fmt.Sprintf("Failed to parse '%s' as an integer!" ))
        } 
    }
	return result
}

// parses a matrix represented as lines of comma-separated ints 
func ParseMatrix(r io.Reader) IntMatrix {
    
    scanner := bufio.NewScanner(r)
    result := make([][]int64, 0, 10)
    for scanner.Scan() {
	    result = append(result, parseCsInts(scanner.Text()))        
    }
    
    return *NewIntMatrix(result)
}

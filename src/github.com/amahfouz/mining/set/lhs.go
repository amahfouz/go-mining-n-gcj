package set

import "github.com/amahfouz/util/math"

type stringErr struct {
    message string
}

func (e stringErr) Error() string {
    return e.message
}

func ComputeLhs() int {
    
    //sig := make([][]int, 6)
    
//    sig[0] := make(int, 7)
//    sig[0][0] = 1
    
    //sig[0] = [...]int{1, 2, 1,	1, 2, 5, 4}
    
    return 0 
}

// computes candidate pairs for similarity given the signature matrix
// this method assumes that the buckets are infinite, which translates
// to two pairs being candidate iff they have the exact same values for
// any of the bands
func ComputeLhsCandidatePairs(signature math.IntMatrix, bands int) (math.PointSet, error) {
    if bands <= 0 {
        return math.PointSet{}, stringErr{"Band size has to be a positive integer."}
    }
    if signature.Rows % bands != 0 {
        return math.PointSet{}, stringErr{"Number of signature rows has to be divisible by band size."}
    }
    
    // for each band compare each row for each pair of columns
    
	candidates := new(math.PointSet)
    
	rowsPerBand := signature.Rows / bands
 
    for band := 0; band < bands; band++ {
		firstRowInBand := band * rowsPerBand
        lastRowInBand := firstRowInBand + rowsPerBand - 1

        for i := 0; i < signature.Columns - 1; i++ {
            for j := i + 1; j < signature.Columns; j++ {
                columnsDoMatch := true
                for r := firstRowInBand; r <= lastRowInBand; r++ {
                    if signature.M[r][i] != signature.M[r][j] {
                        columnsDoMatch = false
                        break; // row different; columns are not a candidate
                    }
                }
                if columnsDoMatch {
                    candidates.Add(math.Point{float64(i+1), float64(j+1)}) // col numbers in problem are 1-based
                }
            } 
        }
    }

    return *candidates, nil
}
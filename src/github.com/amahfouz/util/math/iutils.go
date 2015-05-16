package math

// Integer math utils


func Abs(x uint8) uint8 {
    if x > 0 {
        return x
    }
    return -x
}
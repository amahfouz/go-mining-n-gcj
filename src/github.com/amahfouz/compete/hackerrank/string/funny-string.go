package main

import "fmt"
import "strconv"
import "bufio"
import "os"

func main() {
//    f, err := os.Open("C:\\Users\\amahfouz\\Desktop\\input.txt")
//    
//    if err != nil {
//    	panic(err)
//    }

    r := bufio.NewReader(os.Stdin)
    s := bufio.NewScanner(r)
    s.Split(bufio.ScanWords)
    
    s.Scan()   // scan num cases
    
    numCases, _ := strconv.Atoi(s.Text())
    
    for i := 0; i < numCases; i++ {
		s.Scan()
		solveCase(s.Text())
	}       
}

func solveCase(str string) {
    runes := []rune(str)
    var length = len(runes)
    for j := 0; j < length / 2; j++ {
        if Abs(runes[j+1] - runes[j]) != Abs(runes[length - j - 1] - runes[length - j - 2]) {
            fmt.Println("Not Funny")
            return
        }
    } 
    
    fmt.Println("Funny")
}

func Abs(x rune) rune {
    if x > 0 {
        return x
    }
    return -x
}
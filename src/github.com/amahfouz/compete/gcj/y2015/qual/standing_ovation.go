package main

import "fmt"
import "strings"
import "bufio"
import "os"
import "strconv"

func main() {
    f, err := os.Open("C:\\Users\\amahfouz\\Desktop\\input.txt")
    
    if err != nil {
    	panic(err)
    }
    
    r := bufio.NewReader(f)
    s := bufio.NewScanner(r)
    s.Split(bufio.ScanWords)
    
    s.Scan();   // scan num cases
    
    numCases, err := strconv.Atoi(s.Text())
    
    var numClasses int
    
    for i := 0; i < numCases; i++ {
        s.Scan()
        numClasses, err = strconv.Atoi(s.Text())
        s.Scan()
		classes := s.Text()
		
		arr := make([] int, len(classes))
		splits := strings.SplitN(classes, "", numClasses + 1)
		
		for j := range arr {
		    arr[j], err = strconv.Atoi(splits[j])
		}          
		
		solveCase(i, arr)
    }

}

func solveCase(caseNum int, arr [] int) {
    standing := 0
    invites := 0
    for i, num := range arr {
        if num == 0 {
            continue
        }
        
        if standing < i {
            invites += i - standing
            standing += invites
        }
        standing += num
    } 
    fmt.Printf("Case #%d: %d\n", caseNum+1, invites)
}

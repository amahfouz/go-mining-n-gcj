package main

import (
	"fmt"
	"github.com/amahfouz/mining/set"
)

func main() {
	s1 := "ABRACADABRA"
	s2 := "BRICABRAC"

	fmt.Println(set.Jaccard(s1, s2, 2))
}

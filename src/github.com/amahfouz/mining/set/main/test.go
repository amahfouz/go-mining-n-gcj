package main

import (
	"github.com/amahfouz/mining/set"
	"github.com/amahfouz/util/math"
)
import (
	"bufio"
	"fmt"
	"os"
)

func x() {
	var file, err = os.Open(os.Getenv("GOPATH") + "/resource/test/util/math/6by7.matrix")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	fmt.Println("dddddd")
	signatures := math.ParseMatrix(bufio.NewReader(file))
	fmt.Println(set.ComputeLhsCandidatePairs(signatures, 3))
}

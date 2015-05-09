package main

import (
    "fmt"
    "strconv"
)
import (
    "github.com/amahfouz/mining/cluster"
    m "github.com/amahfouz/util/math"
)

// Week 5B Q2

//    points := []m.Point { m.NewPoint(5,10),
//        				  m.NewPoint(20,5),
//        				  
//        				  m.NewPoint(7,12),
//        				  m.NewPoint(7,8),
//        				  m.NewPoint(12,8),
//        				  m.NewPoint(12,12),
//        				  m.NewPoint(16,19),
//        				  m.NewPoint(16,12),
//        				  m.NewPoint(25,12),
//        				  m.NewPoint(25,19) }
//    
//    pointSet := m.NewPointSetWithPoints(points)
//    	
//    cs := cluster.RunKMeans(2, pointSet, 20)


// final exam #18

func main() {
    points := []m.Point { m.NewPoint(1,1),
        				  m.NewPoint(4,4),
        				  
        				  m.NewPoint(2,1),
        				  m.NewPoint(2,2),
        				  m.NewPoint(3,3),
        				  m.NewPoint(4,2),
        				  m.NewPoint(2,4) }
    
    pointSet := m.NewPointSetWithPoints(points)
    	
    cs := cluster.RunKMeans(2, pointSet, 20)

	for i,c := range cs.Clusters {
	    fmt.Println(strconv.Itoa(i) + ") " + c.String())
	}    
    
}

// Week 5B Q1

func mainQ1() {
    
    // first 10 points are initial centroids
    points := []m.Point { m.NewPoint(25,125), 
        				  m.NewPoint(44,105), 
        				  m.NewPoint(29,97), 
        				  m.NewPoint(35,63), 
        				  m.NewPoint(55,63), 
        				  m.NewPoint(42,57), 
        				  m.NewPoint(23,40), 
        				  m.NewPoint(64,37), 
        				  m.NewPoint(33,22), 
        				  m.NewPoint(55,20), 
        				  /* end of centroids */
        				  m.NewPoint(28,145),
        				  m.NewPoint(65,140),
        				  m.NewPoint(50,130),
        				  m.NewPoint(38,115),
        				  m.NewPoint(55,118),
        				  m.NewPoint(50,90),
        				  m.NewPoint(63,88),
        				  m.NewPoint(43,83),
        				  m.NewPoint(50,60),
        				  m.NewPoint(50,30) } 

    pointSet := m.NewPointSetWithPoints(points)
    	
    cs := cluster.RunKMeans(10, pointSet, 1)

	for i,c := range cs.Clusters {
	    fmt.Println(strconv.Itoa(i) + ") " + c.String())
	}    
}
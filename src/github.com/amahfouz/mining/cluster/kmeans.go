package cluster

import "math"
import myMath "github.com/amahfouz/util/math"
import "github.com/amahfouz/mining/distance"

// k-means clustering 

type ClusterSet struct {
    
    Clusters []Cluster
    // centroid for each cluster (parallel to array of clusters)
    centroids []myMath.Point
    // assignment of points to cluster (cluster index is one-based)
    assignment map[myMath.Point] int
}

func NewClusterSet(initCentroids myMath.PointSet, k int) ClusterSet {
    cs := ClusterSet{ make([]Cluster, k), 
        			  make([]myMath.Point, k),
        			  make(map[myMath.Point]int) }
    
    for i := range cs.Clusters {
		cs.Clusters[i] = NewCluster()
		cs.centroids[i] = initCentroids.At(i)
		cs.assign(cs.centroids[i], i)
    }
    return cs
}

// first k points in the set are chosen as the initial centroids for the 'k' clusters

func RunKMeans(k int, points myMath.PointSet, maxIters int) ClusterSet {
    
    if points.NumElements() < k {
        panic("Too few points!")
    }
    
    // init clusters
    
    set := NewClusterSet(points, k)
    
    // classify points
    
    for i := 0; i < maxIters; i++ {
		runOneIteration(points, &set)
    }
    
    return set
}

/////////////////////////////////////////////////////////////////////////////
// private methods

func runOneIteration(points myMath.PointSet, set *ClusterSet) {
	iter := points.Iterator()
    valid := true
    var p myMath.Point
        
    // classify each point
            
    p, valid = iter()
    for valid {
    	set.classify(p)
    	p, valid = iter()
	}
    
    // recompute centroid of every cluster
    
    for i,_ := range set.Clusters {
        set.centroids[i] = set.Clusters[i].Centroid()
    }
}

func (cs ClusterSet) classify(point myMath.Point) {
    minDist := math.Inf(1)
    minIndex := -1
    for index,_ := range cs.Clusters {
        dist := distance.ComputeL2Dist(point, cs.centroids[index])
        if dist < minDist {
            minDist = dist
            minIndex = index
        }
    }
    if minIndex < 0 {
        panic("Point not clustered!")
    }

    cs.assign(point, minIndex)
}

func (cs ClusterSet) assign(point myMath.Point, clusterIndex int) {
    mappedMinIndex := clusterIndex + 1 
    prevMappedIndex := cs.assignment[point]
    
    if prevMappedIndex == mappedMinIndex {
        // point is already assigned to this cluster
        return
    }

	// add point to cluster
    cs.Clusters[clusterIndex].Add(point)
    cs.assignment[point] = clusterIndex + 1

    // check to remove from old cluster    
    if prevMappedIndex != 0 {
        cs.Clusters[prevMappedIndex - 1].Remove(point)
    }
}


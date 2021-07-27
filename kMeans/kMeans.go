package kMeans

import (
	"fmt"
	"math"
)

//Point represent the multidimensional data
type Point []float64

//Compute the distance between the centroid and the points
func distance(x Point, y Point) float64 {
	return math.Sqrt(math.Pow(x[0]-y[0],2) + math.Pow(x[1]-y[1],2))	 
}

func Main() {

	//Set of data points
	points := []Point{
		Point{3, 3},
		Point{2, 3},
		Point{2, 2},
		Point{1, 1},
		Point{0, 0},
		Point{1, -1},
		Point{-1, -1},
		Point{-3, -2},
		Point{-2, 0},
		Point{-2, -2},		
	}

	//Initialize centroids to some of the existing data points
	centroids := []Point{
		Point{2, 2},
		Point{1, 1},
		Point{-3, -2},
		Point{-2, -2},		
	}

	//Call the kmeans algorithm
	maxIter := 10
	eta := 0.4
	centroids = kmeans(points, centroids, maxIter, eta)

	//Print new centroids
	fmt.Println(centroids)
	//Print map of centroids and corresponding points
	fmt.Println(classify(points, centroids))
}

func kmeans(points []Point, centroids []Point, maxIter int, eta float64) []Point {

	centroidsMap := classify(points, centroids)
	updatedCentroids := updateCentroids(centroids,centroidsMap)
	if (maxIter>0) {
		return kmeans(points, updatedCentroids, maxIter-1, eta)
	} else {
		return updatedCentroids
	}

	// for _,v := range points{
	// 	currentPoint := v
	// 	fmt.Println("Point:",currentPoint)
	
	// 	for i := 0; i< len(centroids); i++ {
	// 		currentCentroid:= centroids[i]
	// 		fmt.Println("Centroid:",currentCentroid)
		
	// 		fmt.Println("Distance:",distance(currentCentroid,currentPoint))
	// 	}
	// }	
}

//Result contains the 'details of a point', 'nearest centroid index' and their distance
type Result struct {
	point Point
	centroidInd int
	distance float64
}

//Find the closest centroid to a point
func closest(p Point, centroids []Point) int {
	ch := make (chan Result)

	for i := 0; i < len(centroids); i++ {
		go func(j int, ch chan Result) {
			ch <- Result {
				centroidInd: j,
				distance:    distance(p,centroids[j]),
			}
		}(i, ch)
	}

	result := <-ch
	minDistance := result.distance
	minInd := result.centroidInd

	for i := 1; i < len(centroids); i++ {
		result = <-ch
		if result.distance < minDistance {
			minDistance = result.distance
			minInd = result.centroidInd
		}
	}

	return minInd
}

//Classify the data points to their nearest centroids
func classify(points []Point, centroids []Point) map[int][]Point {
	ch := make(chan Result)

	for i := range points {
		go func(j int, ch chan Result) {
			ch<- Result {				
				centroidInd: closest(points[j], centroids),
				point: points[j],
			}
		}(i, ch)
	}

	centroidsMap := make(map[int][]Point)

	for i := 0; i < len(points); i++ {
		result := <-ch
		centroidsMap[result.centroidInd] = append(centroidsMap[result.centroidInd], result.point)
	}
	return centroidsMap
}

//Compute new centroids by averaging data points corresponding to each centroid
func updateCentroids(oldCentroids []Point, centroidsMap map[int][]Point) []Point {

	ch := make(chan Result)
	for i := 0; i < len(oldCentroids); i++ {
		go func(j int, ch chan Result) {
			if len(centroidsMap[j]) == 0 {
				ch <- Result{
					centroidInd: j,
					point:       oldCentroids[j],
				}
			} else {
				ch <- Result{
					centroidInd: j,
					point:       getAverage(centroidsMap[j]),
				}
			}
		}(i, ch)
	}

	newCentroids := make([]Point, len(oldCentroids))
	for i := 0; i < len(oldCentroids); i++ {
		result := <-ch
		newCentroids[result.centroidInd] = result.point
	}

	return newCentroids
}

//Compute average point for a given list of points
func getAverage(points []Point) Point {
	newPoint := make(Point, len(points[0]))
	for _, point := range points {
		for j, val := range point {
			newPoint[j] += val
		}
	}
	for i := range newPoint {
		newPoint[i] = newPoint[i]/float64(len(points))
	}
	return newPoint
}
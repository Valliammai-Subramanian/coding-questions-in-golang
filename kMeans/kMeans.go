package kMeans

import (
	"fmt"
)

//Point represent the multidimensional data
type Point []float64

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
	updatedCentroids := []Point{}
	return updatedCentroids
}

//Classify the data points to their nearest centroids
func classify(points []Point, centroids []Point) map[int][]Point {
	centroidsMap := make(map[int][]Point)
	return centroidsMap
}

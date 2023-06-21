package main

import (
	"fmt"
	"math/rand"
)

func displayPoints(points []Point) {
	for index, point := range points {
		fmt.Printf("%3d: Point { %6f, %6f} \n", index, point.x, point.y)
	}
}

func main() {
	var numPoints int = 100
	var points = [100]Point{}
	for i := 0; i < numPoints; i++ {
		tempX := rand.Float64() * 100.0
		tempY := rand.Float64() * 100.0
		points[i] = Point{x: tempX, y: float64(tempY)}
	}

	fmt.Println("Welcome to the kmeans visualization...")
	displayPoints(points[:])

}

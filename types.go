package main

type Point struct {
	x float64
	y float64
}

type Cluster struct {
	center Point
	points []Point
}

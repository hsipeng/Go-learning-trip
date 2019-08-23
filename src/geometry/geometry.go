package main

import (
	"fmt"
	"log"

	"./rectangle"
)

var rectLen, rectWidth float64 = 6, 7

func init() {
	fmt.Println("main package init ...")

	if rectLen < 0 {
		log.Fatal("length is less than zero")
	}

	if rectWidth < 0 {
		log.Fatal("width is less than zero")
	}
}

func main() {
	fmt.Println("hello world!")

	// Area
	fmt.Printf("area of rectangle %.2f\n", rectangle.Area(rectLen, rectWidth))

	fmt.Printf("Diagonal of rectangle %.2f\n", rectangle.Diagonal(rectLen, rectWidth))
}

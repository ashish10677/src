package main

import (
	"fmt"
	"geometry/rectangle"
	"log"
)

var length, breadth float64 = 9, 8

func init() {
	fmt.Println("Main package initialized")
	if length < 0 {
		log.Fatal("Length cannot be less than 0")
	}

	if breadth < 0 {
		log.Fatal("Breadth cannot be less than 0")
	}
}

func main() {
	fmt.Println("Area:", rectangle.Area(length, breadth))
	fmt.Printf("Diagonal: %.2f\n", rectangle.Diagonal(length, breadth))
}

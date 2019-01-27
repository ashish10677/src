package rectangle

import (
	"fmt"
	"math"
)

func init() {
	fmt.Println("Rectangle package initialized")
}

//Area function determines area of a rectangle
func Area(l, b float64) float64 {
	a := l * b
	return a
}

//Diagonal function determines diagonal length of a rectangle
func Diagonal(l, b float64) float64 {
	d := math.Sqrt((l * l) + (b * b))
	return d
}

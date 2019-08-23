package rectangle

import (
	"fmt"
	"math"
)

func init() {
	fmt.Println("rectangle init ....")
}

// Area 求面积
func Area(len, wid float64) float64 {
	area := len * wid
	return area
}

// Diagonal 求平方
func Diagonal(len, wid float64) float64 {
	diagonal := math.Sqrt((len * len) + (wid * wid))
	return diagonal
}

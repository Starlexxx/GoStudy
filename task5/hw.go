package geometry
// hw is a bad name for a package, so I renamed it to geometry

import (
	"math"
)

// Task conditions: coordinates cannot be less than 0.

// stuct is for distance calculation is an overengineered solution
// type Geom struct {
// 	X1, Y1, X2, Y2 float64
// }

// CalculateDistance is too long name for a one line function
func Distance(x1, y1, x2, y2 float64) (distance float64) {
	// condition said that coordinates cannot be less than 0, so we should not check it
	// if geom.X1 < 0 || geom.X2 < 0 || geom.Y1 < 0 || geom.Y2 < 0 {
	// 	fmt.Println("Координаты не могут быть меньше нуля")
	// 	return -1

	return math.Sqrt(math.Pow(x2-x1, 2) + math.Pow(y2-y1, 2))
}

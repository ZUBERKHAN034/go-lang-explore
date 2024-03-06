package main

import (
	"fmt"
	"math"
)

func main() {
	var f32 float32 = 1.123456789
	var f64 float64 = 1.123456789

	fmt.Printf("float32: %.9f\n", f32)
	fmt.Printf("float64: %.15f\n", f64)

	// Performing calculations
	f32Result := f32 * 10
	f64Result := f64 * 10

	fmt.Printf("float32 result: %.9f\n", f32Result)
	fmt.Printf("float64 result: %.15f\n", f64Result)

	// Special values
	fmt.Printf("float32 max value: %f\n", math.MaxFloat32)
	fmt.Printf("float64 max value: %f\n", math.MaxFloat64)
}

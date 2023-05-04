package main

import (
	"fmt"
	"math"
)

func main() {
	var r float64
	fmt.Scan(&r)
	fmt.Printf("VOLUME = %.3f\n", ((4 / 3.0) * 3.14159 * math.Pow(r, 3)))
}

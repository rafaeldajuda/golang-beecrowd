package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b, c, d float64
	fmt.Scan(&a, &b, &c, &d)
	fmt.Printf("%.4f\n", math.Sqrt((math.Pow((c-a), 2) + math.Pow((d-b), 2))))
}

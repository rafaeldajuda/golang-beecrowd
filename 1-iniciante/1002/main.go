package main

import (
	"fmt"
	"math"
)

const n = 3.14159

func main() {

	var raio float64
	fmt.Scan(&raio)
	fmt.Printf("A=%.4f\n", n*(math.Pow(raio, 2)))

}

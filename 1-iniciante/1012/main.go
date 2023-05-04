package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b, c float64
	fmt.Scan(&a, &b, &c)
	fmt.Printf("TRIANGULO: %.3f\n", (a * c / 2))
	fmt.Printf("CIRCULO: %.3f\n", (3.14159 * math.Pow(c, 2)))
	fmt.Printf("TRAPEZIO: %.3f\n", ((a + b) * c / 2))
	fmt.Printf("QUADRADO: %.3f\n", (b * b))
	fmt.Printf("RETANGULO: %.3f\n", (a * b))
}

package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b, c int
	fmt.Scan(&a, &b, &c)

	maior := (a + b + int(math.Abs(float64(a)-float64(b)))) / 2
	if c > maior {
		maior = c
	}

	fmt.Printf("%d eh o maior\n", maior)
}

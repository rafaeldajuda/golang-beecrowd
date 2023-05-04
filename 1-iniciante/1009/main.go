package main

import "fmt"

func main() {
	var a string
	var b, c float64
	fmt.Scan(&a, &b, &c)
	fmt.Printf("TOTAL = R$ %.2f\n", (b + (c * 15 / 100)))
}

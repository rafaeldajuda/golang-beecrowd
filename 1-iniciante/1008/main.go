package main

import "fmt"

func main() {
	var a, b, c float64
	fmt.Scan(&a, &b, &c)
	fmt.Printf("NUMBER = %.0f\n", a)
	fmt.Printf("SALARY = U$ %.2f\n", (b * c))
}

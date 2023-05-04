package main

import "fmt"

func main() {
	var a, b, c, d, e, f float64
	fmt.Scan(&a, &b, &c, &d, &e, &f)
	fmt.Printf("VALOR A PAGAR: R$ %.2f\n", (b*c + e*f))
}

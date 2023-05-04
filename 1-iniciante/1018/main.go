package main

import (
	"fmt"
)

func main() {
	var notas = []int{100, 50, 20, 10, 5, 2, 1}
	var a int
	fmt.Scan(&a)

	fmt.Printf("%d\n", a)
	for _, n := range notas {
		fmt.Printf("%d nota(s) de R$ %d,00\n", int(a)/n, n)
		a = a % n
	}
}

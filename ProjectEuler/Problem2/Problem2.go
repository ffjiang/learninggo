// Project Euler Problem 2
// Even Fibonacci numbers

package main

import "fmt"

func main() {
	x := 1
	y := 1
	sum := 0

	for y < 4000000 {
		if y % 2 == 0 {
			sum += y
		}
		temp := x
		x = y
		y = x + temp
	}

	fmt.Printf("%d\n", sum)
}


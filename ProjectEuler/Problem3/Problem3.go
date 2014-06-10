// Project Euler Problem 3
// Largest Prime Factor
//
// What is the largest prime factor of the nubmer 600851475143

package main

import (
	"fmt"
	"math"
	"os"
)

var num int = 600851475143

func main() {
	fmt.Printf("%d\n", num)
	for i := 2; i <= num / 2; i++ {
		if isFactor(i) {
			if isPrime(num / i) {
				fmt.Printf("%d\n", num / i)
				os.Exit(0)
			}
		}
	}
}

func isFactor(divisor int) bool {
	if (num % divisor == 0) {
		return true
	} else {
		return false
	}
}

func isPrime(quotient int) bool {
	sqrt := int(math.Sqrt(float64(quotient)))
	for i := 2; i < sqrt; i++ {
		if quotient % i == 0 {
			return false
		}
	}

	return true
}
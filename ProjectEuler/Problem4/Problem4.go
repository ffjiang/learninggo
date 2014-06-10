// Project Euler Problm 4
// Largest palindrome product
//
// Find the largest palindrome made form the product of two 3-digit numbers.

package main

import (
	"fmt"
	"strconv"
)

func main() {
	largest := 0

	for x := 999; x > 0; x-- {
		for y := 999; y > 0; y-- {
			if x * y > largest {
				if isPalindrome(x * y) {
					largest = x * y;
				}
			} else {
				break
			}
		}
	}

	fmt.Printf("%d\n", largest)
}

func isPalindrome(num int) bool {
	str := strconv.Itoa(num)
	length := len(str)

	for i := 0; i < length / 2; i++ {
		if str[i] != str[length - 1 - i] {
			return false
		}
	}

	return true


}
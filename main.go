package main

import (
	"fmt"
	"math"
)

func IsPrimo(n int) bool {
	if n < 2 {
		return false
	}

	if n == 2 {
		return true
	}

	if n%2 == 0 {
		return false
	}

	limite := int(math.Sqrt(float64(n)))
	for i := 3; i <= limite; i += 2 {
		if n%i == 0 {
			return false
		}
	}

	return true
}

func main() {
	fmt.Println("Hello, World!")

	x := 0
	for {
		if IsPrimo(x) {
			fmt.Println("El numero ", x, " es primo")
		}
		x++
	}
}

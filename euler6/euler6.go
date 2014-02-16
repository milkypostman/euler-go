package main

import "fmt"

func main() {
	sumOfSquares := 0
	squaredSum := 0
	for i := 1; i <= 100; i++ {
		sumOfSquares += i * i
		squaredSum += i
	}
	fmt.Printf("%v\n", squaredSum*squaredSum-sumOfSquares)
}

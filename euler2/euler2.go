package main

import "fmt"

func main() {
	sum := 0
	for f1, f2 := 1, 1; f2 <= 4000000; f1, f2 = f2, f1+f2 {
		if f2%2 == 0 {
			sum += f2
		}
	}
	fmt.Printf("%v", sum)
}

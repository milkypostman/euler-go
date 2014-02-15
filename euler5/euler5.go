package main

import "fmt"

func evenlyDiv(num int) bool {
	for i := 2; i <= 20; i++ {
		if num%i != 0 {
			return false
		}
	}
	return true
}

func main() {
	smallest := 0
	for i := 2; ; i++ {
		if evenlyDiv(i) {
			smallest = i
			break
		}
	}
	fmt.Printf("%v\n", smallest)
}

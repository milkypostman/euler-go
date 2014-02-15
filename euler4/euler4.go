package main

import "fmt"
import "strconv"

func isPalindrome(p string) bool {
	plen := len(p)
	for i, j := 0, plen-1; i <= j; i, j = i+1, j-1 {
		if p[i] != p[j] {
			return false
		}
	}
	return true
}

func main() {
	max := 0
	for n := 100; n < 1000; n++ {
		for m := 100; m < 1000; m++ {
			prod := n * m
			if isPalindrome(strconv.Itoa(prod)) && prod > max {
				max = prod
			}
		}
	}
	fmt.Printf("%v\n", max)
}

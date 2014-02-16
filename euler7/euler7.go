package main

import "fmt"

func isprime(val int, known []int) bool {
	for i := 0; i < len(known); i++ {
		if val%known[i] == 0 {
			return false
		}
	}
	return true
}

func main() {
	primes := []int{2}
	nextval := 3
	for len(primes) < 10001 {
		if isprime(nextval, primes) {
			primes = append(primes, nextval)
		}
		nextval++
	}
	fmt.Printf("%v\n", primes[len(primes)-1])
}

package main

import "fmt"

var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

func main() {
	// range returns index of iterated element and copy of the element
	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}

	// one of values could be skipped
	for _, value := range pow {
		fmt.Printf("%d\n", value)
	}
}

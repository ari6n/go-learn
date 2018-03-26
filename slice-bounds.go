package main

import "fmt"

func main() {
	var a [10]int
	fmt.Println(a[:10])
	fmt.Println(a[0:])
	fmt.Println(a[:])

	s := []int{2, 3, 5, 7, 11, 13}

	s = s[1:4]
	fmt.Println(s)

	s = s[:2]
	fmt.Println(s)

	s = s[1:]
	fmt.Println(s)
}

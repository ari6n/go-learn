package main

import "fmt"

func main() {
	// names := [4]string{
	names := [...]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}
	fmt.Println(names)

	a := names[0:2]
	b := names[1:3]
	fmt.Println(a, b)
	c := append(b, "Alex", "Susie", "Mike")

	b[0] = "XXX"
	fmt.Println(a, b, c)
	printSlice(a)
	printSlice(b)
	printSlice(c)
	fmt.Println(names)
}

func printSlice(s []string) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

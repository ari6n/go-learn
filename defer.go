package main

import "fmt"

func main() {
	i := 1
	i++
	defer fmt.Println("world", i)
	i++
	fmt.Println("hello", i)
}

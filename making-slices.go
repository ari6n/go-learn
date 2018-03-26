package main

import "fmt"

func main() {
	a := make([]int, 5)
	printSlice("a", a)

	b := make([]int, 0, 5)
	printSlice("b", b)

	c := b[:2]
	for i := range c {
		c[i] = i
	}
	printSlice("c", c)

	d := c[2:5]
	printSlice("d", d)

	c = c[:cap(c)]
	printSlice("c", c)
	for i := range c {
		c[i] = i + 2
	}

	b = b[:cap(b)]
	printSlice("b", b)
	printSlice("c", c)
	printSlice("d", d)

}

func printSlice(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n",
		s, len(x), cap(x), x)
}

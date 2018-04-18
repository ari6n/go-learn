package main

import "fmt"

func main() {
	a := make([]int, 5)
	PrintSlice("a", a)

	b := make([]int, 0, 5)
	PrintSlice("b", b)

	c := b[:2]
	for i := range c {
		c[i] = i
	}
	PrintSlice("c", c)

	d := c[2:5]
	PrintSlice("d", d)

	c = c[:cap(c)]
	PrintSlice("c", c)
	for i := range c {
		c[i] = i + 2
	}

	b = b[:cap(b)]
	PrintSlice("b", b)
	PrintSlice("c", c)
	PrintSlice("d", d)

}

func PrintSlice(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n",
		s, len(x), cap(x), x)
}

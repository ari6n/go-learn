package main

import "golang.org/x/tour/pic"
//import "fmt"

func Pic(dx, dy int) [][]uint8 {
	res := make([][]uint8, dx)
	for x := range res {
		res[x] = make([]uint8, dy)
		for y := range res[x]{
			res[x][y] = uint8(x^y)
		}
	}
	return res
}

func main() {
	pic.Show(Pic)
}

package main

import (
	"fmt"
	"math"
)

func Sqrt1(x float64) float64 {
	res := float64(1)
	for i, tmp := 0, 0.0; i < 10; i++ {
		tmp = res - (res*res-x)/2*res
		fmt.Println("i = ", i, ", res = ", res, ", diff = ", math.Abs(res-tmp))
		res = tmp
	}
	return res
}

func Sqrt2(x float64) float64 {
	res := float64(1)
	dm := math.Abs(x / 500)
	fmt.Println("DiffMin = ", dm)
	for i, dif, tmp := 0, 100.0, 0.0; dif > dm && i < 5000; i++ {
		tmp = res - (res*res-x)/2*res
		dif = math.Abs(res - tmp)
		fmt.Println("i = ", i, ", res = ", res, ", diff = ", dif)
		res = tmp
	}
	return res
}

func main() {
	n := 2.0
	fmt.Println("Evaluated: ", Sqrt2(n))
	fmt.Println("From math: ", math.Sqrt(n))
}

package main

import (
  "fmt"
  "io/ioutil"
  "log"
  "regexp"
)

var digitRegexp = regexp.MustCompile("[0-9]+")

func main() {
  p := []int{2, 3, 5}
  p = append(p, 6, 7, 8, 9)

  // x := Filter(p, isEven)
  x := Filter(p, func(i int) bool {
    return bool((i/2)*2 == i)
  })

  printSlice("X", x)

  b := CopyDigits("/Users/ogolovko/tmp/BPM_Jun_SRX100.cfg")
  printSliceB("B", b)
}

func FindDigits(filename string) []byte {
  b, err := ioutil.ReadFile(filename)
  if err != nil {
		log.Fatal(err)
	}
  return digitRegexp.Find(b)
}

func CopyDigits(filename string) []byte {
    b, _ := ioutil.ReadFile(filename)
    b = digitRegexp.Find(b)
    c := make([]byte, len(b))
    copy(c, b)
    return c
}

func isEven(i int) bool {
  return bool((i/2)*2 == i)
}

// Filter returns a new slice holding only
// the elements of s that satisfy f()
func Filter(s []int, fn func(int) bool) []int {
  var p []int // == nil
  for _, v := range s {
    if fn(v) {
        p = append(p, v)
    }
  }
  return p
}

func printSlice(name string, s []int) {
	fmt.Printf("Slice %s len=%d cap=%d %v\n", name, len(s), cap(s), s)
}

func printSliceB(name string, s []byte) {
	fmt.Printf("Slice %s len=%d cap=%d %v\n", name, len(s), cap(s), s)
}

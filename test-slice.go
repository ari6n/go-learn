package main

import "fmt"

func main() {
	var s []byte
	s = make([]byte, 5, 5)
	// s == []byte{0, 0, 0, 0, 0}
	for i := range s {
		s[i] = byte(i + 3)
	}
	printSlice("S", s)

	t := make([]byte, len(s), (cap(s)+1)*2) // +1 in case cap(s) == 0
	for i, v := range s {
		t[i] = s[i]
		fmt.Printf("s[%d] = %d\n", i, v)
	}
	s = t

	x := AppendByte(s, 5, 4, 3)

	printSlice("X", x)

	p := []byte{2, 3, 5}
	p = AppendByte(p, 7, 11, 13)
	printSlice("P", p)

	p = append(p, s...)
	printSlice("P'", p)
}

func AppendByte(slice []byte, data ...byte) []byte {
	m := len(slice)
	n := m + len(data)
	if n > cap(slice) { // if necessary, reallocate
		// allocate double what's needed, for future growth.
		newSlice := make([]byte, (n+1)*2)
		copy(newSlice, slice)
		slice = newSlice
	}
	slice = slice[0:n]
	copy(slice[m:n], data)

	return slice
}

func printSlice(name string, s []byte) {
	fmt.Printf("Slice %s len=%d cap=%d %v\n", name, len(s), cap(s), s)
}

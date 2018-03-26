package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Print("Go runs on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		// freebsd, openbsd,
		// plan9, windows...
		fmt.Printf("%s.", os)
	}
}

/* other cases:

// No expression (same as "switch true ...")
func unhex(c byte) byte {
	switch {
	case '0' <= c && c <= '9':
		return c - '0'
	case 'a' <= c && c <= 'f':
		return c - 'a' + 10
	case 'A' <= c && c <= 'F':
		return c - 'A' + 10
	}
	return 0
}

// Break
command := ReadCommand()
argv := strings.Fields(command)
switch argv[0] {
case "echo":
	fmt.Print(argv[1:]...)
case "cat":
	if len(argv) <= 1 {
		fmt.Println("Usage: cat <filename>")
		break
	}
	PrintFile(argv[1])
default:
	fmt.Println("Unknown command; try 'echo' or 'cat'")
}

// fallthrough (fallthrough must be last command in case-section)
// Unpack 4 bytes into uint32 to repack into base 85 5-byte.
var v uint32
switch len(src) {
default:
	v |= uint32(src[3])
	fallthrough
case 3:
	v |= uint32(src[2]) << 8
	fallthrough
case 2:
	v |= uint32(src[1]) << 16
	fallthrough
case 1:
	v |= uint32(src[0]) << 24
}

// work around this by using a 'labeled' fallthrough:
switch {
case f():
	if g() {
		goto nextCase // Works now!
	}
	h()
    break
nextCase:
    fallthrough
default:
	error()
}

// multiple cases
func letterOp(code int) bool {
	switch chars[code].category {
	case "Lu", "Ll", "Lt", "Lm", "Lo":
		return true
	}
	return false
}

// Type switch
func typeName(v interface{}) string {
	switch v.(type) {
	case int:
		return "int"
	case string:
		return "string"
	default:
		return "unknown"
	}
}
// or
func do(v interface{}) string {
	switch u := v.(type) {
	case int:
		return strconv.Itoa(u*2) // u has type int
	case string:
		mid := len(u) / 2 // split - u has type string
		return u[mid:] + u[:mid] // join
	}
	return "unknown"
}

do(21) == "42"
do("bitrab") == "rabbit"
do(3.142) == "unknown"

// Noop case
func pluralEnding(n int) string {
	ending := ""

	switch n {
	case 1:
	default:
		ending = "s"
	}

	return ending
}

fmt.Sprintf("foo%s\n", pluralEnding(1))  == "foo"
fmt.Sprintf("bar%s\n", pluralEnding(2))  == "bars"

*/

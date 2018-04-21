/* Exercise for maps:
Implement WordCount. It should return a map of the counts of each “word” in the string s.
The wc.Test function runs a test suite against the provided function and prints success or failure.
*/
package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	//fmt.Printf("Fields are: %q", strings.Fields(s))
	res := make(map[string]int)
	for _, ss := range strings.Fields(s){
		res[ss] ++;
	}
	return res
}

func main() {
	wc.Test(WordCount)
}


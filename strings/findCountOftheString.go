/ You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"
	"strings"
)

func main() {

	s := "go is fast and go is simple"

	val := CountOccurenceofWord(s)
	for k, v := range val {
		fmt.Printf("%s:%d\n", k, v)

	}

}

func CountOccurenceofWord(s string) map[string]int {
	words := strings.Fields(s)
	freq := make(map[string]int)
	for _, count := range words {
		freq[count]++
	}
	return freq

}
package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

// WordCount : Returns a map with the occurences of a word in sring
func WordCount(s string) (counted map[string]int) {
	words := strings.Fields(s)
	counted = make(map[string]int)

	for _, val := range words {
		counted[val]++
	}
	return
}

func main() {
	wc.Test(WordCount)
}

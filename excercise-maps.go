package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	words := strings.Fields(s)
	wcMap := make(map[string]int)
	for _,word := range words {
		count, _ := wcMap[word]
		count+=1
		wcMap[word] = count;
	}
	return wcMap
}

func main() {
	wc.Test(WordCount)
}

package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	var fields []string = strings.Fields(s)
	m := make(map[string]int)
	for _, word := range fields {
		v, ok := m[word]
		if ok == true {
			m[word] = v + 1
		} else {
			m[word] = 1
		}
	}
	return m
}

func main() {
	wc.Test(WordCount)
}
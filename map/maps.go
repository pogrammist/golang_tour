package main

import (
	"strings"

	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	entries := strings.Split(s, " ")
	m := make(map[string]int)
	for _, e := range entries {
		m[e] = 1
	}
	return m
}

func main() {
	wc.Test(WordCount)
}

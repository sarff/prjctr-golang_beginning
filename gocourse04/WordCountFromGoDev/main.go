package main

import (
	"fmt"
	"strings"

	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	separateWord := strings.Fields(s)
	m := make(map[string]int)
	for _, word := range separateWord {
		v, _ := m[word]
		m[word] = v + 1
	}
	return m
}

func main() {
	fmt.Println(WordCount("hello world hello hello word world"))
	wc.Test(WordCount)
}

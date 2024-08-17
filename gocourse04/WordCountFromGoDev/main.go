package main

import (
	"fmt"
	"golang.org/x/tour/wc"
	"strings"
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

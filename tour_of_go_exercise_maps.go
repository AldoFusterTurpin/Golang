//ALDO FUSTER TURPIN

package main

import (
	"strings"

	"golang.org/x/tour/wc"
)

func WordCount_Go_is_great(s string) map[string]int {
	m := make(map[string]int)

	for _, c := range strings.Fields(s) {
		m[c] += 1
	}
	return m
}

func WordCount(s string) map[string]int {
	m := make(map[string]int)

	for _, c := range strings.Fields(s) {
		_, ok := m[c]
		if ok {
			m[c] += 1
		} else {
			m[c] = 1
		}
	}
	return m
}

func WordCount_v2(s string) map[string]int {
	m := make(map[string]int)

	for _, c := range strings.Fields(s) {
		elem, ok := m[c]
		if ok {
			m[c] = elem + 1
		} else {
			m[c] = 1
		}
	}
	return m
}

func main() {
	wc.Test(WordCount_Go_is_great)
}

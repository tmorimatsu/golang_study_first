package main

import (
	"fmt"
	"github.com/labstack/gommon/log"
	"sort"
)

var prereqs = map[string][]string{
	"algorithms": {"data structures"},
	"calculus":   {"linear algebra"},

	"compilers": {
		"data structures",
		"formal languages",
		"computer organization",
	},

	"data structures":       {"discrete math"},
	"databases":             {"data structures"},
	"discrete math":         {"intro to programming"},
	"formal languages":      {"discrete math"},
	"networks":              {"operating systems"},
	"operating systems":     {"data structures", "computer organization"},
	"programming languages": {"data structures", "computer organization"},
	"linear algebra":        {"calculus"}, // 循環を作成
	// "intro to programming": {"formal languages"},
}

func main() {
	for i, course := range topoSort(prereqs) {
		fmt.Printf("%d:\t%s\n", i+1, course)
	}
}

func topoSort(m map[string][]string) []string {
	var order []string
	try := make(map[string]bool)
	seen := make(map[string]bool)
	var visitAll func(items []string)

	visitAll = func(items []string) {
		for _, item := range items {
			if !seen[item] {
				if try[item] {
					log.Fatal("cyclic")
				}
				try[item] = true
				visitAll(m[item])
				order = append(order, item)
				seen[item] = true
			}
		}
	}

	var keys []string
	for key := range m {
		fmt.Println(key)
		keys = append(keys, key)
	}
	sort.Strings(keys)
	visitAll(keys)
	return order
}

package main

import (
	"fmt"
	"log"
)

func main() {
	max, err := max(1, 0, 9, -19, 20)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(max)

	min, err := min(1, 0, 9, -19, 20)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(min)
}

func max(vals ...int) (int, error) {
	if len(vals) == 0 {
		return 0, fmt.Errorf("no arguments is given")
	}
	m := vals[0]
	for _, val := range vals[1:] {
		if m < val {
			m = val
		}
	}
	return m, nil
}

func min(vals ...int) (int, error) {
	if len(vals) == 0 {
		return 0, fmt.Errorf("no arguments is given")
	}
	m := vals[0]
	for _, val := range vals[1:] {
		if m > val {
			m = val
		}
	}
	return m, nil
}

func sum(vals ...int) int {
	total := 0
	for _, val := range vals {
		total += val
	}
	return total
}

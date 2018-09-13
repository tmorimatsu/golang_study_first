package main

import(
	"fmt"
	_"unicode/utf8"
)

func main() {

	var c ByteCounter
	c.Write([]byte("hello"), "test")
	fmt.Println(c)

	var r RowCounter
	rows, _ := r.Write([]byte(`testtest
		eee
		eee`))

	fmt.Printf("%d\n", rows)

	var w WordCounter
	words, _ := w.Write([]byte(`testtest
		eee
		eee`))

	fmt.Printf("%d\n", words)
}

type ByteCounter int
type RowCounter int
type WordCounter int

func (c *ByteCounter) Write(p []byte, t string) (int, error) {

	*c += ByteCounter(len(p))
	return len(p), nil
}

func (c *RowCounter) Write(p []byte) (int, error) {
	r := []rune(string(p))
	rows := 1
	for i := range r {
		if r[i] == '\n' {
			rows++
		}
	}
	return rows, nil
}

func (c *WordCounter) Write(p []byte) (int, error) {

	r := []rune(string(p))
	var words int
	for _ = range r {
		words++
	}
	return words, nil
}

func (c *ByteCounter) countRowsAndWords(s string) (int, int) {
	r := []rune(s)
	var words, rows int
	for i := range r {
		if r[i] == '\n' {
			rows++
		}
		words++
	}
	return rows, words
}


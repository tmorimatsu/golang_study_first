package main

import (
	"bytes"
	"fmt"
	"io"
	"strings"
)

func main() {
	r := LimitReader(strings.NewReader("tetete"), 3)
	buf := new(bytes.Buffer)
	buf.ReadFrom(r)
	fmt.Println(buf)
}

type limitReader struct {
	r        io.Reader
	n, limit int
}

// TODO:fixme
func (r *limitReader) Read(p []byte) (n int, err error) {
	lim := r.limit - r.n
	if len(p) < lim {
		lim = len(p)
	}
	n, err = r.r.Read(p[:lim])
	r.n += n
	if r.n >= r.limit {
		err = io.EOF
	}
	return
}

func LimitReader(r io.Reader, limit int) io.Reader {
	return &limitReader{r: r, limit: limit}
}

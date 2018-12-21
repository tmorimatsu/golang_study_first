package sexpr

import (
	"reflect"
	"testing"
)

func TestUnmarshal(t *testing.T) {
	type Movie struct {
		Title, Subtitle string
		Year            int
		Oscars          []string
		Sequel          *string
	}
	strangelove := Movie{
		Title:    "Dr. Strangelove",
		Subtitle: "How I Learned to Stop Worrying and Love the Bomb",
		Year:     1964,
		Oscars: []string{
			"Best Actor (Nomin.)",
			"Best Adapted Screenplay (Nomin.)",
			"Best Director (Nomin.)",
			"Best Picture (Nomin.)",
		},
	}

	data := "( (Title \"Dr. Strangelove\") (Subtitle \"How I Learned to Stop Worrying and Love the Bomb\") (Year 1964) (Oscars (\"Best Actor (Nomin.)\" \"Best Adapted Screenplay (Nomin.)\" (\"Best Director (Nomin.)\" (\"Best Picture (Nomin.)\")))"
	var movie Movie
	b := []byte(data)
	err := Unmarshal(b, &movie)
	if err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(movie, strangelove) {
		t.Fatal("not equal")
	}
}

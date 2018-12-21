package sexpr

import (
	"testing"
)

// TODO : 追加した分のテストコード

func TestMarshal(t *testing.T) {
	// type Movie struct {
	// 	Title, Subtitle string
	// 	Year            int
	// 	Actor           map[string]string
	// 	Oscars          []string
	// 	Sequel          *string
	// }
	// strangelove := Movie{
	// 	Title:    "Dr. Strangelove",
	// 	Subtitle: "How I Learned to Stop Worrying and Love the Bomb",
	// 	Year:     1964,
	// 	Actor: map[string]string{
	// 		"Dr. Strangelove":            "Peter Sellers",
	// 		"Grp. Capt. Lionel Mandrake": "Peter Sellers",
	// 		"Pres. Merkin Muffley":       "Peter Sellers",
	// 		"Gen. Buck Turgidson":        "George C. Scott",
	// 		"Brig. Gen. Jack D. Ripper":  "Sterling Hayden",
	// 		`Maj. T.J. "King" Kong`:      "Slim Pickens",
	// 	},
	// 	Oscars: []string{
	// 		"Best Actor (Nomin.)",
	// 		"Best Adapted Screenplay (Nomin.)",
	// 		"Best Director (Nomin.)",
	// 		"Best Picture (Nomin.)",
	// 	},
	// }
	ts := []struct {
		i        interface{}
		expected string
	}{
		{true, "t"},
		{false, "nil"},
		{float64(10.3), "10.3"},
		{complex(10, 3), "#C(10 3)"},
		{complex(-300, -300), "#C(-300 -300)"},
	}

	for _, k := range ts {
		actual, err := Marshal(k.i)
		if err != nil {
			t.Fatalf("Marshal failed: %v", err)
		}
		if k.expected != string(actual) {
			t.Fatalf("exptected %v, but actual %v", k.expected, actual)
		}
	}
}

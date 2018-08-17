package main

import "testing"

func TestComma(t *testing.T) {
	t.Log("3桁区切りでコンマが挿入されること")
	expected := "9,223,372,036,854,775,808"
	actual := comma("9223372036854775808")
	if actual != expected {
		t.Errorf("actual %v\nwant %v", actual, expected)
	}
}

func TestCommaWithSign(t *testing.T) {
	t.Log("符号付の数に対応していること")
	expected := "-1,234"
	actual := comma("-1234")
	if actual != expected {
		t.Errorf("actual %v\nwant %v", actual, expected)
	}
}

func TestCommaWithFloatingPoint(t *testing.T) {
	t.Log("浮動小数点の数に対応していること")
	expected := "+1,234.6543"
	actual := comma("+1234.6543")
	if actual != expected {
		t.Errorf("actual %v\nwant %v", actual, expected)
	}
}

func TestTmp(t *testing.T) {
	tests := []struct {
		s, want string
	}{
		{"", ""},
		{"1", "1"},
		{"12", "12"},
		{"123", "123"},
		{"1234", "1,234"},
		{"-1", "-1"},
		{"-1.2", "-1.2"},
		{"-123", "-123"},
		{"-12.34", "-12.34"},
		{"-1234", "-1,234"},
	}
	for _, test := range tests {
		got := comma(test.s)
		if got != test.want {
			t.Errorf("comma(%q), got %q, want %q", test.s, got, test.want)
		}
	}
}

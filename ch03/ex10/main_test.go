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

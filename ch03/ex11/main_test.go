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

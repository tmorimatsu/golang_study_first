package main

import "testing"

func TestReverseUTF8(t *testing.T) {
	t.Log("文字列を逆順にする")
	expected := "トステとすて"
	actual := string(reverseUTF8([]byte("てすとテスト")))
	if actual != expected {
		t.Errorf("actual %v\nwant %v", actual, expected)
	}
}

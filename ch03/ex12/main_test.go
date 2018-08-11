package main

import "testing"

func TestIsAnagram(t *testing.T) {
	t.Log("シングルバイト文字でアナグラムを検知できること")
	expected := true
	actual := isAnagram("testt", "ttest")
	if actual != expected {
		t.Errorf("actual %v\nwant %v", actual, expected)
	}
}

func TestIsAnagramFalse(t *testing.T) {
	t.Log("シングルバイト文字でアナグラムじゃない文字列を検知できること")
	expected := false
	actual := isAnagram("testt", "twest")
	if actual != expected {
		t.Errorf("actual %v\nwant %v", actual, expected)
	}
}

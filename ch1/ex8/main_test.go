package main

import "testing"

func TestAddUrlPrefix(t *testing.T) {
	expected := "http://test.com"
	actual := addUrlPrefix("test.com")
	if actual != expected {
		t.Errorf("actual %v\nwant %v", actual, expected)
	}
}
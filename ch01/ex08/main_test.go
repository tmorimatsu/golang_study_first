package main

import (
	"fmt"
	"testing"
)

func TestAddUrlPrefixWithNoPrefix(t *testing.T) {
	fmt.Println("http://, https://がついていないURLにはhttp://をつける")
	expected := "http://test.com"
	actual := addUrlPrefix("test.com")
	if actual != expected {
		t.Errorf("actual %v\nwant %v", actual, expected)
	}
}

func TestAddUrlPrefixWithHttp(t *testing.T) {
	fmt.Println("http://がついているURLはそのままにする")
	expected := "http://test.com"
	actual := addUrlPrefix("http://test.com")
	if actual != expected {
		t.Errorf("actual %v\nwant %v", actual, expected)
	}
}

func TestAddUrlPrefixWithHttps(t *testing.T) {
	fmt.Println("https://がついているURLはそのままにする")
	expected := "https://test.com"
	actual := addUrlPrefix("https://test.com")
	if actual != expected {
		t.Errorf("actual %v\nwant %v", actual, expected)
	}
}

package main

import (
	"crypto/sha256"
	"crypto/sha512"
	"testing"
)

func TestConvertToSHA256(t *testing.T) {
	expected := sha256.Sum256([]byte("tete"))
	actual := convertToSHA256("tete")
	if expected != actual {
		t.Errorf("actual %v\nwant %v", actual, expected)
	}
}

func TestConvertToSHA384(t *testing.T) {
	expected := sha512.Sum384([]byte("tete"))
	actual := convertToSHA384("tete")
	if expected != actual {
		t.Errorf("actual %v\nwant %v", actual, expected)
	}
}

func TestConvertToSHA512(t *testing.T) {
	expected := sha512.Sum512([]byte("tete"))
	actual := convertToSHA512("tete")
	if expected != actual {
		t.Errorf("actual %v\nwant %v", actual, expected)
	}
}

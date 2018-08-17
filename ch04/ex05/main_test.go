package main

import (
	"reflect"
	"testing"
)

func TestNonContiguousValues(t *testing.T) {
	arr := []string{"test", "tete1", "tete2", "tete2", "tete2", "tete2", "tete3", "test", "test"}
	expected := []string{"test", "tete1", "tete2", "tete3", "test"}
	actual := nonContiguousValues(arr)
	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("actual %v\nwant %v", actual, expected)
	}
}

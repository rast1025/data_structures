package main

import (
	"testing"
	"reflect"
)

func TestNullifyArr(t *testing.T) {
	input := [][]int {
		{1, 2, 3, 4},
		{0, 4, 7, 9},
		{2, 1, 1, 0},
		{4, 9, 8, 6},
	}

	expected := [][]int {
		{0, 2, 3, 0},
		{0, 0, 0, 0},
		{0, 0, 0, 0},
		{0, 9, 8, 0},
	}

	result := nullifyArray(input)

	if !reflect.DeepEqual(expected, result) {
		t.Errorf("expected %v, got %v", expected, result)
	}
}

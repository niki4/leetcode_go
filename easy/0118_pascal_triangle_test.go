package main

import (
	"reflect"
	"testing"
)

func TestGeneratePascalTriangle(t *testing.T) {
	expected := [][]int{
		{1},
		{1, 1},
		{1, 2, 1},
		{1, 3, 3, 1},
		{1, 4, 6, 4, 1},
	}
	if triangle := generatePascalTriangle(5); !reflect.DeepEqual(triangle, expected) {
		t.Errorf("\nwant:\t%v\ngot:\t%v", expected, triangle)
	}
}

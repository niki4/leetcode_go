package medium

import (
	"reflect"
	"testing"
)

func TestThreeSum(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want [][]int
	}{
		{
			name: "Example 1",
			nums: []int{-1, 0, 1, 2, -1, -4},
			want: [][]int{{-1, -1, 2}, {-1, 0, 1}},
		},
		{
			name: "Example 2",
			nums: []int{0, 1, 1},
			want: [][]int{},
		},
		{
			name: "Example 3",
			nums: []int{0, 0, 0},
			want: [][]int{{0, 0, 0}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := threeSum(tt.nums)
			if !equalTriplets(got, tt.want) {
				t.Errorf("threeSum(%v) = %v, want %v", tt.nums, got, tt.want)
			}
		})
	}
}

// equalTriplets checks if two slices of triplets are equal regardless of order.
func equalTriplets(a, b [][]int) bool {
	if len(a) != len(b) {
		return false
	}
	visited := make([]bool, len(b))
	for _, tripletA := range a {
		found := false
		for j, tripletB := range b {
			if visited[j] {
				continue
			}
			if reflect.DeepEqual(tripletA, tripletB) {
				visited[j] = true
				found = true
				break
			}
		}
		if !found {
			return false
		}
	}
	return true
}

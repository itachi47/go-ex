package main

import (
	"reflect"
	"testing"
)

func TestBubbleSort(t *testing.T) {
	tests := []struct {
		name  string
		input []int
		want  []int
	}{
		{
			name:  "basic unsorted",
			input: []int{5, 1, 9, 2, 7},
			want:  []int{1, 2, 5, 7, 9},
		},
		{
			name:  "already sorted",
			input: []int{1, 2, 3, 4, 5},
			want:  []int{1, 2, 3, 4, 5},
		},
		{
			name:  "reverse sorted",
			input: []int{5, 4, 3, 2, 1},
			want:  []int{1, 2, 3, 4, 5},
		},
		{
			name:  "duplicates",
			input: []int{4, 2, 4, 1, 2, 3},
			want:  []int{1, 2, 2, 3, 4, 4},
		},
		{
			name:  "negative numbers",
			input: []int{-1, 3, 0, -5, 2},
			want:  []int{-5, -1, 0, 2, 3},
		},
		{
			name:  "single element",
			input: []int{42},
			want:  []int{42},
		},
		// {
		// 	name:  "empty slice",
		// 	input: []int{},
		// 	want:  []int{},
		// },
	}

	for _, tt := range tests {
		got := append([]int(nil), tt.input...) // copy input

		BubbleSort(got)

		if !reflect.DeepEqual(got, tt.want) {
			t.Fatalf("%s: got %v want %v", tt.name, got, tt.want)
		}
	}
}

package main

import "testing"

func TestToCsvText(t *testing.T) {
	testCases := []struct {
		input    [][]int
		expected string
	}{
		{
			input: [][]int{
				{0, 1, 2, 3, 45},
				{10, 11, 12, 13, 14},
				{20, 21, 22, 23, 24},
				{30, 31, 32, 33, 34},
			},
			expected: "0,1,2,3,45\n10,11,12,13,14\n20,21,22,23,24\n30,31,32,33,34",
		},
		{
			input: [][]int{
				{-25, 21, 2, -33, 48},
				{30, 31, -32, 33, -34},
			},
			expected: "-25,21,2,-33,48\n30,31,-32,33,-34",
		},
		{
			input: [][]int{
				{5, 55, 5, 5, 55},
				{6, 6, 66, 23, 24},
				{666, 31, 66, 33, 7},
			},
			expected: "5,55,5,5,55\n6,6,66,23,24\n666,31,66,33,7",
		},
		{
			input:    [][]int{},
			expected: "",
		},
		{
			input: [][]int{
				{1},
			},
			expected: "1",
		},
		{
			input: [][]int{
				{1, 2, 3},
				{4, 5, 6},
			},
			expected: "1,2,3\n4,5,6",
		},
	}

	for _, tc := range testCases {
		result := ToCsvText(tc.input)
		if result != tc.expected {
			t.Errorf("Expected: %q, Got: %q", tc.expected, result)
		}
	}
}

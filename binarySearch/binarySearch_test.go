package binarySearch

import "testing"

func TestSearch(t *testing.T) {
	type testCase struct {
		input  []int
		target int
		result int
	}

	testCaseList := []testCase{
		{
			input:  []int{-1, 0, 3, 5, 9, 12},
			target: 9,
			result: 4,
		},
		{
			input:  []int{-1, 0, 3, 5, 9, 12},
			target: 2,
			result: -1,
		},
		{
			input:  []int{-8, -1, 0, 3, 5, 9, 12},
			target: 5,
			result: 4,
		},
		{
			input:  []int{-1, 0, 5},
			target: -1,
			result: 0,
		},
		{
			input:  []int{5},
			target: -5,
			result: -1,
		},
	}

	for _, test := range testCaseList {
		result := Search(test.input, test.target)
		if result != test.result {
			t.Log("got", result, "expected", test.result)
			t.Fail()
		}

		if !t.Failed() {
			t.Log("Success", result, "was the correct answer.")
		}
	}
}

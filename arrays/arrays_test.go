package arrays

import (
	_ "net/http/pprof"
	"testing"
)

func TestFindMaxConsecutiveOnes(t *testing.T) {
	type testCase struct {
		input  []int
		result int
	}

	testCaseList := []testCase{
		{
			input:  []int{1, 1, 0, 1, 1, 1},
			result: 3,
		},
		{
			input:  []int{1},
			result: 1,
		},
		{
			input:  []int{},
			result: 0,
		},
	}

	for _, test := range testCaseList {
		t.Logf("Testing %v", test.input)
		result := FindMaxConsecutiveOnes(test.input)
		if result != test.result {
			t.Errorf("Expected %d, got %v", test.result, result)
		}
	}
}

func TestFindNumbers(t *testing.T) {
	type testCase struct {
		input  []int
		result int
	}

	testCaseList := []testCase{
		{
			input:  []int{555, 901, 482, 1771},
			result: 1,
		},
	}

	for _, test := range testCaseList {
		t.Logf("Testing %v", test.input)
		result := FindNumbers(test.input)
		if result != test.result {
			t.Errorf("Expected %d, got %v", test.result, result)
		}
	}
}

func TestSortedSquares(t *testing.T) {
	type testCase struct {
		input  []int
		result []int
	}

	testCaseList := []testCase{
		{
			input:  []int{-7, -3, 2, 3, 11},
			result: []int{4, 9, 9, 49, 121},
		},
	}

	for _, test := range testCaseList {
		t.Logf("Testing %v", test.input)
		result := SortedSquares(test.input)
		for i, num := range result {
			if num != test.result[i] {
				t.Fail()
				break
			}
		}
	}
}

func TestDuplicateZeros(t *testing.T) {
	type testCase struct {
		input  []int
		result []int
	}

	testCaseList := []testCase{
		{
			input:  []int{1, 0, 2, 3, 0, 4, 5, 0},
			result: []int{1, 0, 0, 2, 3, 0, 0, 4},
		},
		{
			input:  []int{1, 0, 2, 3, 0, 0, 5, 0},
			result: []int{1, 0, 0, 2, 3, 0, 0, 0},
		},
	}

	for _, test := range testCaseList {
		t.Logf("Testing %v", test.input)
		DuplicateZeros(test.input)
		for i, num := range test.input {
			if num != test.result[i] {
				t.Fail()
				break
			}
		}
	}
}

func TestMerge(t *testing.T) {
	type testCase struct {
		num1   []int
		m      int
		num2   []int
		n      int
		result []int
	}

	testCaseList := []testCase{
		{
			num1:   []int{1, 2, 3, 0, 0, 0},
			m:      3,
			num2:   []int{2, 5, 6},
			n:      3,
			result: []int{1, 2, 2, 3, 5, 6},
		},
		{
			num1:   []int{1},
			m:      1,
			num2:   []int{},
			n:      0,
			result: []int{1},
		},
	}

	for _, test := range testCaseList {
		t.Logf("Testing %v (%d), %v (%d)", test.num1, test.n, test.num2, test.m)
		Merge(test.num1, test.m, test.num2, test.n)
		for i, num := range test.num1 {
			if num != test.result[i] {
				t.Fail()
				break
			}
		}
	}
}

func TestRemoveElement(t *testing.T) {
	type testCase struct {
		input  []int
		value  int
		result int
	}

	testCaseList := []testCase{
		{
			input:  []int{3, 2, 2, 3},
			value:  3,
			result: 2,
		},
		{
			input:  []int{0, 1, 2, 2, 3, 0, 4, 2},
			value:  2,
			result: 5,
		},
	}

	for _, test := range testCaseList {
		t.Logf("Testing %v (%d)", test.input, test.value)
		result := RemoveElement(test.input, test.value)
		if result != test.result {
			t.Logf("expected: %d - got: %d", test.result, result)
			t.Fail()
			break
		}

		// Test to make sure the number of elements left in the array are correct.
		// Any element that was removed was instead replaced with the rune '_' (95)
		if !t.Failed() {
			for _, num := range test.input[result:] {
				if num != 95 {
					t.Fail()
					break
				}
			}
		}
	}
}

func TestRemoveDuplicates(t *testing.T) {
	type testCase struct {
		input      []int
		value      int
		result     int
		resultList []int
	}

	testCaseList := []testCase{
		{
			input:      []int{2, 2, 3, 3},
			result:     2,
			resultList: []int{2, 3, 111, 111},
		},
		{
			input:      []int{0, 0, 1, 2, 2, 2, 3, 4},
			result:     5,
			resultList: []int{0, 1, 2, 3, 4, 111, 111, 111},
		},
		{
			input:      []int{1, 1, 2},
			result:     2,
			resultList: []int{1, 2, 111},
		},
	}

	for _, test := range testCaseList {
		t.Logf("Testing %v", test.input)
		result := RemoveDuplicates(test.input)
		if result != test.result {
			t.Logf("expected: %d - got: %d", test.result, result)
			t.Fail()
			break
		}

		if !t.Failed() {
			for i := 0; i < result; i++ {
				if test.input[i] != test.resultList[i] {
					t.Logf("expected: %d - got: %d", test.resultList[i], test.input[i])
					t.Fail()
					break
				}
			}
		}
	}
}

func TestCheckIfExist(t *testing.T) {
	type testCase struct {
		input  []int
		result bool
	}

	testCaseList := []testCase{
		{
			input:  []int{10, 2, 5, 3},
			result: true,
		},
		{
			input:  []int{3, 1, 7, 11},
			result: false,
		},
		{
			input:  []int{-20, 8, -6, -14, 0, -19, 14, 4},
			result: true,
		},
	}

	for _, test := range testCaseList {
		t.Logf("Testing %v", test.input)
		result := CheckIfExist(test.input)
		if result != test.result {
			t.Logf("expected: %v - got: %v", test.result, result)
			t.Fail()
			break
		} else {
			t.Logf("Success: output - %v", result)
		}
	}
}

func TestValidMountainArray(t *testing.T) {
	type testCase struct {
		input  []int
		result bool
	}

	testCaseList := []testCase{
		{
			input:  []int{2, 1},
			result: false,
		},
		{
			input:  []int{3, 5, 5},
			result: false,
		},
		{
			input:  []int{0, 2, 3, 4, 5, 2, 1, 0},
			result: true,
		},
		{
			input:  []int{0, 3, 2, 1},
			result: true,
		},
		{
			input:  []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
			result: false,
		},
	}

	for _, test := range testCaseList {
		t.Logf("Testing %v", test.input)
		result := ValidMountainArray(test.input)
		if result != test.result {
			t.Logf("expected: %v - got: %v", test.result, result)
			t.Fail()
			break
		} else {
			t.Logf("Success: output - %v", result)
		}
	}
}

func TestReplaceElements(t *testing.T) {
	type testCase struct {
		input  []int
		result []int
	}

	testCaseList := []testCase{
		{
			input:  []int{17, 18, 5, 4, 6, 1},
			result: []int{18, 6, 6, 6, 1, -1},
		},
	}

	for _, test := range testCaseList {
		t.Logf("Testing %v", test.input)
		result := ReplaceElements(test.input)
		for i, num := range result {
			if num != test.result[i] {
				t.Logf("expected: %v - got: %v", test.result[i], num)
				t.Fail()
				break
			}
		}
	}
}

func TestMoveZeroes(t *testing.T) {
	type testCase struct {
		input  []int
		result []int
	}

	testCaseList := []testCase{
		{
			input:  []int{0, 1, 0, 3, 12},
			result: []int{1, 3, 12, 0, 0},
		},
		{
			input:  []int{0, 0, 1},
			result: []int{1, 0, 0},
		},
	}

	for _, test := range testCaseList {
		t.Logf("Testing %v", test.input)
		MoveZeroes(test.input)
		if len(test.input) != len(test.result) {
			t.Fail()
			continue
		}
		for i, num := range test.input {
			if num != test.result[i] {
				t.Logf("expected: %v - got: %v", test.result[i], num)
				t.Fail()
				continue
			}
		}
	}
}

func TestSortArrayByParity(t *testing.T) {
	type testCase struct {
		input  []int
		result []int
	}

	testCaseList := []testCase{
		{
			input:  []int{3, 1, 2, 4},
			result: []int{2, 4, 3, 1},
		},
		{
			input:  []int{0},
			result: []int{0},
		},
	}

	for _, test := range testCaseList {
		t.Logf("Testing %v", test.input)
		result := SortArrayByParity(test.input)
		if len(result) != len(test.result) {
			t.Fail()
			continue
		}
		for i, num := range result {
			if num != test.result[i] {
				t.Logf("expected: %v - got: %v", test.result[i], num)
				t.Fail()
				continue
			}
		}
	}
}

func TestHeightChecker(t *testing.T) {
	type testCase struct {
		input  []int
		result int
	}

	testCaseList := []testCase{
		{
			input:  []int{1, 1, 4, 2, 1, 3},
			result: 3,
		},
		{
			input:  []int{5, 1, 2, 3, 4},
			result: 5,
		},
		{
			input:  []int{1, 2, 3, 4, 5},
			result: 0,
		},
	}

	for _, test := range testCaseList {
		t.Logf("Testing %v", test.input)
		result := HeightChecker(test.input)
		if result != test.result {
			t.Logf("running with input data: %+v", test.input)
			t.Fail()
			break
		}
	}
}

func TestThirdMax(t *testing.T) {
	type testCase struct {
		input  []int
		result int
	}

	testCaseList := []testCase{
		{
			input:  []int{3, 2, 1},
			result: 1,
		},
		{
			input:  []int{1, 2},
			result: 2,
		},
		{
			input:  []int{2, 2, 3, 1},
			result: 1,
		},
		{
			input:  []int{1, 1, 2},
			result: 2,
		},
		{
			input:  []int{2, 2, 2, 1},
			result: 2,
		},
		{
			input:  []int{1, 1, 1},
			result: 1,
		},
	}

	for _, test := range testCaseList {
		t.Logf("Testing %v", test.input)
		result := ThirdMax(test.input)
		if result != test.result {
			t.Logf("running with input data: %+v", test.input)
			t.Fail()
			break
		}
	}
}

func TestFindDisappearedNumbers(t *testing.T) {
	type testCase struct {
		input  []int
		result []int
	}

	testCaseList := []testCase{
		{
			input:  []int{4, 3, 2, 7, 8, 2, 3, 1},
			result: []int{5, 6},
		},
		{
			input:  []int{3, 3, 3, 3, 3, 3},
			result: []int{1, 2, 4, 5, 6},
		},
		{
			input:  []int{1, 1},
			result: []int{2},
		},
		{
			input:  []int{5, 3, 4, 1, 1, 1, 8, 8},
			result: []int{2, 6, 7},
		},
		{
			input:  []int{2, 2, 5, 5, 5, 5},
			result: []int{1, 3, 4, 6},
		},
		{
			input:  []int{1},
			result: []int{},
		},
	}

	for _, test := range testCaseList {
		t.Logf("Testing %v", test.input)
		result := FindDisappearedNumbers(test.input)

		if len(result) != len(test.result) {
			t.Fail()
			continue
		}

		for i, num := range result {
			if num != test.result[i] {
				t.Fail()
				continue
			}
		}
	}
}

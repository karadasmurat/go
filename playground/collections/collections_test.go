package collections

import "testing"

// data representing a test case:
// inputs, and expected output(s)
type TestCaseData struct {
	Collection  []int
	Q           int
	Expected_v  int
	Expected_ok bool
}

func TestContains(t *testing.T) {

	numbers_empty := []int{}
	numbers := []int{-1, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	// A collection representing different test scenarios for the function to be tested.
	cases := []TestCaseData{
		{numbers_empty, -1, 0, false},
		{numbers, -1, 0, true},
		{numbers, 0, 1, true},
		{numbers, 99, 0, false},
	}

	// For each scenario, call the function to be tested
	for _, v := range cases {

		result, ok := Contains(v.Collection, v.Q)
		if ok != v.Expected_ok || result != v.Expected_v {
			t.Errorf("Contains(%v, %v): %v, Expected (%v, %v)", v.Collection, v.Q, result, v.Expected_v, v.Expected_ok)
		}
	}
}

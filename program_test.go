package main

import "testing"

func TestSumNumbers(t *testing.T){
	result := SumNumbers(1,2,3)
	if result != 6 {
		t.Errorf("Expected sum of %v equals %v. But got %v.", []int{1,2,3}, 6, result)
	}
}

func TestTableSumNumbers(t *testing.T){
	listOfResults := []struct{
		input []int
		expected int
	}{
		{
			input:[]int{4},
			expected:4,
		},
		{
			input:[]int{12,0,3},
			expected:15,
		},
		{
			input:[]int{10,22,100,200,60},
			expected:392,
		},
		{
			input:[]int{20,30},
			expected:50,
		},
	}

	for _, v := range listOfResults {
		if result := SumNumbers(v.input...); result != v.expected {
			t.Errorf("Input: %v, expected: %v, result: %v", v.input, v.expected, result)
		}
	}
}

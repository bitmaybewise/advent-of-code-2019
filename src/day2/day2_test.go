package main

import (
	"testing"
)

func TestExecWithSum(t *testing.T) {
	input := []int{1, 0, 0, 0, 99}
	output := []int{2, 0, 0, 0, 99}
	result := Exec(input)
	for i := 0; i < len(input); i++ {
		if result[i] != output[i] {
			t.Errorf("Given input %d, output should be %d, but output is %d", input[i], output[i], result[i])
		}
	}
}

func TestExecWithMultiplication(t *testing.T) {
	input := []int{2, 3, 0, 3, 99}
	output := []int{2, 3, 0, 6, 99}
	result := Exec(input)
	for i := 0; i < len(input); i++ {
		if result[i] != output[i] {
			t.Errorf("Given input %d, output should be %d, but output is %d", input[i], output[i], result[i])
		}
	}
}

func TestExecWithMixedOperations(t *testing.T) {
	input := []int{1, 1, 1, 4, 99, 5, 6, 0, 99}
	output := []int{30, 1, 1, 4, 2, 5, 6, 0, 99}
	result := Exec(input)
	for i := 0; i < len(input); i++ {
		if result[i] != output[i] {
			t.Errorf("Given input %d at index %d, output should be %d, but output is %d", input[i], i, output[i], result[i])
		}
	}
}

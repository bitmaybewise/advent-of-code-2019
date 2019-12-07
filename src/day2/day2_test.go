package main

import (
	"io/ioutil"
	"strconv"
	"strings"
	"testing"
)

func TestExecAllWithSum(t *testing.T) {
	input := []int{1, 0, 0, 0, 99}
	output := []int{2, 0, 0, 0, 99}
	result := ExecAll(input)
	for i := 0; i < len(input); i++ {
		if result[i] != output[i] {
			t.Errorf("Given input %d, output should be %d, but output is %d", input[i], output[i], result[i])
		}
	}
}

func TestExecAllWithMultiplication(t *testing.T) {
	input := []int{2, 3, 0, 3, 99}
	output := []int{2, 3, 0, 6, 99}
	result := ExecAll(input)
	for i := 0; i < len(input); i++ {
		if result[i] != output[i] {
			t.Errorf("Given input %d, output should be %d, but output is %d", input[i], output[i], result[i])
		}
	}
}

func TestExecAllWithMixedOperations(t *testing.T) {
	input := []int{1, 1, 1, 4, 99, 5, 6, 0, 99}
	output := []int{30, 1, 1, 4, 2, 5, 6, 0, 99}
	result := ExecAll(input)
	for i := 0; i < len(input); i++ {
		if result[i] != output[i] {
			t.Errorf("Given input %d at index %d, output should be %d, but output is %d", input[i], i, output[i], result[i])
		}
	}
}

func TestOutput(t *testing.T) {
	if result := Output(12, 2); result != 1202 {
		t.Errorf("Given noun %d and verb %d, output should be %d, but result is %d", 12, 2, 1202, result)
	}
}

func TestFindInputsOf(t *testing.T) {
	input, _ := ioutil.ReadFile("input")
	lines := strings.Split(string(input), ",")
	var instructions []int
	for _, line := range lines {
		number, _ := strconv.Atoi(line)
		instructions = append(instructions, number)
	}
	expectedNoun := 24
	expectedVerb := 42
	expectedResult := 5141750
	noun, verb := FindInputsOf(expectedResult, instructions)
	if noun != expectedNoun || verb != expectedVerb {
		t.Errorf("Wrong noun %d and verb %d", noun, verb)
	}
}

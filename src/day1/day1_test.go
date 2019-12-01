package main

import "testing"

func TestMassOf(t *testing.T) {
	inputs := []int{12, 14, 1969, 100756}
	outputs := []int{2, 2, 654, 33583}
	for i := 0; i < len(inputs); i++ {
		result := MassOf(inputs[i])
		if result != outputs[i] {
			t.Errorf("Given input %d, output should be %d, but output is %d", inputs[i], outputs[i], result)
		}
	}
}

func TestSumMassOf(t *testing.T) {
	inputs := []int{12, 14, 1969, 100756}
	outputs := []int{2, 2, 654, 33583}
	result := 0
	for _, n := range outputs {
		result += n
	}
	sum := SumMassOf(inputs)
	if result != sum {
		t.Errorf("Given input %v, output should be %d, but output is %d", inputs, result, sum)
	}
}

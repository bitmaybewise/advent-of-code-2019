package main

import (
	"testing"
)

func TestFuelRequirement(t *testing.T) {
	inputs := []int{12, 14, 1969, 100756}
	outputs := []int{2, 2, 654, 33583}
	for i := 0; i < len(inputs); i++ {
		result := FuelRequirement(inputs[i])
		if result != outputs[i] {
			t.Errorf("Given input %d, output should be %d, but output is %d", inputs[i], outputs[i], result)
		}
	}
}

func TestFuelSum(t *testing.T) {
	inputs := []int{12, 14, 1969, 100756}
	outputs := []int{2, 2, 654, 33583}
	result := 0
	for _, n := range outputs {
		result += n
	}
	sum := FuelSum(inputs)
	if result != sum {
		t.Errorf("Given input %v, output should be %d, but output is %d", inputs, result, sum)
	}
}

func TestAdditionalFuelRequirement(t *testing.T) {
	inputs := []int{12, 1969, 100756}
	outputs := []int{2, 966, 50346}
	for i := 0; i < len(inputs); i++ {
		result := AdditionalFuelRequirement(inputs[i])
		if result != outputs[i] {
			t.Errorf("Given input %d, output should be %d, but output is %d", inputs[i], outputs[i], result)
		}
	}
}

func TestAdditionalFuelSum(t *testing.T) {
	inputs := []int{12, 1969, 100756}
	outputs := []int{2, 966, 50346}
	result := 0
	for _, n := range outputs {
		result += n
	}
	sum := AdditionalFuelSum(inputs)
	if result != sum {
		t.Errorf("Given input %v, output should be %d, but output is %d", inputs, result, sum)
	}
}

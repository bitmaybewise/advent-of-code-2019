package main

import (
	"testing"
)

func TestIntToOpcodeLowerOpcodes(t *testing.T) {
	opcodes := []int{1, 2, 3, 4, 5}
	for _, op := range opcodes {
		value := IntToOpcode(op)
		opcode := Opcode{Operation(op), 0, 0, 0}
		if value != opcode {
			t.Errorf("Wrong opcode %v, should be %v", value, opcode)
		}
	}
}

func TestIntToOpcodeWithBigOpcodes(t *testing.T) {
	value := IntToOpcode(1002)
	opcode := Opcode{Operation(2), 0, 1, 0}
	if value != opcode {
		t.Errorf("Wrong opcode %v, should be %v", value, opcode)
	}
	value = IntToOpcode(11101)
	opcode = Opcode{Operation(1), 1, 1, 1}
	if value != opcode {
		t.Errorf("Wrong opcode %v, should be %v", value, opcode)
	}
}

func TestExecAllWithSum(t *testing.T) {
	input := []int{1001, 4, 3, 4, 33, 99}
	output := []int{1001, 4, 3, 4, 36, 99}
	result, _ := ExecAll(input, 1)
	for i := 0; i < len(input); i++ {
		if result[i] != output[i] {
			t.Errorf("Given input %d, output should be %d, but output is %d", input[i], output[i], result[i])
		}
	}
}

func TestExecAllWithMultiplication(t *testing.T) {
	input := []int{1002, 4, 3, 4, 33}
	output := []int{1002, 4, 3, 4, 99}
	result, _ := ExecAll(input, 1)
	for i := 0; i < len(input); i++ {
		if result[i] != output[i] {
			t.Errorf("Given input %d, output should be %d, but output is %d", input[i], output[i], result[i])
		}
	}
}

func TestExecAllWithSave(t *testing.T) {
	input := []int{3, 0, 99}
	output := []int{1, 0, 99}
	result, _ := ExecAll(input, 1)
	for i := 0; i < len(input); i++ {
		if result[i] != output[i] {
			t.Errorf("Given input %d, output should be %d, but output is %d", input[i], output[i], result[i])
		}
	}
}

func TestExecAllWithOutput(t *testing.T) {
	input := []int{4, 0, 99}
	output := []int{4, 0, 99}
	result, v := ExecAll(input, 1)
	if v != 4 {
		t.Errorf("Output should be %d, but is %d", 4, v)
	}
	for i := 0; i < len(input); i++ {
		if result[i] != output[i] {
			t.Errorf("Given input %d, output should be %d, but output is %d", input[i], output[i], result[i])
		}
	}
	input = []int{3, 0, 4, 0, 99}
	output = []int{1, 0, 4, 0, 99}
	result, v = ExecAll(input, 1)
	if v != 1 {
		t.Errorf("Output should be %d, but is %d", 1, v)
	}
	for i := 0; i < len(input); i++ {
		if result[i] != output[i] {
			t.Errorf("Given input %d, output should be %d, but output is %d", input[i], output[i], result[i])
		}
	}
}

func TestExecAllWithNegatives(t *testing.T) {
	input := []int{1101, 100, -1, 4, 0}
	output := []int{1101, 100, -1, 4, 99}
	result, _ := ExecAll(input, 1)
	for i := 0; i < len(input); i++ {
		if result[i] != output[i] {
			t.Errorf("Given input %d, output should be %d, but output is %d", input[i], output[i], result[i])
		}
	}
}

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

func TestExecAllWithEquals(t *testing.T) {
	input := []int{3, 9, 8, 9, 10, 9, 4, 9, 99, -1, 0}
	output := []int{3, 9, 8, 9, 10, 9, 4, 9, 99, 1, 0}
	result, _ := ExecAll(input, 0)
	for i := 0; i < len(input); i++ {
		if result[i] != output[i] {
			t.Errorf("Given input %d, output should be %d, but output is %d", input[i], output[i], result[i])
		}
	}
	input = []int{3, 3, 1108, -1, 0, 3, 4, 3, 99}
	output = []int{3, 3, 1108, 1, 0, 3, 4, 3, 99}
	result, _ = ExecAll(input, 0)
	for i := 0; i < len(input); i++ {
		if result[i] != output[i] {
			t.Errorf("Given input %d, output should be %d, but output is %d", input[i], output[i], result[i])
		}
	}
}

func TestExecAllWithLessThan(t *testing.T) {
	input := []int{3, 9, 7, 9, 10, 9, 4, 9, 99, -1, 8}
	output := []int{3, 9, 7, 9, 10, 9, 4, 9, 99, 1, 8}
	result, _ := ExecAll(input, 0)
	for i := 0; i < len(input); i++ {
		if result[i] != output[i] {
			t.Errorf("Given input %d, output should be %d, but output is %d", input[i], output[i], result[i])
		}
	}
	input = []int{3, 3, 1107, -1, 8, 3, 4, 3, 99}
	output = []int{3, 3, 1107, 1, 8, 3, 4, 3, 99}
	result, _ = ExecAll(input, 0)
	for i := 0; i < len(input); i++ {
		if result[i] != output[i] {
			t.Errorf("Given input %d, output should be %d, but output is %d", input[i], output[i], result[i])
		}
	}
}

func TestExecAllWithJumpIfTrue(t *testing.T) {
	input := []int{3, 3, 1105, -1, 9, 1101, 0, 0, 12, 4, 12, 99, 1}
	_, output := ExecAll(input, 1)
	if output != 1 {
		t.Errorf("Wrong output %d", output)
	}
	_, output = ExecAll(input, 0)
	if output == 1 {
		t.Errorf("Wrong output %d", output)
	}
}

func TestExecAllWithJumpIfFalse(t *testing.T) {
	input := []int{3, 12, 6, 12, 15, 1, 13, 14, 13, 4, 13, 99, -1, 0, 1, 9}
	_, output := ExecAll(input, 1)
	if output != 1 {
		t.Errorf("Wrong output %d", output)
	}
	input = []int{3, 1, 6, 1, 9, 4, 8, 99, 1, 5}
	_, output = ExecAll(input, 0)
	if output != 1 {
		t.Errorf("Wrong output %d", output)
	}
}

func TestExecAllWithMixedInstructions(t *testing.T) {
	input := []int{3, 21, 1008, 21, 8, 20, 1005, 20, 22, 107, 8, 21, 20, 1006, 20, 31,
		1106, 0, 36, 98, 0, 0, 1002, 21, 125, 20, 4, 20, 1105, 1, 46, 104,
		999, 1105, 1, 46, 1101, 1000, 1, 20, 4, 20, 1105, 1, 46, 98, 99}
	_, output := ExecAll(input, 8)
	if output != 1000 {
		t.Errorf("Wrong output %d", output)
	}
	_, output = ExecAll(input, 9)
	if output != 1001 {
		t.Errorf("Wrong output %d", output)
	}
	_, output = ExecAll(input, 7)
	if output != 999 {
		t.Errorf("Wrong output %d", output)
	}
}

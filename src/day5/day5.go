/*
de, which also works like it did before - 99 is written to address 4.

Parameters that an instruction writes to will never be in immediate mode.

Finally, some notes:

    It is important to remember that the instruction pointer should increase by the number of values in the instruction after the instruction finishes. Because of the new instructions, this amount is no longer always 4.
    Integers can be negative: 1101,100,-1,4,0 is a valid program (find 100 + -1, store the result in position 4).

The TEST diagnostic program will start by requesting from the user the ID of the system to test by running an input instruction - provide it 1, the ID for the ship's air conditioner unit.

It will then perform a series of diagnostic tests confirming that various parts of the Intcode computer, like parameter modes, function correctly. For each test, it will run an output instruction indicating how far the result of the test was from the expected value, where 0 means the test was successful. Non-zero outputs mean that a function is not working correctly; check the instructions that were run before the output instruction to see which one failed.

Finally, the program will output a diagnostic code and immediately halt. This final output isn't an error; an output followed immediately by a halt means the program finished. If all outputs were zero except the diagnostic code, the diagnostic program ran successfully.

After providing 1 to the only input instruction and passing all the tests, what diagnostic code does the program produce?

Your puzzle answer was 9938601.

The first half of this puzzle is complete! It provides one gold star: *
--- Part Two ---

The air conditioner comes online! Its cold air feels good for a while, but then the TEST alarms start to go off. Since the air conditioner can't vent its heat anywhere but back into the spacecraft, it's actually making the air inside the ship warmer.

Instead, you'll need to use the TEST to extend the thermal radiators. Fortunately, the diagnostic program (your puzzle input) is already equipped for this. Unfortunately, your Intcode computer is not.

Your computer is only missing a few opcodes:

    Opcode 5 is jump-if-true: if the first parameter is non-zero, it sets the instruction pointer to the value from the second parameter. Otherwise, it does nothing.
    Opcode 6 is jump-if-false: if the first parameter is zero, it sets the instruction pointer to the value from the second parameter. Otherwise, it does nothing.
    Opcode 7 is less than: if the first parameter is less than the second parameter, it stores 1 in the position given by the third parameter. Otherwise, it stores 0.
    Opcode 8 is equals: if the first parameter is equal to the second parameter, it stores 1 in the position given by the third parameter. Otherwise, it stores 0.

Like all instructions, these instructions need to support parameter modes as described above.

Normally, after an instruction is finished, the instruction pointer increases by the number of values in that instruction. However, if the instruction modifies the instruction pointer, that value is used and the instruction pointer is not automatically increased.

For example, here are several programs that take one input, compare it to the value 8, and then produce one output:

    3,9,8,9,10,9,4,9,99,-1,8 - Using position mode, consider whether the input is equal to 8; output 1 (if it is) or 0 (if it is not).
    3,9,7,9,10,9,4,9,99,-1,8 - Using position mode, consider whether the input is less than 8; output 1 (if it is) or 0 (if it is not).
    3,3,1108,-1,8,3,4,3,99 - Using immediate mode, consider whether the input is equal to 8; output 1 (if it is) or 0 (if it is not).
    3,3,1107,-1,8,3,4,3,99 - Using immediate mode, consider whether the input is less than 8; output 1 (if it is) or 0 (if it is not).

Here are some jump tests that take an input, then output 0 if the input was zero or 1 if the input was non-zero:

    3,12,6,12,15,1,13,14,13,4,13,99,-1,0,1,9 (using position mode)
    3,3,1105,-1,9,1101,0,0,12,4,12,99,1 (using immediate mode)

Here's a larger example:

3,21,1008,21,8,20,1005,20,22,107,8,21,20,1006,20,31,
1106,0,36,98,0,0,1002,21,125,20,4,20,1105,1,46,104,
999,1105,1,46,1101,1000,1,20,4,20,1105,1,46,98,99

The above example program uses an input instruction to ask for a single number. The program will then output 999 if the input value is below 8, output 1000 if the input value is equal to 8, or output 1001 if the input value is greater than 8.

This time, when the TEST diagnostic program runs its input instruction to get the ID of the system to test, provide it 5, the ID for the ship's thermal radiator controller. This diagnostic test suite only outputs one number, the diagnostic code.

What is the diagnostic code for system ID 5?

*/

package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type Operation int
type ParamMode int

type Opcode struct {
	operation           Operation
	mode1, mode2, mode3 ParamMode
}

const opSum = Operation(1)
const opMultiply = Operation(2)
const opSave = Operation(3)
const opOutput = Operation(4)

const positionMode = ParamMode(0)
const immediateMode = ParamMode(1)

func main() {
	input, _ := ioutil.ReadFile("input")
	lines := strings.Split(string(input), ",")
	var instructions []int
	for _, line := range lines {
		number, _ := strconv.Atoi(line)
		instructions = append(instructions, number)
	}
	_, output := ExecAll(instructions, 1)
	fmt.Printf("Answer 1: %d\n", output)
}

func IntToOpcode(op int) Opcode {
	s := fmt.Sprintf("%05s", strconv.Itoa(op))
	operation, _ := strconv.Atoi(s[3:])
	param1, _ := strconv.Atoi(s[2:3])
	param2, _ := strconv.Atoi(s[1:2])
	param3, _ := strconv.Atoi(s[0:1])
	return Opcode{Operation(operation), ParamMode(param1), ParamMode(param2), ParamMode(param3)}
}

func ExecAll(instructions []int, inputTest int) ([]int, int) {
	ins := make([]int, len(instructions))
	copy(ins, instructions)
	return execOpcode(0, ins, inputTest, 0)
}

func getParam(instructions []int, pos int, mode ParamMode) int {
	if mode == immediateMode {
		return pos
	}
	return instructions[pos]
}

func execOpcode(pos int, instructions []int, inputTest int, output int) ([]int, int) {
	opcode := IntToOpcode(instructions[pos])
	switch opcode.operation {
	case opSum:
		idx1 := getParam(instructions, pos+1, opcode.mode1)
		idx2 := getParam(instructions, pos+2, opcode.mode2)
		idx3 := getParam(instructions, pos+3, positionMode)
		instructions[idx3] = instructions[idx1] + instructions[idx2]
		return execOpcode(pos+4, instructions[:len(instructions)], inputTest, output)
	case opMultiply:
		idx1 := getParam(instructions, pos+1, opcode.mode1)
		idx2 := getParam(instructions, pos+2, opcode.mode2)
		idx3 := getParam(instructions, pos+3, positionMode)
		instructions[idx3] = instructions[idx1] * instructions[idx2]
		return execOpcode(pos+4, instructions[:len(instructions)], inputTest, output)
	case opSave:
		idx1 := getParam(instructions, pos+1, positionMode)
		instructions[idx1] = inputTest
		return execOpcode(pos+2, instructions[:len(instructions)], inputTest, output)
	case opOutput:
		idx1 := getParam(instructions, pos+1, opcode.mode1)
		return execOpcode(pos+2, instructions[:len(instructions)], inputTest, instructions[idx1])
	default:
		return instructions, output
	}
}

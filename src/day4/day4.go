/*
--- Day 4: Secure Container ---

You arrive at the Venus fuel depot only to discover it's protected by a password. The Elves had written the password on a sticky note, but someone threw it out.

However, they do remember a few key facts about the password:

    It is a six-digit number.
    The value is within the range given in your puzzle input.
    Two adjacent digits are the same (like 22 in 122345).
    Going from left to right, the digits never decrease; they only ever increase or stay the same (like 111123 or 135679).

Other than the range rule, the following are true:

    111111 meets these criteria (double 11, never decreases).
    223450 does not meet these criteria (decreasing pair of digits 50).
    123789 does not meet these criteria (no double).

How many different passwords within the range given in your puzzle input meet these criteria?

Your puzzle input is 357253-892942.

Your puzzle answer was 530.

The first half of this puzzle is complete! It provides one gold star: *
--- Part Two ---

An Elf just remembered one more important detail: the two adjacent matching digits are not part of a larger group of matching digits.

Given this additional criterion, but still ignoring the range rule, the following are now true:

    112233 meets these criteria because the digits never decrease and all repeated digits are exactly two digits long.
    123444 no longer meets the criteria (the repeated 44 is part of a larger group of 444).
    111122 meets the criteria (even though 1 is repeated more than twice, it still contains a double 22).

How many different passwords within the range given in your puzzle input meet all of the criteria?

Your puzzle input is still 357253-892942.

Your puzzle answer was 324.

That's the right answer! You are one gold star closer to rescuing Santa.
*/

package main

import (
	"fmt"
	"strconv"
	"strings"
)

const input = "357253-892942"

func main() {
	fmt.Printf("day 4 puzzle 1: %d\n", CountDifferentPasswords(input, MeetCriteria))
	fmt.Printf("day 4 puzzle 2: %d\n", CountDifferentPasswords(input, MeetCriteriaNoAdjacentGroups))
}

func MeetCriteria(digits []int) bool {
	var hasDouble bool
	for i := 0; i < len(digits)-1; i++ {
		if digits[i] > digits[i+1] {
			return false
		}
		if hasDouble {
			continue
		}
		if digits[i] == digits[i+1] {
			hasDouble = true
		}
	}
	return hasDouble
}

func MeetCriteriaNoAdjacentGroups(digits []int) bool {
	counter := make(map[int]int)
	size := len(digits) - 1
	for i := 0; i < size; i++ {
		if digits[i] > digits[i+1] {
			return false
		}
		counter[digits[i]]++
	}
	counter[digits[size]]++
	for _, v := range counter {
		if v == 2 {
			return true
		}
	}
	return false
}

func intToArray(i int) []int {
	var r []int
	s := strconv.Itoa(i)
	for _, v := range s {
		n, _ := strconv.Atoi(string(v))
		r = append(r, n)
	}
	return r
}

func CountDifferentPasswords(givenRange string, criteria func([]int) bool) int {
	values := strings.Split(givenRange, "-")
	start, _ := strconv.Atoi(values[0])
	finish, _ := strconv.Atoi(values[1])
	var count int
	for start <= finish {
		digits := intToArray(start)
		if criteria(digits) {
			count++
		}
		start++
	}
	return count
}

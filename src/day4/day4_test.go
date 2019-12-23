package main

import (
	"testing"
)

func TestMeetCriteria(t *testing.T) {
	digits := []int{1, 1, 1, 1, 1, 1}
	if !MeetCriteria(digits) {
		t.Errorf("Should be valid digits %v", digits)
	}
	digits = []int{2, 2, 3, 4, 5, 0}
	if MeetCriteria(digits) {
		t.Errorf("Should be invalid %v, they never decrease", digits)
	}
	digits = []int{1, 2, 3, 7, 8, 9}
	if MeetCriteria(digits) {
		t.Errorf("Should be invalid %v, no double found", digits)
	}
}

func TestMeetCriteriaNoAdjacentGroups(t *testing.T) {
	digits := []int{1, 1, 2, 2, 3, 3}
	if !MeetCriteriaNoAdjacentGroups(digits) {
		t.Errorf("Should be valid digits %v", digits)
	}
	digits = []int{1, 2, 3, 4, 4, 4}
	if MeetCriteriaNoAdjacentGroups(digits) {
		t.Errorf("Should be invalid digits %v, repeated 44 is part of a larger group", digits)
	}
	digits = []int{1, 1, 1, 1, 2, 2}
	if !MeetCriteriaNoAdjacentGroups(digits) {
		t.Errorf("Should be valid digits %v", digits)
	}
}

func TestCountDifferentPasswords(t *testing.T) {
	input := "111111-111116"
	if count := CountDifferentPasswords(input, MeetCriteria); count != 6 {
		t.Errorf("Input %s should have %d valid passwords, but returned %d", input, 6, count)
	}
	input = "111111-111120"
	if count := CountDifferentPasswords(input, MeetCriteria); count != 9 {
		t.Errorf("Input %s should have %d valid passwords, but returned %d", input, 9, count)
	}
}

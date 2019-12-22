package main

import (
	"testing"
)

func TestManhattanDistance(t *testing.T) {
	p1 := Coord{6, 6}
	p2 := Coord{3, 3}
	if dist := ManhattanDistance(p1, p2); dist != 6 {
		t.Errorf("Wrong distance %d", dist)
	}
}

func TestParseMoveUp(t *testing.T) {
	p := ParseMove("U10")
	if p.x != 0 || p.y != 10 {
		t.Errorf("Wrong cartesian point (%d,%d)", p.x, p.y)
	}
}

func TestParseMoveDown(t *testing.T) {
	p := ParseMove("D10")
	if p.x != 0 || p.y != -10 {
		t.Errorf("Wrong cartesian point (%d,%d)", p.x, p.y)
	}
}

func TestParseMoveLeft(t *testing.T) {
	p := ParseMove("L10")
	if p.x != -10 || p.y != 0 {
		t.Errorf("Wrong cartesian point (%d,%d)", p.x, p.y)
	}
}

func TestParseMoveRight(t *testing.T) {
	p := ParseMove("R10")
	if p.x != 10 || p.y != 0 {
		t.Errorf("Wrong cartesian point (%d,%d)", p.x, p.y)
	}
}

func TestClosestIntersectionDistance(t *testing.T) {
	line1 := "R8,U5,L5,D3"
	line2 := "U7,R6,D4,L4"
	dist := ClosestIntersectionDistance(line1, line2)
	if dist != 6 {
		t.Errorf("Wrong distance %d, expected %d", dist, 6)
	}
	line1 = "R75,D30,R83,U83,L12,D49,R71,U7,L72"
	line2 = "U62,R66,U55,R34,D71,R55,D58,R83"
	dist = ClosestIntersectionDistance(line1, line2)
	if dist != 159 {
		t.Errorf("Wrong distance %d, expected %d", dist, 159)
	}
	line1 = "R98,U47,R26,D63,R33,U87,L62,D20,R33,U53,R51"
	line2 = "U98,R91,D20,R16,D67,R40,U7,R15,U6,R7"
	dist = ClosestIntersectionDistance(line1, line2)
	if dist != 135 {
		t.Errorf("Wrong distance %d, expected %d", dist, 135)
	}
}

/*
--- Day 3: Crossed Wires ---

The gravity assist was successful, and you're well on your way to the Venus refuelling station. During the rush back on Earth, the fuel management system wasn't completely installed, so that's next on the priority list.

Opening the front panel reveals a jumble of wires. Specifically, two wires are connected to a central port and extend outward on a grid. You trace the path each wire takes as it leaves the central port, one wire per line of text (your puzzle input).

The wires twist and turn, but the two wires occasionally cross paths. To fix the circuit, you need to find the intersection point closest to the central port. Because the wires are on a grid, use the Manhattan distance for this measurement. While the wires do technically cross right at the central port where they both start, this point does not count, nor does a wire count as crossing with itself.

For example, if the first wire's path is R8,U5,L5,D3, then starting from the central port (o), it goes right 8, up 5, left 5, and finally down 3:

...........
...........
...........
....+----+.
....|....|.
....|....|.
....|....|.
.........|.
.o-------+.
...........

Then, if the second wire's path is U7,R6,D4,L4, it goes up 7, right 6, down 4, and left 4:

...........
.+-----+...
.|.....|...
.|..+--X-+.
.|..|..|.|.
.|.-X--+.|.
.|..|....|.
.|.......|.
.o-------+.
...........

These wires cross at two locations (marked X), but the lower-left one is closer to the central port: its distance is 3 + 3 = 6.

Here are a few more examples:

    R75,D30,R83,U83,L12,D49,R71,U7,L72
    U62,R66,U55,R34,D71,R55,D58,R83 = distance 159
    R98,U47,R26,D63,R33,U87,L62,D20,R33,U53,R51
    U98,R91,D20,R16,D67,R40,U7,R15,U6,R7 = distance 135

What is the Manhattan distance from the central port to the closest intersection?

Your puzzle answer was 896.

The first half of this puzzle is complete! It provides one gold star: *
--- Part Two ---

It turns out that this circuit is very timing-sensitive; you actually need to minimize the signal delay.

To do this, calculate the number of steps each wire takes to reach each intersection; choose the intersection where the sum of both wires' steps is lowest. If a wire visits a position on the grid multiple times, use the steps value from the first time it visits that position when calculating the total value of a specific intersection.

The number of steps a wire takes is the total number of grid squares the wire has entered to get to that location, including the intersection being considered. Again consider the example from above:

...........
.+-----+...
.|.....|...
.|..+--X-+.
.|..|..|.|.
.|.-X--+.|.
.|..|....|.
.|.......|.
.o-------+.
...........

In the above example, the intersection closest to the central port is reached after 8+5+5+2 = 20 steps by the first wire and 7+6+4+3 = 20 steps by the second wire for a total of 20+20 = 40 steps.

However, the top-right intersection is better: the first wire takes only 8+5+2 = 15 and the second wire takes only 7+6+2 = 15, a total of 15+15 = 30 steps.

Here are the best steps for the extra examples from above:

    R75,D30,R83,U83,L12,D49,R71,U7,L72
    U62,R66,U55,R34,D71,R55,D58,R83 = 610 steps
    R98,U47,R26,D63,R33,U87,L62,D20,R33,U53,R51
    U98,R91,D20,R16,D67,R40,U7,R15,U6,R7 = 410 steps

What is the fewest combined steps the wires must take to reach an intersection?

Your puzzle answer was 16524.

That's the right answer! You are one gold star closer to rescuing Santa.
*/

package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"strconv"
	"strings"
)

type Coord struct {
	x, y, steps    int
	isIntersection bool
}

func (coord Coord) String() string {
	return fmt.Sprintf("%d,%d", coord.x, coord.y)
}

func (coord Coord) Eq(o *Coord) bool {
	return coord.x == o.x && coord.y == o.y
}

func initCoord() Coord {
	return Coord{0, 0, 0, false}
}

func strToIntersection(s string) Coord {
	var coord Coord
	v := strings.Split(s, ",")
	coord.x, _ = strconv.Atoi(v[0])
	coord.y, _ = strconv.Atoi(v[1])
	coord.isIntersection = true
	return coord
}

func main() {
	input, _ := ioutil.ReadFile("input")
	lines := strings.Split(string(input), "\n")
	distance := ClosestIntersectionDistance(lines[0], lines[1])
	bestStep := BestSteps(lines[0], lines[1])
	fmt.Printf("day 3 puzzle 1: %d\n", distance)
	fmt.Printf("day 3 puzzle 2: %d\n", bestStep)
}

func ManhattanDistance(p1 Coord, p2 Coord) int {
	x := float64(p1.x - p2.x)
	y := float64(p1.y - p2.y)
	return int(math.Abs(x) + math.Abs(y))
}

func ParseMove(mov string) *Coord {
	var x, y int
	direction := string(mov[0])
	value, _ := strconv.Atoi(mov[1:])
	switch {
	case direction == "U":
		y = value
	case direction == "D":
		y -= value
	case direction == "R":
		x = value
	case direction == "L":
		x -= value
	}
	c := Coord{x, y, 0, false}
	return &c
}

func ParseMovements(movements string) []Coord {
	var coords []Coord
	for _, move := range strings.Split(movements, ",") {
		coords = append(coords, *ParseMove(move))
	}
	return coords
}

type Visits = map[string]bool

func makeVisits(movements []Coord) Visits {
	vs := make(Visits)
	ignoreIntersection := initCoord()
	current := initCoord()
	for _, move := range movements {
		visit(&move, &current, vs, &ignoreIntersection)
	}
	return vs
}

func intersections(visits1 Visits, visits2 Visits) []Coord {
	var intersections []Coord
	for k := range visits1 {
		if visits2[k] {
			coord := strToIntersection(k)
			intersections = append(intersections, coord)
		}
	}
	return intersections
}

type Stop = bool

func visit(moveP *Coord, coord *Coord, visits Visits, intersection *Coord) Stop {
	i := moveP.x
	for i != 0 {
		if moveP.x > 0 {
			i--
			coord.x++
		} else {
			i++
			coord.x--
		}
		visits[coord.String()] = true
		coord.steps++
		if intersection.isIntersection && coord.Eq(intersection) {
			return true
		}
	}
	i = moveP.y
	for i != 0 {
		if moveP.y > 0 {
			i--
			coord.y++
		} else {
			i++
			coord.y--
		}
		visits[coord.String()] = true
		coord.steps++
		if intersection.isIntersection && coord.Eq(intersection) {
			return true
		}
	}
	return false
}

func ClosestIntersectionDistance(movements1 string, movements2 string) int {
	movs1 := ParseMovements(movements1)
	movs2 := ParseMovements(movements2)
	visits1, visits2 := makeVisits(movs1), makeVisits(movs2)
	intersections := intersections(visits1, visits2)
	closestDistance := math.MaxUint16
	for _, coord := range intersections {
		distance := ManhattanDistance(coord, initCoord())
		if closestDistance > distance {
			closestDistance = distance
		}
	}
	return closestDistance
}

func countSteps(movements []Coord, visits Visits, intersection *Coord) int {
	current := initCoord()
	for _, move := range movements {
		stop := visit(&move, &current, visits, intersection)
		if stop {
			break
		}
	}
	return current.steps
}

func BestSteps(movements1 string, movements2 string) int {
	movs1 := ParseMovements(movements1)
	movs2 := ParseMovements(movements2)
	visits1, visits2 := makeVisits(movs1), makeVisits(movs2)
	intersections := intersections(visits1, visits2)
	best := math.MaxInt16
	for _, coord := range intersections {
		steps1 := countSteps(movs1, visits1, &coord)
		steps2 := countSteps(movs2, visits2, &coord)
		steps := steps1 + steps2
		if best > steps {
			best = steps
		}
	}
	return best
}

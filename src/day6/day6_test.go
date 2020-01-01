package main

import (
	"testing"
)

func TestCountOrbits1DirectOrbit(t *testing.T) {
	orbits := []string{"COM)B"}
	value := CountOrbits(orbits)
	if value != 1 {
		t.Errorf("Wrong result %d, should be %d", value, 1)
	}
}

func TestCountOrbits2DirectOrbits(t *testing.T) {
	orbits := []string{"COM)B", "COM)C"}
	value := CountOrbits(orbits)
	if value != 2 {
		t.Errorf("Wrong result %d, should be %d", value, 2)
	}
	// inverting orbits
	orbits = []string{"COM)C", "COM)B"}
	value = CountOrbits(orbits)
	if value != 2 {
		t.Errorf("Wrong result %d, should be %d", value, 2)
	}
}

func TestCountOrbits1IndirectOrbit(t *testing.T) {
	orbits := []string{"COM)B", "B)C"}
	value := CountOrbits(orbits)
	if value != 3 {
		t.Errorf("Wrong result %d, should be %d", value, 3)
	}
	// inverting orbits
	orbits = []string{"B)C", "COM)B"}
	value = CountOrbits(orbits)
	if value != 3 {
		t.Errorf("Wrong result %d, should be %d", value, 3)
	}
}

func TestCountOrbits2IndirectOrbits(t *testing.T) {
	orbits := []string{"COM)B", "B)C", "C)D"}
	value := CountOrbits(orbits)
	if value != 6 {
		t.Errorf("Wrong result %d, should be %d", value, 6)
	}
}

func TestCountOrbits2DirectAndIndirectOrbits(t *testing.T) {
	orbits := []string{"COM)B", "B)C", "COM)D", "D)E"}
	value := CountOrbits(orbits)
	if value != 6 {
		t.Errorf("Wrong result %d, should be %d", value, 6)
	}
}

func TestCountOrbitsMultipleConnections(t *testing.T) {
	orbits := []string{"COM)B", "B)C", "C)D", "D)E", "E)F", "B)G", "G)H", "D)I", "E)J", "J)K", "K)L"}
	value := CountOrbits(orbits)
	if value != 42 {
		t.Errorf("Wrong result %d, should be %d", value, 42)
	}
}

func TestCountOrbitsWithMultipleCOM(t *testing.T) {
	orbits := []string{"COM)B", "COM)C", "C)D"}
	value := CountOrbits(orbits)
	if value != 4 {
		t.Errorf("Wrong result %d, should be %d", value, 4)
	}
}

func TestCountOrbitsWithCOMAppendedLater(t *testing.T) {
	orbits := []string{"A)B", "B)C", "COM)A"}
	value := CountOrbits(orbits)
	if value != 6 {
		t.Errorf("Wrong result %d, should be %d", value, 3)
	}
	orbits = []string{"COM)A", "B)C", "A)B", "B)D"}
	value = CountOrbits(orbits)
	if value != 9 {
		t.Errorf("Wrong result %d, should be %d", value, 6)
	}
}

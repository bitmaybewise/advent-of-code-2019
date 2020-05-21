package main

import (
	"testing"
)

func TestCountOrbits(t *testing.T) {
	t.Skip()

	t.Run("1 Direct Orbit", func(t *testing.T) {
		orbits := []string{"COM)B"}
		value := CountOrbits(orbits)
		if value != 1 {
			t.Errorf("Wrong result %d, should be %d", value, 1)
		}
	})

	t.Run("2 Direct Orbits", func(t *testing.T) {
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
	})

	t.Run("1 Indirect Orbit", func(t *testing.T) {
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
	})

	t.Run("2 Indirect Orbits", func(t *testing.T) {
		orbits := []string{"COM)B", "B)C", "C)D"}
		value := CountOrbits(orbits)
		if value != 6 {
			t.Errorf("Wrong result %d, should be %d", value, 6)
		}
	})

	t.Run("2 Direct and Indirect Orbits", func(t *testing.T) {
		orbits := []string{"COM)B", "B)C", "COM)D", "D)E"}
		value := CountOrbits(orbits)
		if value != 6 {
			t.Errorf("Wrong result %d, should be %d", value, 6)
		}
	})

	t.Run("Multiple Connections", func(t *testing.T) {
		orbits := []string{"COM)B", "B)C", "C)D", "D)E", "E)F", "B)G", "G)H", "D)I", "E)J", "J)K", "K)L"}
		value := CountOrbits(orbits)
		if value != 42 {
			t.Errorf("Wrong result %d, should be %d", value, 42)
		}
	})

	t.Run("Multiple COM", func(t *testing.T) {
		orbits := []string{"COM)B", "COM)C", "C)D"}
		value := CountOrbits(orbits)
		if value != 4 {
			t.Errorf("Wrong result %d, should be %d", value, 4)
		}
	})

	t.Run("COM Appended Later", func(t *testing.T) {
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
	})
}

func TestCountOrbitalTransfers(t *testing.T) {
	t.Run("A Few Transfers", func(t *testing.T) {
		orbits := []string{"COM)A", "A)B", "B)D", "B)C", "D)E"}
		value := CountOrbitalTransfers(orbits, "E", "A")
		expected := 1
		if value != expected {
			t.Errorf("Wrong result %d, should be %d", value, expected)
		}
	})

	t.Run("You and Santa", func(t *testing.T) {
		orbits := []string{"COM)B", "B)C", "C)D", "D)E", "E)F", "B)G", "G)H", "D)I", "E)J", "J)K", "K)L", "K)YOU", "I)SAN"}
		value := CountOrbitalTransfers(orbits, "YOU", "SAN")
		expected := 4
		if value != expected {
			t.Errorf("Wrong result %d, should be %d", value, expected)
		}
	})
}

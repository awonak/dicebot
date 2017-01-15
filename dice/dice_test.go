package dice

import (
	"testing"
)

var dicePatternTests = []struct {
	pattern  string
	numDice  int
	sides    int
	expected int
}{
	{"1d6", 1, 6, 1},
	{"2d20", 2, 20, 30},
	{"4d10", 4, 10, 20},
}

func TestValidDiceRollPatterns(t *testing.T) {

	for _, tt := range dicePatternTests {

		d, _ := NewDiceRollWithSeed(tt.pattern, 0)

		// Test pattern parsed as expected
		if d.NumDice != tt.numDice {
			t.Errorf("NumDice value incorrect, expected: %v got: %v.", tt.numDice, d.NumDice)
		}
		if d.Sides != tt.sides {
			t.Errorf("Sides value incorrect, expected: %v got: %v.", tt.sides, d.Sides)
		}

		// Test the roll value is within expected range
		roll := d.Roll()

		if roll != tt.expected {
			t.Errorf("Roll %v result invalid. expected: %v, got: %v.", d, tt.expected, roll)
		}

		// Test the string representation of the DiceRoll
		if d.String() != tt.pattern {
			t.Errorf("Expected %v got %v", tt.pattern, d)
		}
	}
}

var invalidDicePatternTests = []struct {
	pattern  string
	expected error
}{
	{"", ErrInvalidPattern},
	{"zzz", ErrInvalidPattern},
	{"1X6", ErrInvalidPattern},
	{"2.2d3.3", ErrInvalidPattern},
	{"1d2", ErrInvalidSides},
	{"1d101", ErrInvalidSides},
	{"0d20", ErrInvalidNumDice},
	{"20d20", ErrInvalidNumDice},
}

func TestInvalidDiceRollPatterns(t *testing.T) {

	for _, tt := range invalidDicePatternTests {

		_, err := NewDiceRoll(tt.pattern)
		if err != tt.expected {
			t.Errorf("expected %v, got %v", tt.expected, err)
		}
	}
}

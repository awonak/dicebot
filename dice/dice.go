package dice

import (
	"errors"
	"fmt"
	"math/rand"
	"regexp"
	"strconv"
	"time"
)

var (
	// Dice config
	MIN_DICE  = 1
	MAX_DICE  = 6
	MIN_SIDES = 3
	MAX_SIDES = 100

	// Error types
	ErrInvalidPattern = errors.New("Invalid pattern")
	ErrInvalidNumDice = errors.New(fmt.Sprintf("Number of dice must be between %d and %d", MIN_DICE, MAX_DICE))
	ErrInvalidSides   = errors.New(fmt.Sprintf("Dice sides must be between %d and %d", MIN_SIDES, MAX_SIDES))
)

type DiceRoll struct {
	NumDice int
	Sides   int
	seed    int64
}

func NewDiceRoll(pattern string) (*DiceRoll, error) {
	return NewDiceRollWithSeed(pattern, time.Now().UnixNano())
}

func NewDiceRollWithSeed(pattern string, seed int64) (*DiceRoll, error) {
	// Parse Pattern regex
	re := regexp.MustCompile("^(\\d+)d(\\d+)$")
	result := re.FindStringSubmatch(pattern)

	if len(result) == 0 {
		return nil, ErrInvalidPattern
	}

	// Parse the number strings to ints
	numDice, err := strconv.Atoi(result[1])
	if err != nil {
		return nil, err
	}

	sides, err := strconv.Atoi(result[2])
	if err != nil {
		return nil, err
	}

	// Validation rules
	if numDice < MIN_DICE || numDice > MAX_DICE {
		return nil, ErrInvalidNumDice
	}
	if sides < MIN_SIDES || sides > MAX_SIDES {
		return nil, ErrInvalidSides
	}

	return &DiceRoll{numDice, sides, seed}, nil
}

func (d DiceRoll) String() string {
	return fmt.Sprintf("%vd%v", d.NumDice, d.Sides)
}

func (d DiceRoll) Roll() int {
	results := 0

	for i := 0; i < d.NumDice; i++ {
		rand.Seed(d.seed)
		v := rand.Intn(d.Sides) + 1
		results += v
	}

	return results
}

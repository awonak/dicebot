package dice

import (
	"fmt"
	"math/rand"
	"regexp"
	"strconv"
	"time"
)

type DiceRoll struct {
	NumDice int
	Sides   int
	roll    string
}

func NewDiceRoll(roll string) (*DiceRoll, error) {
	re := regexp.MustCompile("(\\d+)d(\\d+)")
	result := re.FindStringSubmatch(roll)

	numDice, err := strconv.Atoi(result[1])
	if err != nil {
		return nil, err
	}

	sides, err := strconv.Atoi(result[2])
	if err != nil {
		return nil, err
	}

	return &DiceRoll{numDice, sides, roll}, nil
}

func (d DiceRoll) String() string {
	return fmt.Sprintf("%v", d.roll)
}

func (d DiceRoll) Roll() int{
	results := 0

	for i := 0; i < d.NumDice; i++ {
		rand.Seed(time.Now().UnixNano())
		v := rand.Intn(d.Sides) + 1
		results += v
	}

	return results
}


package robotname

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

// Robot is a friendly automaton
type Robot struct {
	name string
}

// Max available names [A-Z][A-Z][0-999]
const nameLimit = 26 * 26 * 1000

// Keep all used names
var nameUsed = map[string]bool{}

var seededRand *rand.Rand

func init() {
	seededRand = rand.New(rand.NewSource(time.Now().UnixNano()))
}

// Generating Robot Names
func generateName() string {
	r1 := seededRand.Intn(26) + 'A'
	r2 := seededRand.Intn(26) + 'A'
	digits := seededRand.Intn(1000)

	return fmt.Sprintf("%c%c%03d", r1, r2, digits)
}

// Name returns this robot's name
func (r *Robot) Name() (string, error) {
	if r.name != "" {
		return r.name, nil
	}

	// In case all names are used
	if len(nameUsed) >= nameLimit {
		return "", errors.New("robots > 26 * 26 * 1000")
	}

	// Find a free name slot
	for r.name == "" || nameUsed[r.name] == true {
		r.name = generateName()
	}
	nameUsed[r.name] = true

	return r.name, nil
}

// Reset will reset the robot's name
func (r *Robot) Reset() error {
	r.name = ""
	return nil
}

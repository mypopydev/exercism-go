package thefarm

import (
	"errors"
	"fmt"
)	

// See types.go for the types defined for this exercise.

// SillyNephewError returns error.
type SillyNephewError struct {
	cows int
}

func (e *SillyNephewError) Error() string {
	return fmt.Sprintf("silly nephew, there cannot be %d cows", e.cows)
} 

// DivideFood computes the fodder amount per cow for the given cows.
func DivideFood(weightFodder WeightFodder, cows int) (float64, error) {
	if cows == 0 {
		return 0.0, errors.New("division by zero")
	}
	if cows < 0 {
		return 0.0, &SillyNephewError{cows}
	}

	amount, err := weightFodder.FodderAmount()
	if err != nil {
		if err == ErrScaleMalfunction {
			if amount < 0 {
				return 0.0, errors.New("negative fodder")
			} else {
				return (amount * 2.0) / float64(cows), nil
			}
		} else {
			return 0.0, err
		}
	}

	// Manage fodder amount negative
	if amount < 0 {
		return 0.0, errors.New("negative fodder")
	}

	return (amount / float64(cows)), nil
}

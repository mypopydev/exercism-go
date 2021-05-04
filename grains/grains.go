package grains

import "errors"

// Square get the 2^(input-1)
func Square(input int) (uint64, error) {
	if input < 1 || input > 64 {
		return 0, errors.New("input not in [1 - 64]")
	}

	return 1 << (input - 1), nil
}

// Total calculate the number of grains of wheat on a chessboard
// given that the number on each square doubles
func Total() uint64 {
	return 1<<64 - 1
}

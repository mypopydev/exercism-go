package collatzconjecture

import "errors"

// CollatzConjecture calc the 3x+1 problem, given a number i, return the
// number of steps required to reach 1
func CollatzConjecture(i int) (int, error) {
	if i <= 0 {
		return -1, errors.New("negative or zero value")
	}
	if i == 1 {
		return 0, nil
	}

	if i%2 == 0 {
		i = i / 2
	} else {
		i = 3*i + 1
	}

	val, err := CollatzConjecture(i)

	return 1 + val, err
}

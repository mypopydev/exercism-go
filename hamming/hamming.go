package hamming

import "errors"

// Distance calculate the Hamming Distance between two strings.
func Distance(a, b string) (int, error) {
	lenA := len(a)
	lenB := len(b)

	if lenA != lenB {
		return -1, errors.New("string is not equal length")
	}

	num := 0

	arrayA := []byte(a)
	arrayB := []byte(b)

	for i := 0; i < lenA; i++ {
		if arrayA[i] != arrayB[i] {
			num++
		}
	}

	return num, nil
}

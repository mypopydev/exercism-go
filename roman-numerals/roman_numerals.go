package romannumerals

import "errors"

// ToRomanNumera to convert from normal numbers to Roman Numerals.
// input arabic range (0, 3000]
func ToRomanNumeral(arabic int) (string, error) {
	if arabic > 3000 || arabic <= 0 {
		return "", errors.New("arabic number is out of boundaries")
	}

	var (
		ones      = []string{"", "I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX"}
		tens      = []string{"", "X", "XX", "XXX", "XL", "L", "LX", "LXX", "LXXX", "XC"}
		hundreds  = []string{"", "C", "CC", "CCC", "CD", "D", "DC", "DCC", "DCCC", "CM"}
		thousands = []string{"", "M", "MM", "MMM"}
	)

	return thousands[arabic%1e4/1e3] + hundreds[arabic%1e3/1e2] + tens[arabic%1e2/1e1] + ones[arabic%1e1/1e0], nil
}

package twelve

// Verse takes an int to identify a verse from the twelve days of
// christmas and returns it as a string
func Verse(input int) string {
	verses := []string{
		"",
		"and a Partridge in a Pear Tree.",
		"two Turtle Doves,",
		"three French Hens,",
		"four Calling Birds,",
		"five Gold Rings,",
		"six Geese-a-Laying,",
		"seven Swans-a-Swimming,",
		"eight Maids-a-Milking,",
		"nine Ladies Dancing,",
		"ten Lords-a-Leaping,",
		"eleven Pipers Piping,",
		"twelve Drummers Drumming,",}

	// relative to verses
	days := map[int]string{
		1: "first",
		2: "second",
		3: "third",
		4: "fourth",
		5: "fifth",
		6: "sixth",
		7: "seventh",
		8: "eighth",
		9: "ninth",
		10: "tenth",
		11: "eleventh",
		12: "twelfth"}

	outVerse := "On the " + days[input] +
		" day of Christmas my true love gave to me:"

	// remove "and"
	if input == 1 {
		outVerse += verses[input][3:]
		return outVerse
	}

	for i := input; i >= 2; i-- {
		outVerse += " " + verses[i]
	}
	outVerse += " " + verses[1]

	return outVerse
}

// Song returns the entire twelve days of christmas
func Song() string {
	var out string

	for i := 1; i <= 11; i++ {
		out += Verse(i) + "\n"
	}

	out += Verse(12)

	return out
}

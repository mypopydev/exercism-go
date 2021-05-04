package proverb

// Proverb will output the full text of this proverbial rhyme
func Proverb(rhyme []string) []string {
	if len(rhyme) == 0 {
		return nil
	}

	proverbs := make([]string, len(rhyme))

	for i := 0; i < len(rhyme)-1; i++ {
		proverbs[i] = "For want of a " + rhyme[i] + " the " + rhyme[i+1] + " was lost."
	}
	proverbs[len(rhyme)-1] = "And all for the want of a " + rhyme[0] + "."

	return proverbs
}

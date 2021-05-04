package letter

// FreqMap records the frequency of each rune in a given text.
type FreqMap map[rune]int

// Frequency counts the frequency of each rune in a given text and returns this
// data as a FreqMap.
func Frequency(s string) FreqMap {
	m := FreqMap{}
	for _, r := range s {
		m[r]++
	}
	return m
}

// ConcurrentFrequency parallel counts the frequency of each rune in a given text
// and returns this data as a FreqMap.
func ConcurrentFrequency(strings []string) FreqMap {
	m := FreqMap{}
	count := len(strings)
	results := make(chan FreqMap, count)

	// To calculate the frequency map for each string per goroutine
	for _, s := range strings {
		go func(s string) {
			results <- Frequency(s)
		}(s)
	}

	// Sum of the frequency maps passed back via the results channel.
	for i := 0; i < count; i++ {
		for r, freq := range <-results {
			m[r] += freq
		}
	}

	return m
}

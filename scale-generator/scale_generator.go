package scale

/*
Sharpe: A, A#, B, C, C#, D, D#, E, F, F#, G, G#
Flats:  A, Bb, B, C, Db, D, Eb, E, F, Gb, G, Ab

No Sharps or Flats:
    C major a minor

Use Sharps:
    G, D, A, E, B, F# major e, b, f#, c#, g#, d# minor

Use Flats:
    F, Bb, Eb, Ab, Db, Gb major d, g, c, f, bb, eb minor

interval offset
 m        1
 M        2
 A        3
*/

var sharps = []string{
	//0    1     2    3    4     5    6     7    8    9    10    11
	"A", "A#", "B", "C", "C#", "D", "D#", "E", "F", "F#", "G", "G#",
}

var flats = []string{
	//0    1     2    3    4     5    6     7    8    9    10    11
	"A", "Bb", "B", "C", "Db", "D", "Eb", "E", "F", "Gb", "G", "Ab",
}

var sharpsMap = map[string]int{
	//tonic, index in the map sharps
	"C":  3,
	"G":  10,
	"D":  5,
	"A":  0,
	"E":  7,
	"B":  2,
	"F#": 9,

	"e":  7,
	"b":  2,
	"f#": 9,
	"c#": 4,
	"g#": 11,
	"d#": 6,

	"a": 0,
}

var flatsMap = map[string]int{
	//tonic, index in the map flats
	"F":  8,
	"Bb": 1,
	"Eb": 6,
	"Ab": 11,
	"Db": 4,
	"Gb": 9,

	"d":  5,
	"g":  10,
	"c":  3,
	"f":  8,
	"bb": 1,
	"eb": 6,
}

// Scale use a tonic, or starting note, and a set of intervals,
// generate the musical scale starting with the tonic and
// following the specified interval pattern.
func Scale(tonic, interval string) []string {
	var i, offset int
	var notes []string

	// If empty string, get all notes
	if interval == "" {
		interval = "mmmmmmmmmmmm"
	}

	// Used the sharps or flats
	i, ok := sharpsMap[tonic]
	notes = sharps
	if !ok {
		i, ok = flatsMap[tonic]
		notes = flats
	}

	scale := make([]string, len(interval))
	for j, c := range interval {
		scale[j] = notes[i]
		switch c {
		case 'm':
			offset = 1
		case 'M':
			offset = 2
		case 'A':
			offset = 3
		}
		i = (i + offset) % len(notes)
	}

	return scale
}

package strand

// ToRNA will transcrib DNA to RNA
func ToRNA(dna string) string {
	dna2rna := map[byte]byte{
		'C': 'G',
		'G': 'C',
		'T': 'A',
		'A': 'U',
	}

	b := []byte(dna)
	for i, c := range b {
		b[i] = dna2rna[c]
	}

	return string(b)
}

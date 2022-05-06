package chessboard

// Rank which stores if a square is occupied by a piece - this will be a slice of bools
type Rank []bool

// Chessboard which contains a map of eight Ranks, accessed with keys from "A" to "H"
type Chessboard map[string]Rank

// CountInRank returns how many squares are occupied in the chessboard,
// within the given rank
func CountInRank(cb Chessboard, rank string) int {
	rankV, ok:= cb[rank]; if !ok {
		return 0
	}
	
	var i int
	for _, v := range rankV {
		if v {
			i++
		}
	}

	return i
}

// CountInFile returns how many squares are occupied in the chessboard,
// within the given file
func CountInFile(cb Chessboard, file int) int {
	if file > 8 || file < 1 {
		return 0
	}

	var v int
	
	for _, rankV:= range cb {
		if rankV[file-1] {
			v++
		}
	}

	return v
}

// CountAll should count how many squares are present in the chessboard
func CountAll(cb Chessboard) int {
	return len(cb) * len(cb["A"])
}

// CountOccupied returns how many squares are occupied in the chessboard
func CountOccupied(cb Chessboard) int {
	var v int
	
	for _, rankV := range cb {
		for _, j := range rankV {
			if j {
				v++
			}
		}
	}

	return v
}

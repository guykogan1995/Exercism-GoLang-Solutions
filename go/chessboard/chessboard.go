package chessboard

// Declare a type named File which stores if a square is occupied by a piece - this will be a slice of bools
type File [8]bool

// Declare a type named Chessboard which contains a map of eight Files, accessed with keys from "A" to
type Chessboard map[string]File

// CountInFile returns how many squares are occupied in the chessboard,
// within the given file.
func CountInFile(cb Chessboard, file string) int {
	sum := 0
	for _, x := range cb[file] {
		if x == true {
			sum++
		}
	}
	return sum
}

// CountInRank returns how many squares are occupied in the chessboard,
// within the given rank.
func CountInRank(cb Chessboard, rank int) int {
	sum := 0
	if rank < 1 || rank > 8 {
		return 0
	} else {
		for _, x := range cb {
			if x[rank-1] == true {
				sum++
			}
		}
		return sum
	}
}

// CountAll should count how many squares are present in the chessboard.
func CountAll(cb Chessboard) int {
	rows := len(cb)
	cols := len(cb)
	return rows * cols
}

// CountOccupied returns how many squares are occupied in the chessboard.
func CountOccupied(cb Chessboard) int {
	sum := 0
	for rankKey, _ := range cb {
		sum += CountInFile(cb, rankKey)
	}
	return sum
}

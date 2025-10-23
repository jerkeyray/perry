package chess

import "fmt"

var (
	KingAttacks [64]uint64
	KnightAttacks [64]uint64
	BishopAttacks [64]uint64
	RookAttacks [64]uint64
	QueenAttacks [64]uint64
	WhitePawnAttacks [64]uint64
	BlackPawnAttacks [64]uint64
	AllAttacks [64]uint64
)

func init() {
	fmt.Println("Initializing attacks...")
	initLeaperAttacks()
	initPawnAttacks()

	fmt.Println("Attacks initialized")
}

// calculate knight and king attacks
func initLeaperAttacks() {
	knightOffsets := []int{17, 15, 10, 6, -17, -15, -10, -6}
	kingOffsets := []int{9, 8, 7, 1, -9, -8, -7, -1} 

	for sq := uint(0); sq < 64; sq++ {
		// knight attacks
		for _, offset := range knightOffsets {
			targetSq := int(sq) + offset
			if targetSq > 0 && targetSq < 64 {
				startFile := int(sq % 8)
				targetFile := targetSq % 8
				if abs(startFile - targetFile) <= 2 {
					KnightAttacks[sq] = setBit(KnightAttacks[sq], uint(targetSq))
				}
			}
		}

		// king attacks
		for _, offset := range kingOffsets {
			targetSq := int(sq) + offset
			if targetSq > 0 && targetSq < 64 {
				startFile := int(sq % 8)
				targetFile := targetSq % 8
				if abs(startFile - targetFile) <= 1 {
					KingAttacks[sq] = setBit(KingAttacks[sq], uint(targetSq))
				}
			}
		}
	}
}

func initPawnAttacks() {
	for sq := uint(0); sq < 64; sq++ {
		// White Pawn Attacks (North-East, North-West)
		targetNE := int(sq) + 9
		targetNW := int(sq) + 7

		if targetNE >= 0 && targetNE < 64 {
			startFile := int(sq % 8)
			targetFile := targetNE % 8
			if abs(startFile-targetFile) == 1 { // Check file wrap-around
				WhitePawnAttacks[sq] = setBit(WhitePawnAttacks[sq], uint(targetNE))
			}
		}
		if targetNW >= 0 && targetNW < 64 {
			startFile := int(sq % 8)
			targetFile := targetNW % 8
			if abs(startFile-targetFile) == 1 { // Check file wrap-around
				WhitePawnAttacks[sq] = setBit(WhitePawnAttacks[sq], uint(targetNW))
			}
		}

		// Black Pawn Attacks (South-East, South-West)
		targetSE := int(sq) - 7
		targetSW := int(sq) - 9

		if targetSE >= 0 && targetSE < 64 {
			startFile := int(sq % 8)
			targetFile := targetSE % 8
			if abs(startFile-targetFile) == 1 { // Check file wrap-around
				BlackPawnAttacks[sq] = setBit(BlackPawnAttacks[sq], uint(targetSE))
			}
		}
		if targetSW >= 0 && targetSW < 64 {
			startFile := int(sq % 8)
			targetFile := targetSW % 8
			if abs(startFile-targetFile) == 1 { // Check file wrap-around
				BlackPawnAttacks[sq] = setBit(BlackPawnAttacks[sq], uint(targetSW))
			}
		}
	}
}

func abs(x int) int {
	if(x < 0) {
		return -x
	}
	return x
}
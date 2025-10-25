package chess

import "fmt"

// Move represents a chess move compactly using a uint32.
// Bits layout:
// 0-5:   From Square (6 bits, 0-63)
// 6-11:  To Square   (6 bits, 0-63)
// 12-14: Promotion Piece Type (3 bits, Knight=1, Bishop=2, Rook=3, Queen=4) - Only used if FlagPromotion is set.
// 15:    Unused (or potentially for capture piece type later)
// 16-19: Flags (Promotion, Capture, EnPassant, Castle, DoublePawnPush)
// ... remaining bits unused for now ...
type Move uint32

const (
	// --- Square Masks (Binary Notation) ---
	fromSquareMask Move = 0b00000000000000000000000000111111 // Lower 6 bits
	toSquareMask   Move = 0b00000000000000000000111111000000 // Next 6 bits

	// --- Promotion Piece Type --- (Relative type, NOT shifted piece constants)
	// Stored in bits 12-14
	promoTypeMask Move = 0b00000000000000000111000000000000 // Bits 12, 13, 14
	promoTypeKnight uint = 1 // Use these small ints when creating promo moves
	promoTypeBishop uint = 2
	promoTypeRook   uint = 3
	promoTypeQueen  uint = 4

	// --- Flags (Using bits 16 onwards) ---
	FlagPromotion      Move = 1 << 16 // Indicates a pawn promotion
	FlagCapture        Move = 1 << 17 // Indicates a capture (including en passant)
	FlagEnPassant      Move = 1 << 18 // Indicates an en passant capture specifically
	FlagCastleKing     Move = 1 << 19 // Indicates king-side castling
	FlagCastleQueen    Move = 1 << 20 // Indicates queen-side castling
	FlagDoublePawnPush Move = 1 << 21 // Indicates a pawn moving two squares forward
)

// --- Helper Constructors ---

// NewMove creates a standard quiet (non-special) move.
func NewMove(from, to uint) Move {
	return Move(from) | (Move(to) << 6)
}

// NewCaptureMove creates a standard capture move.
func NewCaptureMove(from, to uint) Move {
	return NewMove(from, to) | FlagCapture
}

// NewPromotionMove creates a promotion move (quiet promotion).
// promoPieceType should be one of promoTypeKnight, promoTypeBishop, etc.
func NewPromotionMove(from, to uint, promoPieceType uint) Move {
	return NewMove(from, to) | (Move(promoPieceType) << 12) | FlagPromotion
}

// NewPromotionCaptureMove creates a promotion move that is also a capture.
func NewPromotionCaptureMove(from, to uint, promoPieceType uint) Move {
	return NewPromotionMove(from, to, promoPieceType) | FlagCapture
}

// NewEnPassantMove creates an en passant capture move.
func NewEnPassantMove(from, to uint) Move {
	return NewMove(from, to) | FlagCapture | FlagEnPassant
}

// NewCastleKingMove creates a king-side castling move.
func NewCastleKingMove(from, to uint) Move {
	return NewMove(from, to) | FlagCastleKing
}

// NewCastleQueenMove creates a queen-side castling move.
func NewCastleQueenMove(from, to uint) Move {
	return NewMove(from, to) | FlagCastleQueen
}

// NewDoublePawnPush creates a double pawn push move.
func NewDoublePawnPush(from, to uint) Move {
	return NewMove(from, to) | FlagDoublePawnPush
}

// --- Getter Methods ---

func (m Move) From() uint {
	return uint(m & fromSquareMask)
}

func (m Move) To() uint {
	return uint((m & toSquareMask) >> 6)
}

// PromotionType returns the type of piece promoted to (e.g., promoTypeKnight)
// or 0 if it's not a promotion move.
func (m Move) PromotionType() uint {
	if !m.IsPromotion() {
		return 0 // Or a specific 'NoPromo' constant
	}
	return uint((m & promoTypeMask) >> 12)
}

func (m Move) IsPromotion() bool {
	return (m & FlagPromotion) != 0
}

func (m Move) IsCapture() bool {
	return (m & FlagCapture) != 0
}

func (m Move) IsEnPassant() bool {
	return (m & FlagEnPassant) != 0
}

func (m Move) IsCastleKing() bool {
	return (m & FlagCastleKing) != 0
}

func (m Move) IsCastleQueen() bool {
	return (m & FlagCastleQueen) != 0
}

func (m Move) IsCastle() bool { // Convenience for either castle
	return (m & (FlagCastleKing | FlagCastleQueen)) != 0
}

func (m Move) IsDoublePawnPush() bool {
	return (m & FlagDoublePawnPush) != 0
}

// --- String Representation (UCI Format) ---

// String converts the move to UCI standard notation (e.g., "e2e4", "e7e8q").
func (m Move) String() string {
	from := m.From()
	to := m.To()
	promoType := m.PromotionType()

	// Convert square indices (0-63) back to algebraic notation (a1-h8)
	fromStr := fmt.Sprintf("%c%d", 'a'+(from%8), 1+(from/8))
	toStr := fmt.Sprintf("%c%d", 'a'+(to%8), 1+(to/8))

	promoStr := ""
	if m.IsPromotion() {
		switch promoType {
		case promoTypeKnight:
			promoStr = "n"
		case promoTypeBishop:
			promoStr = "b"
		case promoTypeRook:
			promoStr = "r"
		case promoTypeQueen:
			promoStr = "q"
		}
	}

	return fromStr + toStr + promoStr
}
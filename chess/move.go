package chess

import "fmt"

// Move represents a chess move compactly.
// Bits layout (example using uint16):
// 0-5:   From Square (6 bits, 0-63)
// 6-11:  To Square   (6 bits, 0-63)
// 12-14: Promotion Piece Type (3 bits, Knight=1, Bishop=2, Rook=3, Queen=4)
// 15:    Special Flag (e.g., Castling, En Passant, Promotion) - More bits might be needed
type Move uint16

const (
	// Masks to extract information
	fromSquareMask Move = 0b0000000000111111 // Lower 6 bits
	toSquareMask   Move = 0b0000111111000000 // Next 6 bits
	promoPieceMask Move = 0b0111000000000000 // Next 3 bits
	specialFlagMask Move = 0b1000000000000000 // Top bit (can expand later)

    // Promotion piece type values (shifted to correct position)
    PromoKnight Move = 1 << 12
    PromoBishop Move = 2 << 12
    PromoRook   Move = 3 << 12
    PromoQueen  Move = 4 << 12

    // Special Flags (example, needs more thought)
    FlagCastle    Move = 1 << 15
    FlagEnPassant Move = 1 << 15 // Overlap? Need more bits or smarter encoding
    FlagPromotion Move = 1 << 15
)

// NewMove creates a standard move.
func NewMove(from, to uint) Move {
	return Move(from) | (Move(to) << 6)
}

// NewPromotionMove creates a promotion move.
func NewPromotionMove(from, to uint, promoType Move) Move {
    // Assuming promoType is already shifted (PromoKnight, etc.)
	return NewMove(from, to) | promoType | FlagPromotion // Set promotion flag too
}

// Getters for move properties
func (m Move) From() uint {
	return uint(m & fromSquareMask)
}

func (m Move) To() uint {
	return uint((m & toSquareMask) >> 6)
}

func (m Move) Promotion() Move { // Returns the shifted value (PromoKnight etc.) or 0
	return m & promoPieceMask
}

func (m Move) IsCastle() bool {
	// Example: Needs refinement based on how you encode flags
	return (m & specialFlagMask) != 0 // && ... specific check for castle type
}
// ... Add IsEnPassant(), IsPromotion() etc. ...

// String representation (useful for debugging)
func (m Move) String() string {
    from := m.From()
    to := m.To()
    promo := m.Promotion()
    // Convert square indices (0-63) back to algebraic notation (a1-h8)
    fromStr := fmt.Sprintf("%c%d", 'a'+(from%8), 1+(from/8))
    toStr := fmt.Sprintf("%c%d", 'a'+(to%8), 1+(to/8))
    promoStr := ""
    switch promo {
        case PromoKnight: promoStr = "n"
        case PromoBishop: promoStr = "b"
        case PromoRook: promoStr = "r"
        case PromoQueen: promoStr = "q"
    }
	return fromStr + toStr + promoStr // UCI format
}
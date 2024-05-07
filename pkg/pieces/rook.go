package pieces

import "github.com/ikaroly/gobot/pkg/bitboard"



type Rook interface {
	RookHorizontal() uint64
	RookVertical() uint64
}

func (p Piece) RookHorizontal() uint64{
	// bitboard.GetRank(p.Position) 
	return uint64(256 << 8 * (8 - bitboard.GetRank(p.Position)))
}

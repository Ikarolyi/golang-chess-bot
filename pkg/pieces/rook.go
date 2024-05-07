package pieces

import "github.com/ikaroly/gobot/pkg/bitboard"



type Rook interface {
	RookHorizontal() uint64
	RookVertical() uint64
	RookCombined() uint64
}

func (p Piece) RookHorizontal() uint64{
	// bitboard.GetRank(p.Position) 
	return uint64(bitboard.Rank8 << (8 * (8 - bitboard.GetRank(p.Position))))
}

func (p Piece) RookVertical() uint64{
	return uint64(bitboard.FileA >> bitboard.GetFile(p.Position))
}

func (p Piece) RookCombined() uint64{
	return p.RookVertical() & p.RookHorizontal()
}
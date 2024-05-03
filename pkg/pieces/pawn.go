package pieces

import "github.com/ikaroly/gobot/pkg/bitboard"

type Pawn interface {
	pawn_double_pushes() uint64
	pawn_pushes() uint64
	pawn_kicks() uint64
	is_pawn_untouched() bool
	en_passant() uint64
}

func (p Piece) is_pawn_untouched() bool{
	if p.IsWhite{
		return bitboard.GetRank(p.Position) == 8
	}else{
		return bitboard.GetRank(p.Position) == 0
	}
}

func (p Piece) pawn_kicks() uint64{
	var retVal uint64

	// Push forward
	if p.IsWhite{
		retVal = p.Position << 8
	}else{
		retVal = p.Position >> 8
	}
	

	retVal = (retVal << 1) | (retVal >> 1)

	return retVal
}

func (p Piece) pawn_pushes() uint64{
	var retVal uint64

	//first row
	if p.IsWhite{
		retVal = p.Position << 8
	}else{
		retVal = p.Position >> 8
	}

	if p.is_pawn_untouched(){
		retVal |= p.pawn_double_pushes()
	}

	return retVal
}

func (p Piece) pawn_double_pushes() uint64{
	var retVal uint64

	//first row
	if p.IsWhite{
		retVal = p.Position << 16
	}else{
		retVal = p.Position >> 16
	}

	return retVal
}

func (p Piece) en_passant() uint64 {
	// TODO: ?????????????????????????????????????????????????????????????????????????????????????????????????????????????
	return 0
}
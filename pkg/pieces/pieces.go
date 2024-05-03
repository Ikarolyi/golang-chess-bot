package pieces

import "github.com/ikaroly/gobot/pkg/bitboard"

// import "github.com/ikaroly/gobot/pkg/bitboard"

//sakkok
const PAWN = 0

//tiszák
const KNIGHT = 1
const BISHOP = 2
const ROOK = 3
const QUEEN = 4
const KING = 5


type Piece struct {
	IsWhite bool;
	Position uint64
	Class int
}

type Moves interface {
	get_moves() uint64
	get_value() int
	on_back_rank() bool
}

func (p Piece) get_moves() uint64{
	switch class := p.Class; class{
		case PAWN:
			return p.pawn_pushes()
		case KNIGHT:
			return 3
		case BISHOP:
			return 3
		case ROOK:
			return 5
		case QUEEN:
			return 9
		case KING:
			return 0
	}
	return 0 //Shall never happen
}

func (p Piece) get_value() int{
	switch class := p.Class; class{
		case PAWN:
			return 1
		case KNIGHT:
			return 3
		case BISHOP:
			return 3
		case ROOK:
			return 5
		case QUEEN:
			return 9
		case KING:
			return 0
	}
	return 0 //Shall never happen
}

func (p Piece) on_back_rank() bool{
	if p.IsWhite{
		return bitboard.GetRank(p.Position) == 8
	}else{
		return bitboard.GetRank(p.Position) == 0
	}
} 
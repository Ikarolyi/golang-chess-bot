package pieces

import (
	"math"
	"unicode"

	"github.com/ikaroly/gobot/pkg/bitboard"
)

// import "github.com/ikaroly/gobot/pkg/bitboard"

//sakkok
const PAWN = 0

//tiszák
const KNIGHT = 1
const BISHOP = 2
const ROOK = 3
const QUEEN = 4
const KING = 5

const WHITE = 1
const BLACK = -1


type Piece struct {
	Color int;
	Position uint64
	Class int
}

type Moves interface {
	get_moves() uint64
	GetValue() int
	on_back_rank() bool
	GetColorStr() string
	ToString() string
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

func (p Piece) GetValue() int{
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
	if p.Color == WHITE{
		return bitboard.GetRank(p.Position) == 8
	}else{
		return bitboard.GetRank(p.Position) == 0
	}
}

func NewPiece(char rune, exponent float64) Piece{
	var new_piece = new(Piece)
	switch unicode.ToUpper(char) {
		case 'P':
			new_piece.Class = PAWN
		case 'N':
			new_piece.Class = KNIGHT
		case 'B':
			new_piece.Class = BISHOP
		case 'R':
			new_piece.Class = ROOK
		case 'Q':
			new_piece.Class = QUEEN
		case 'K':
			new_piece.Class = KING
	}

	if unicode.IsUpper(char) {
		new_piece.Color = WHITE
	}else{
		new_piece.Color = BLACK
	}


	new_piece.Position = uint64(math.Pow(2, exponent))

	return *new_piece
}

func (p Piece) GetColorStr() string{
	if p.Color == WHITE{
		return "W"
	}else{
		return "B"
	}
}

func (p Piece) ToString() string{
	var classRepr string
	switch p.Class{
		case PAWN:
			classRepr = "P"
		case BISHOP:
			classRepr = "B"
		case KNIGHT:
			classRepr = "N"
		case ROOK:
			classRepr = "R"
		case QUEEN:
			classRepr = "Q"
		case KING:
			classRepr = "K"
	}

	return p.GetColorStr() + " " + classRepr + "@" + bitboard.ToString(p.Position)
}
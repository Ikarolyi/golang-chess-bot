package pieces

import (
	"math"
	"unicode"

	"github.com/ikaroly/gobot/pkg/bitboard"
)


const PAWN = 0

const KNIGHT = 1
const BISHOP = 2
const ROOK = 3
const QUEEN = 4
const KING = 5

const WHITE int8 = 1
const BLACK int8 = -1

const EMPTYINDEX int = -1


type Piece struct {
	Color int8;
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

func (p Piece) GetMoves(boardCombined bitboard.CombinedBoard, enPassantTarget uint64) []bitboard.Move{
	var result []bitboard.Move
	switch class := p.Class; class{
		case PAWN:
			result = p.GetPawnMoves(boardCombined, enPassantTarget)
		case KNIGHT:
			result = p.GetKnightMoves(boardCombined)
		case BISHOP:
			result = p.GetBishopMoves(boardCombined)
		case ROOK:
			result = p.GetRookMoves(boardCombined)
		case QUEEN:	
			result = p.GetQueenMoves(boardCombined)
		case KING:
			result = nil			
	}

	return result
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

	return p.GetColorStr() + " " + classRepr + "@" + bitboard.Decode(p.Position)
}
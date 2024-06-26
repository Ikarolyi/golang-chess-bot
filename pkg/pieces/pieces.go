package pieces

import (
	"unicode"

	"github.com/ikaroly/gobot/pkg/bitboard"
	"github.com/ikaroly/gobot/pkg/tables"
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
	BoundCastlingAbilities byte
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
			result = p.GetKingMoves(boardCombined)
	}

	return result
}

func GetThreatenedSquares(boardCombined bitboard.CombinedBoard, color int8, pieces []Piece) uint64{
	var result uint64
	for _, piece := range pieces{
		if piece.Color != color{
			continue
		}
		
		var pieceMoves []bitboard.Move
		if piece.Class != PAWN{
			pieceMoves = piece.GetMoves(boardCombined, 0)
		}else{
			pieceMoves = piece.pawnKick(boardCombined, 0)
		}

		for _, pieceMove := range pieceMoves{
			result |= pieceMove.To
		}
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
			return 65536
	}
	return 0 //Shall never happen
}

func NewPiece(char rune, fenIndex int) Piece{
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


	// Init piece position (complicated formula to convert fen order to normal)
	new_piece.Position = 1 << (56 - (int(fenIndex/8) * 8) + (fenIndex % 8)) // TODO separating formula

	// Init the castling abilities bound to the piece
	switch *new_piece{
		case Piece{Class: KING, Color: WHITE, Position: tables.CastlingKingMove[0].From}:
			new_piece.BoundCastlingAbilities = tables.CastlingAbility[0] | tables.CastlingAbility[1]
		case Piece{Class: KING, Color: BLACK, Position: tables.CastlingKingMove[2].From}:
			new_piece.BoundCastlingAbilities = tables.CastlingAbility[2] | tables.CastlingAbility[3]
		case Piece{Class: ROOK, Color: WHITE, Position: tables.CastlingRookMove[0].From}:
			new_piece.BoundCastlingAbilities = tables.CastlingAbility[0]
		case Piece{Class: ROOK, Color: WHITE, Position: tables.CastlingRookMove[1].From}:
			new_piece.BoundCastlingAbilities = tables.CastlingAbility[1]
		case Piece{Class: ROOK, Color: BLACK, Position: tables.CastlingRookMove[2].From}:
			new_piece.BoundCastlingAbilities = tables.CastlingAbility[2]
		case Piece{Class: ROOK, Color: BLACK, Position: tables.CastlingRookMove[3].From}:
			new_piece.BoundCastlingAbilities = tables.CastlingAbility[3]
	}
	
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
	var result rune
	switch p.Class{
		case PAWN:
			result = 'P'
		case BISHOP:
			result = 'B'
		case KNIGHT:
			result = 'N'
		case ROOK:
			result = 'R'
		case QUEEN:
			result = 'Q'
		case KING:
			result = 'K'
	}

	if p.Color == BLACK{
		result = unicode.ToLower(result)
	}

	return string(result)
}
package pieces

import "github.com/ikaroly/gobot/pkg/bitboard"

type Pawn interface {
	GetPawnMoves() []bitboard.Move
}

var pawnForward = bitboard.Vector{X: 0, Y: 1}

var pawnKickVectors = [...]bitboard.Vector{
	{X: 1, Y: 1},
	{X: -1, Y: 1},
}

func (p Piece) pawnKick(boardCombined bitboard.CombinedBoard, enPassantTarget uint64) []bitboard.Move{
	var result []bitboard.Move

	if p.Color == WHITE{
		boardCombined.Black |= enPassantTarget
	}else{
		boardCombined.White |= enPassantTarget
	}

	for _, kickMove := range pawnKickVectors{
		print()
		result = append(result, bitboard.RayCastMovement(p.Position, -p.Color, boardCombined, kickMove, 1, true)...)
	}


	return result
}

func (p Piece) GetPawnMoves(boardCombined bitboard.CombinedBoard, enPassantTarget uint64) []bitboard.Move {
	var result []bitboard.Move

	var push_len = 1
	var on_default_place = false

	if p.Color == WHITE{
		on_default_place = (bitboard.GetRank(p.Position) == 2)
	}else{
		on_default_place = (bitboard.GetRank(p.Position) == 7)
	}

	if on_default_place{
		push_len = 2
	}

	result = append(result, pawnPushRayCast(p.Position, p.Color, boardCombined, pawnForward.Multiply(int(p.Color)), push_len)...)

	result = append(result, p.pawnKick(boardCombined, enPassantTarget)...)


	return result
}

func pawnPushRayCast(square uint64, color int8, boardCombined bitboard.CombinedBoard, direction bitboard.Vector, limit int) []bitboard.Move {
	boardCombined.White = boardCombined.GetTrueCombined()
	boardCombined.Black = boardCombined.White

	result := bitboard.RayCastMovement(square, color, boardCombined, direction, limit, false)
	
	// Init en passant value on the double push, with the value of the short push's Move.To
	if len(result) == 2 {
		result[1].NewEnPassantTarget = result[0].To
	}

	return result
}